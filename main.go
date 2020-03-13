package errs

import (
	"fmt"
	"net/http"
)

type IError interface {
	Error() string
	GetCode() int
	GetDetails() []IError
	PushDetail(IError)
	GetMessage() string
	Http(code int) IError
}

type appError struct {
	code    int
	message string
	detail  []IError
}

func (e *appError) PushDetail(ae IError) {
	e.detail = append(e.detail, ae)
}

func (e *appError) Error() (er string) {
	er += fmt.Sprintf("Code: %v; ", e.code)
	er += "Msg: " + e.message + ";  "
	if len(e.GetDetails()) == 0 {
		return
	}
	er += " Details: {"
	for idx := range e.GetDetails() {
		er += e.GetDetails()[idx].Error()
	}
	er += "}"
	return
}

func (e *appError) GetCode() int {
	return e.code
}

func (e *appError) GetMessage() string {
	return e.message
}

func (e *appError) GetDetails() []IError {
	return e.detail
}

func (e *appError) Http(code int) IError {
	e.code = code
	return e
}

func New(message string) (e IError) {
	e = &appError{code: http.StatusInternalServerError, message: message}
	return
}
