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

func PanicIf(expression bool, exception Exception) {
	if expression {
		panic(exception)
	}
}

func PanicIfErr(err error, exception Exception) {
	if err != nil {
		panic(exception)
	}
}

func Catch(f func()) (e *Exception) {
	defer func() {
		if panicData := recover(); panicData != nil {
			if except, ok := panicData.(Exception); ok {
				e = &except
			} else {
				temp := InternalError.WithRawError(fmt.Errorf("%v", panicData))
				e = &temp
			}
			return
		}
	}()
	f()
	return
}
