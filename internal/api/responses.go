package api

import (
	"net/http"
)

type Error struct {
	Message string `json:"message,omitempty"`
	Code    int64  `json:"code,omitempty"`
}

type Response struct {
	Result string      `json:"result"`
	Error  *Error      `json:"error,omitempty"`
	Data   interface{} `json:"data"`
}

func (r *Response) WithData(data interface{}) *Response {
	r.Data = data
	return r
}

var (
	Ok                  = Response{Result: "ok"}
	ErrorUnauthorized   = &Error{Code: http.StatusUnauthorized, Message: "Unauthorized"}
	ErrorInternalServer = &Error{Code: http.StatusInternalServerError, Message: "Internal Server Error"}
	ErrorForbidden      = &Error{Code: http.StatusForbidden, Message: "Forbidden"}
	ErrorBadRequest     = &Error{Code: http.StatusBadRequest, Message: "Bad Request"}
	ErrorNotFound       = &Error{Code: http.StatusNotFound, Message: "Not Found"}
)

func responseWithError(err *Error) Response {
	return Response{
		Result: "error",
		Error:  err,
	}
}
