package apperror

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"runtime"
	"time"

	"task-management/util"
)

var (
	ErrPermission   = &AppError{Code: http.StatusForbidden, Message: ErrMsgPermission}
	ErrBadRequest   = &AppError{Code: http.StatusBadRequest, Message: ErrMsgBadRequest}
	ErrJsonDecode   = &AppError{Code: http.StatusBadRequest, Message: ErrMsgJsonDecode}
	ErrNotFound     = &AppError{Code: http.StatusNotFound, Message: ErrMsgNotFound}
	ErrDatabase     = &AppError{Code: http.StatusInternalServerError, Message: ErrMsgDatabase}
	ErrAccessDenied = &AppError{Code: http.StatusForbidden, Message: ErrMsgAccessDenied}
	ErrUnauthorized = &AppError{Code: http.StatusUnauthorized, Message: ErrMsgUnauthorized}
	ErrRateLimited  = &AppError{Code: http.StatusTooManyRequests, Message: ErrMsgRateLimited}
	ErrServer       = &AppError{Code: http.StatusInternalServerError, Message: ErrMsgServer}
	Info            = &AppError{Code: http.StatusOK, Message: LogMsgInfo}
)

type AppError struct {
	Code    int
	Message string
	File    string
	Line    int
}

func (e *AppError) Error() string {
	if e.File == "" || e.Line == 0 {
		_, file, line, _ := runtime.Caller(1)
		e.File = file
		e.Line = line
	}

	return e.Message
}

func (e *AppError) StatusCode() int {
	return e.Code
}

func (e *AppError) Customize(message string) *AppError {
	return &AppError{
		Code:    e.Code,
		Message: message,
	}
}

// WithLocation creates a new AppError with file and line info
func (e *AppError) LogWithLocation() *AppError {
	_, file, line, _ := runtime.Caller(1) // Capture the caller's location
	e.File = file
	e.Line = line
	log.SetFlags(0)
	log.Printf("%sError: %s \n File: %s \n Line: %d \n Date: %s%s", "\033[31m",
		e.Message, file, line, time.Now().Format(time.RFC3339), "\033[0m")
	return e
}

func Respond(w http.ResponseWriter, err error) {
	if appErr, ok := err.(*AppError); ok {
		w.WriteHeader(appErr.Code)
		response := util.SetResponse(nil, 0, appErr.Message)
		json.NewEncoder(w).Encode(response)
		return
	}

	json.NewEncoder(w).Encode(util.SetResponse(nil, 0, fmt.Sprintf("%s : %s", ErrMsgUnhandledError, err.Error())))
	http.Error(w, err.Error(), ErrServer.Code)
}
