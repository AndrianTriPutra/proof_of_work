package util

import (
	"errors"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

var ErrorNotFound = errors.New("content not found")

var (
	ErrInternalServer   = errors.New("internal server error")
	ErrBadRequest       = errors.New("bad request")
	ErrNotFound         = errors.New("not found")
	ErrConflict         = errors.New("resource conflict")
	ErrUnauthorized     = errors.New("unauthorized access")
	ErrPasswordMismatch = bcrypt.ErrMismatchedHashAndPassword
)

func GetHttpStatusCode(err error) int {
	switch err {
	case ErrInternalServer:
		return http.StatusInternalServerError
	case ErrBadRequest:
		return http.StatusBadRequest
	case ErrNotFound:
		return http.StatusNotFound
	case ErrConflict:
		return http.StatusConflict
	case ErrUnauthorized:
		return http.StatusUnauthorized
	default:
		return http.StatusInternalServerError
	}
}

type CustomError struct {
	ErrorType error
	Message   string
	Cause     interface{}
}

func (e CustomError) Error() string {
	return fmt.Sprintf("%s. Cause: %s", e.Message, e.Cause)
}
