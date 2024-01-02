package controllers

import (
	"context"
	"errors"
	"regexp"

	"github.com/goravel/framework/facades"

	protouser "github.com/goravel-ecosystem/market-backend/proto/user"
)

func validateGetEmailRegisterCodeRequest(ctx context.Context, req *protouser.GetEmailRegisterCodeRequest) error {
	return validateEmailValid(ctx, req.GetEmail())
}

func validateEmailLoginRequest(ctx context.Context, req *protouser.EmailLoginRequest) error {
	if err := validateEmailValid(ctx, req.GetEmail()); err != nil {
		return err
	}
	if req.GetPassword() == "" {
		err, _ := facades.Lang(ctx).Get("required.password")

		return errors.New(err)
	}

	return nil
}

func validateEmailRegisterRequest(ctx context.Context, req *protouser.EmailRegisterRequest) error {
	if err := validateEmailValid(ctx, req.GetEmail()); err != nil {
		return err
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

func validateEmailValid(ctx context.Context, email string) error {
	if email == "" {
		requiredEmail, _ := facades.Lang(ctx).Get("required.email")

		return errors.New(requiredEmail)
	}

	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	if matched, err := regexp.MatchString(pattern, email); !matched || err != nil {
		invalidEmail, _ := facades.Lang(ctx).Get("invalid.email")

		return errors.New(invalidEmail)
	}

	return nil
}
