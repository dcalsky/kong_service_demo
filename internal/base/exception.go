package base

import (
	"fmt"
)

type Exception struct {
	error             // original error
	StatusCode int    // error http status code
	Code       string // error code
	Message    string // standard error message to user
	Public     bool   // public to user
	RawError   string // original error text
}

func NewException(statusCode int, textCode, msg string) Exception {
	return Exception{
		StatusCode: statusCode,
		Code:       textCode,
		Message:    msg,
		Public:     false,
		RawError:   "",
	}
}

func (e Exception) Error() string {
	return fmt.Sprintf("StatusCode: %v, Code: %v, Message: %v, RawError: %v", e.StatusCode, e.Code, e.Message, e.RawError)
}

func (e Exception) WithRawError(err error) Exception {
	e.error = err
	e.RawError = ""
	if e.error != nil {
		e.RawError = e.error.Error()
	}
	return e
}

func (e Exception) WithMessage(msg string) Exception {
	e.Message = msg
	return e
}
