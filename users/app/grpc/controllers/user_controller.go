package controllers

import (
	"context"

	"github.com/goravel-ecosystem/market-backend/gateway/proto/users"
)

type UserController struct {
	users.UnimplementedUserServer
}

func NewUserController() *UserController {
	return &UserController{}
}

func (r *UserController) GetUser(ctx context.Context, req *user.GetUserRequest) (protoBook *user.GetUserResponse, err error) {
	return &users.GetUserResponse{
		Message: req.Name + " user " + req.Avatar,
	}, nil
}
