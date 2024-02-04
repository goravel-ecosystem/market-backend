package errors

import (
	"net/http"
)

type ErrorWithCode interface {
	Code() int
	Error() string
}

type Validate struct {
	code    int
	message string
}

func NewValidate(message string) *Validate {
	return &Validate{
		code:    http.StatusBadRequest,
		message: message,
	}
}

func (r *Validate) Code() int {
	return r.code
}

func (r *Validate) Error() string {
	return r.message
}
