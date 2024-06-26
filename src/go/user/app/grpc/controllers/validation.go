package controllers

import (
	"context"
	"regexp"

	"github.com/goravel/framework/contracts/translation"
	"github.com/goravel/framework/facades"

	protouser "market.goravel.dev/proto/user"
	utilserrors "market.goravel.dev/utils/errors"
)

func validateEmailLoginRequest(ctx context.Context, req *protouser.EmailLoginRequest) error {
	if err := validateEmailValid(ctx, req.GetEmail()); err != nil {
		return err
	}
	if req.GetPassword() == "" {
		return utilserrors.NewBadRequest(facades.Lang(ctx).Get("required.password"))
	}

	return nil
}

func validateEmailRegisterRequest(ctx context.Context, req *protouser.EmailRegisterRequest) error {
	if err := validateEmailValid(ctx, req.GetEmail()); err != nil {
		return err
	}
	if req.GetName() == "" {
		return utilserrors.NewBadRequest(facades.Lang(ctx).Get("required.name"))
	}
	if req.GetPassword() == "" {
		return utilserrors.NewBadRequest(facades.Lang(ctx).Get("required.password"))
	}
	if len(req.GetPassword()) < 6 {
		return utilserrors.NewBadRequest(facades.Lang(ctx).Get("invalid.password.min"))
	}
	if req.GetCode() == "" {
		return utilserrors.NewBadRequest(facades.Lang(ctx).Get("required.code"))
	}
	if req.GetCodeKey() == "" {
		return utilserrors.NewBadRequest(facades.Lang(ctx).Get("required.code_key"))
	}

	return nil
}

func validateEmailValid(ctx context.Context, email string) error {
	if email == "" {
		return utilserrors.NewBadRequest(facades.Lang(ctx).Get("required.email"))
	}

	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	if matched, err := regexp.MatchString(pattern, email); !matched || err != nil {
		return utilserrors.NewBadRequest(facades.Lang(ctx).Get("invalid.email"))
	}

	return nil
}

func validateGetEmailRegisterCodeRequest(ctx context.Context, req *protouser.GetEmailRegisterCodeRequest) error {
	return validateEmailValid(ctx, req.GetEmail())
}

func validateUpdateUserRequest(ctx context.Context, req *protouser.UpdateUserRequest) error {
	name := req.GetName()
	summery := req.GetSummary()
	password := req.GetPassword()
	userID := req.GetUserId()
	id := req.GetId()

	translate := facades.Lang(ctx)
	if id == "" {
		return utilserrors.NewBadRequest(translate.Get("required.id"))
	}
	if userID == "" {
		return utilserrors.NewBadRequest(translate.Get("required.user_id"))
	}
	if name == "" {
		return utilserrors.NewBadRequest(translate.Get("required.name"))
	}
	if len(name) > 50 {
		return utilserrors.NewBadRequest(translate.Get("invalid.name.max", translation.Option{
			Replace: map[string]string{
				"max": "50",
			},
		}))
	}

	if len(summery) > 200 {
		return utilserrors.NewBadRequest(translate.Get("invalid.summery.max", translation.Option{
			Replace: map[string]string{
				"max": "200",
			},
		}))
	}

	if password != "" {
		if len(password) < 6 {
			return utilserrors.NewBadRequest(translate.Get("invalid.password.min"))
		}

		if len(password) > 50 {
			return utilserrors.NewBadRequest(translate.Get("invalid.password.max", translation.Option{
				Replace: map[string]string{
					"max": "50",
				},
			}))
		}
	}

	return nil
}
