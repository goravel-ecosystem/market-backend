package controllers

import (
	httpcontract "github.com/goravel/framework/contracts/http"

	"market.goravel.dev/proto/base"
)

func NewBadRequestStatus(err error) *base.Status {
	return &base.Status{
		Code:  httpcontract.StatusBadRequest,
		Error: err.Error(),
	}
}

func NewNotFoundStatus(err error) *base.Status {
	return &base.Status{
		Code:  httpcontract.StatusNotFound,
		Error: err.Error(),
	}
}

func NewOkStatus() *base.Status {
	return &base.Status{
		Code: httpcontract.StatusOK,
	}
}
