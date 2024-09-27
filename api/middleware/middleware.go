package middleware

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"runtime/debug"
	"strings"
	"task-management/app/jwtauth"
	"task-management/apperror"
	"task-management/util"
	"time"

	"github.com/sirupsen/logrus"
)

type MiddlewareConfig struct {
	CookieName     string
	MaxContentSize int64
	ProxyCount     int
	JwtService     jwtauth.Service
}

func (a *MiddlewareConfig) MiddlewareHandler(f HandlerFuncWithCTX, auth ...bool) http.HandlerFunc {
	checkAuth := auth[0]

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		ctx := &Context{
			Logger: logrus.StandardLogger(),
		}

		fmt.Println("API -", r.URL.Path)
		r.Body = http.MaxBytesReader(w, r.Body, a.MaxContentSize*1024*1024)
		beginTime := time.Now()
		hijacker, _ := w.(http.Hijacker)
		ctx = ctx.WithRemoteAddress(a.IPAddressForRequest(r))
		ctx = ctx.WithLogger(ctx.Logger.WithField("request_id", base64.RawURLEncoding.EncodeToString(util.NewID())))

		w = &StatusCodeRecorder{
			ResponseWriter: w,
			Hijacker:       hijacker,
		}
		if checkAuth {
			userID, err := validateUser(a, r)
			if err != nil {
				apperror.Respond(w, err)
				return
			}

			ctx = ctx.WithUser(userID)
		}

		defer func() {
			statusCode := w.(*StatusCodeRecorder).StatusCode
			if statusCode == 0 {
				statusCode = 200
			}

			duration := time.Since(beginTime)
			logger := ctx.Logger.WithFields(logrus.Fields{
				"duration":    duration,
				"status_code": statusCode,
				"remote":      ctx.RemoteAddress,
			})

			logger.Info(r.Method + " " + r.URL.RequestURI())
		}()

		defer func() {
			if localRecover := recover(); localRecover != nil {
				ctx.Logger.Error(fmt.Errorf("recovered from panic\n %v: %s", localRecover, debug.Stack()))
				apperror.Respond(w, apperror.ErrServer)
				return
			}
		}()

		if err := f(ctx, w, r); err != nil {
			apperror.Respond(w, err)
			return
		}
	}
}

// IPAddressForRequest determines IP address for request
func (a *MiddlewareConfig) IPAddressForRequest(r *http.Request) string {
	addr := r.RemoteAddr
	if a.ProxyCount > 0 {
		h := r.Header.Get("X-Forwarded-For")
		if h != "" {
			clients := strings.Split(h, ",")
			if a.ProxyCount > len(clients) {
				addr = clients[0]
			} else {
				addr = clients[len(clients)-a.ProxyCount]
			}
		}
	}
	return strings.Split(strings.TrimSpace(addr), ":")[0]
}

// Basic authentication middleware
func BasicAuthMiddleware(next http.HandlerFunc, username, password string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()
		if !ok {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			apperror.Respond(w, apperror.ErrUnauthorized)
			return
		}

		if user != username || pass != password {
			apperror.Respond(w, apperror.ErrUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}
}
