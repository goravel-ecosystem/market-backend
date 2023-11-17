package controllers

import (
	"context"

	"github.com/goravel/framework/contracts/http"

	"github.com/goravel-ecosystem/market-backend/proto/base"
	"github.com/goravel-ecosystem/market-backend/proto/users"
)

type UsersController struct {
	users.UnimplementedUsersServiceServer
}

func NewUsersController() *UsersController {
	return &UsersController{}
}

func (r *UsersController) EmailRegister(ctx context.Context, req *users.EmailRegisterRequest) (protoBook *users.EmailRegisterResponse, err error) {
	return &users.EmailRegisterResponse{
		Status: &base.Status{
			Code: http.StatusOK,
		},
		User: &users.User{
			Id:   "uuid",
			Name: req.GetEmail(),
		},
	}, nil
}
