package controllers

import (
	httpcontract "github.com/goravel/framework/contracts/http"

	"github.com/goravel-ecosystem/market-backend/proto/base"
)

func NewBadRequestStatus(err error) *base.Status {
	return &base.Status{
		Code:  httpcontract.StatusBadRequest,
		Error: err.Error(),
	}
}

func NewOkStatus() *base.Status {
	return &base.Status{
		Code: httpcontract.StatusOK,
	}
}
