package controllers

import (
	"context"
	"errors"

	"github.com/goravel/framework/facades"

	protouser "github.com/goravel-ecosystem/market-backend/proto/user"
)

func validateEmailLoginRequest(ctx context.Context, req *protouser.EmailLoginRequest) error {
	if req.GetEmail() == "" {
		err, _ := facades.Lang(ctx).Get("required.email")

		return errors.New(err)
	}

	if req.GetPassword() == "" {
		err, _ := facades.Lang(ctx).Get("required.password")

		return errors.New(err)
	}

	return nil
}

func validateEmailRegisterRequest(ctx context.Context, req *protouser.EmailRegisterRequest) error {
	if req.GetEmail() == "" {
		err, _ := facades.Lang(ctx).Get("required.email")

		return errors.New(err)
	}

	if req.GetPassword() == "" {
		err, _ := facades.Lang(ctx).Get("required.password")

		return errors.New(err)
	}

	if req.GetCode() == "" {
		err, _ := facades.Lang(ctx).Get("required.code")

		return errors.New(err)
	}

	return nil
}
