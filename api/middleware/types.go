package middleware

import (
	"net/http"
)

// HandlerFuncWithCTX - type is an adapter to use handlerfunc with ctx
type HandlerFuncWithCTX func(*Context, http.ResponseWriter, *http.Request) error

type StatusCodeRecorder struct {
	http.ResponseWriter
	http.Hijacker
	StatusCode int
}

func (r *StatusCodeRecorder) WriteHeader(statusCode int) {
	r.StatusCode = statusCode
	r.ResponseWriter.WriteHeader(statusCode)
}
