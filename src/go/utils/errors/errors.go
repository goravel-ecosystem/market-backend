package errors

import (
	"net/http"

	"github.com/goravel/framework/facades"

	"market.goravel.dev/proto/base"
)

type response interface {
	GetStatus() *base.Status
}

type ErrorWithCode interface {
	Code() int32
	Error() string
}

type ErrorWithCodeImpl struct {
	code    int32
	message string
}

func (r *ErrorWithCodeImpl) Code() int32 {
	return r.code
}

func (r *ErrorWithCodeImpl) Error() string {
	return r.message
}

func New(code int32, message string) ErrorWithCode {
	return &ErrorWithCodeImpl{
		code:    code,
		message: message,
	}
}

func NewBadRequest(message string) error {
	return New(http.StatusBadRequest, message)
}

func NewUnauthorized(message string) ErrorWithCode {
	return New(http.StatusUnauthorized, message)
}

func NewNotFound(message string) ErrorWithCode {
	return New(http.StatusNotFound, message)
}

func NewInternalServerError(err error) ErrorWithCode {
	facades.Log().Errorf("internal server error: %+v", err)

	return New(http.StatusInternalServerError, err.Error())
}

func NewResponse(resp response, err error) ErrorWithCode {
	if err != nil {
		return NewInternalServerError(err)
	}

	if resp.GetStatus().GetCode() >= http.StatusMultipleChoices {
		return New(resp.GetStatus().GetCode(), resp.GetStatus().GetError())
	}

	return nil
}
