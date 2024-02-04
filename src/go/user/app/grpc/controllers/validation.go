package controllers

import (
	"context"
	"regexp"

	"github.com/goravel/framework/facades"

	protouser "market.goravel.dev/proto/user"
	utilserrors "market.goravel.dev/utils/errors"
)

func validateGetEmailRegisterCodeRequest(ctx context.Context, req *protouser.GetEmailRegisterCodeRequest) error {
	return validateEmailValid(ctx, req.GetEmail())
}

func validateEmailLoginRequest(ctx context.Context, req *protouser.EmailLoginRequest) error {
	if err := validateEmailValid(ctx, req.GetEmail()); err != nil {
		return err
	}
	if req.GetPassword() == "" {
		return utilserrors.NewValidate(facades.Lang(ctx).Get("required.password"))
	}

	return nil
}

func validateEmailRegisterRequest(ctx context.Context, req *protouser.EmailRegisterRequest) error {
	if err := validateEmailValid(ctx, req.GetEmail()); err != nil {
		return err
	}
	if req.GetName() == "" {
		return utilserrors.NewValidate(facades.Lang(ctx).Get("required.name"))
	}
	if req.GetPassword() == "" {
		return utilserrors.NewValidate(facades.Lang(ctx).Get("required.password"))
	}
	if len(req.GetPassword()) < 6 {
		return utilserrors.NewValidate(facades.Lang(ctx).Get("invalid.password.min"))
	}
	if req.GetCode() == "" {
		return utilserrors.NewValidate(facades.Lang(ctx).Get("required.code"))
	}
	if req.GetCodeKey() == "" {
		return utilserrors.NewValidate(facades.Lang(ctx).Get("required.code_key"))
	}

	return nil
}

func validateEmailValid(ctx context.Context, email string) error {
	if email == "" {
		return utilserrors.NewValidate(facades.Lang(ctx).Get("required.email"))
	}

	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	if matched, err := regexp.MatchString(pattern, email); !matched || err != nil {
		return utilserrors.NewValidate(facades.Lang(ctx).Get("invalid.email"))
	}

	return nil
}
