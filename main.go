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
	HTTP(code int) IError
	SetID(string) IError
	GetID() string
}

type AppError struct {
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Detail  []IError `json:"detail,omitempty"`
	ID      string   `json:"id,omitempty"`
}

func (e *AppError) PushDetail(ae IError) {
	e.Detail = append(e.Detail, ae)
}

func (e *AppError) SetID(name string) IError {
	e.ID = name
	return e
}

func (e *AppError) GetID() string {
	return e.ID
}

func (e *AppError) Error() (er string) {
	er += fmt.Sprintf("Code: %v; ", e.Code)

	er += "Msg: " + e.Message + ";  "

	if len(e.GetDetails()) == 0 {
		return
	}

	er += " Details: {"

	for idx := range e.GetDetails() {
		er += e.GetDetails()[idx].Error()
	}

	er += "}"

	return er
}

func (e *AppError) GetCode() int {
	return e.Code
}

func (e *AppError) GetMessage() string {
	return e.Message
}

func (e *AppError) GetDetails() []IError {
	return e.Detail
}

func (e *AppError) HTTP(code int) IError {
	e.Code = code

	return e
}

// deprecated
func New(message string) IError {
	e := &AppError{Code: http.StatusInternalServerError, Message: message}

	return e
}

func create(message string) IError {
	e := &AppError{Code: http.StatusInternalServerError, Message: message}

	return e
}

func BadRequest(message string) IError {
	return create(message).HTTP(http.StatusBadRequest)
}

func Unauthorized(message string) IError {
	return create(message).HTTP(http.StatusUnauthorized)
}

func Forbidden(message string) IError {
	return create(message).HTTP(http.StatusForbidden)
}

func NotFound(message string) IError {
	return create(message).HTTP(http.StatusNotFound)
}

func NotAcceptable(message string) IError {
	return create(message).HTTP(http.StatusNotAcceptable)
}

func Conflict(message string) IError {
	return create(message).HTTP(http.StatusConflict)
}

func Unprocessable(message string) IError {
	return create(message).HTTP(http.StatusUnprocessableEntity)
}

func Internal(message string) IError {
	return create(message).HTTP(http.StatusInternalServerError)
}




