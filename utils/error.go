package utils

import (
	"fmt"
	"net/http"
)

const (
	PageNotFound ErrorStatus = iota + 1
	NotAuthorized
	Forbidden
	NotFound
	IncorrectParam
	FieldNotPermitted
	OperationFailed
	InternalError
	Success
	MultiStatus
	NoContent
	CallServicesFailed
	GetDataFromDB
)

type ErrorStatus int

type errorMessage struct {
	message string
	code    int
}

var messages = []errorMessage{
	PageNotFound:       {message: "Page not found", code: http.StatusNotFound},
	NotAuthorized:      {message: "You are not authorized", code: http.StatusUnauthorized},
	Forbidden:          {message: "Forbidden endpoint", code: http.StatusForbidden},
	NotFound:           {message: "Data not found", code: http.StatusNotFound},
	IncorrectParam:     {message: "Wrong parameter", code: http.StatusBadRequest},
	FieldNotPermitted:  {message: "Data on field '%s' not permitted", code: http.StatusBadRequest},
	OperationFailed:    {message: "%s failed", code: http.StatusBadRequest},
	InternalError:      {message: "Internal error", code: http.StatusInternalServerError},
	Success:            {message: "%s success", code: http.StatusOK},
	MultiStatus:        {message: "Parallel Multistatus", code: http.StatusMultiStatus},
	NoContent:          {message: "No Content", code: http.StatusNoContent},
	CallServicesFailed: {message: "Failed connect to services %s", code: http.StatusInternalServerError},
	GetDataFromDB:      {message: "getting data from DB %s Failed", code: http.StatusBadRequest},
}

// Error used for get message for display
type Error struct {
	Code       string `json:"code,omitempty"`
	Message    string `json:"message,omitempty"`
	StatusCode int    `json:"status_code,omitempty"`
}

// SetError used for set 'Error' structure
func SetError(code ErrorStatus, err error, fields ...string) *Error {
	if err == nil {
		return nil
	}

	if int(code) > len(messages) {
		code = InternalError
	}
	msg := messages[code]
	switch code {
	case FieldNotPermitted, OperationFailed, Success, MultiStatus, CallServicesFailed, GetDataFromDB:
		if len(fields) > 0 {
			msg.message = fmt.Sprintf(msg.message, fields[0])
		}
	}
	return &Error{
		Message: msg.message,
		Code:    fmt.Sprintf("%03d", code),
	}
}
