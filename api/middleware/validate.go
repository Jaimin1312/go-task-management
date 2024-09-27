package middleware

import (
	"net/http"
	"task-management/apperror"
)

func validateUser(middlewareConf *MiddlewareConfig, r *http.Request) (string, error) {
	var token string

	token = r.Header.Get(middlewareConf.CookieName)
	if token == "" {
		c, err := r.Cookie(middlewareConf.CookieName)
		if err != nil || c.Value == "" {
			return "", apperror.ErrUnauthorized.Customize("Token is not present").LogWithLocation()
		}

		token = c.Value
	}

	claims, err := middlewareConf.JwtService.Validatejwt(token)
	if err != nil {
		return "", apperror.ErrUnauthorized
	}

	return claims.UserID, nil
}
