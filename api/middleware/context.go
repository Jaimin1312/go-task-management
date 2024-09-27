package middleware

import (
	"github.com/sirupsen/logrus"
)

// Context per request state
type Context struct {
	Logger        logrus.FieldLogger
	RemoteAddress string
	UserID        string
}

// WithLogger sets logger for context
func (ctx *Context) WithLogger(logger logrus.FieldLogger) *Context {
	ret := *ctx
	ret.Logger = logger
	return &ret
}

// WithRemoteAddress sets remote address for context
func (ctx *Context) WithRemoteAddress(address string) *Context {
	ret := *ctx
	ret.RemoteAddress = address
	return &ret
}

// WithAccount sets account for context
func (ctx *Context) WithUser(userID string) *Context {
	ret := *ctx
	ret.UserID = userID
	return &ret
}
