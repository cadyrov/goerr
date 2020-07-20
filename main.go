package goerr

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

type AppError struct {
	code    int
	Message string   `json:"message"`
	Detail  []IError `json:"detail,omitempty"`
}

func (e *AppError) PushDetail(ae IError) {
	e.Detail = append(e.Detail, ae)
}

func (e *AppError) Error() (er string) {
	er += fmt.Sprintf("Code: %v; ", e.code)
	er += "Msg: " + e.Message + ";  "
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

func (e *AppError) GetCode() int {
	return e.code
}

func (e *AppError) GetMessage() string {
	return e.Message
}

func (e *AppError) GetDetails() []IError {
	return e.Detail
}

func (e *AppError) Http(code int) IError {
	e.code = code
	return e
}

func New(message string) (e IError) {
	e = &AppError{code: http.StatusInternalServerError, Message: message}
	return
}
