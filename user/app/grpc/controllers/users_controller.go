package controllers

import (
	"context"
	"errors"

	"github.com/goravel/framework/facades"
	"github.com/goravel/framework/http"

	protouser "github.com/goravel-ecosystem/market-backend/proto/user"
	"github.com/goravel-ecosystem/market-backend/user/app/services"
)

type UsersController struct {
	protouser.UnimplementedUserServiceServer
	notificationService services.Notification
}

func NewUsersController() *UsersController {
	return &UsersController{
		notificationService: services.NewNotification(),
	}
}

func (r *UsersController) EmailLogin(ctx context.Context, req *protouser.EmailLoginRequest) (*protouser.EmailLoginResponse, error) {
	if err := validateEmailLoginRequest(ctx, req); err != nil {
		return &protouser.EmailLoginResponse{
			Status: NewBadRequestStatus(err),
		}, nil
	}

	token, err := facades.Auth(http.Background()).LoginUsingID("uuid")
	if err != nil {
		return &protouser.EmailLoginResponse{
			Status: NewBadRequestStatus(err),
		}, nil
	}

	return &protouser.EmailLoginResponse{
		Status: NewOkStatus(),
		User: &protouser.User{
			Id:   "uuid",
			Name: req.GetEmail(),
		},
		Token: "Bearer " + token,
	}, nil
}

func (r *UsersController) EmailRegister(ctx context.Context, req *protouser.EmailRegisterRequest) (*protouser.EmailRegisterResponse, error) {
	if err := validateEmailRegisterRequest(ctx, req); err != nil {
		return &protouser.EmailRegisterResponse{
			Status: NewBadRequestStatus(err),
		}, nil
	}

	if !r.notificationService.VerifyRegisterEmailCode(req.GetEmail(), req.GetCode()) {
		invalidCode, _ := facades.Lang(ctx).Get("invalid.code")

		return &protouser.EmailRegisterResponse{
			Status: NewBadRequestStatus(errors.New(invalidCode)),
		}, nil
	}

	return &protouser.EmailRegisterResponse{
		Status: NewOkStatus(),
		User: &protouser.User{
			Id:   "uuid",
			Name: req.GetEmail(),
		},
	}, nil
}

func (r *UsersController) GetEmailCode(ctx context.Context, req *protouser.GetEmailCodeRequest) (*protouser.GetEmailCodeResponse, error) {
	if req.GetEmail() == "" {
		requiredEmail, _ := facades.Lang(ctx).Get("required_email")

		return &protouser.GetEmailCodeResponse{
			Status: NewBadRequestStatus(errors.New(requiredEmail)),
		}, nil
	}

	if err := r.notificationService.SendRegisterEmailCode(ctx, req.GetEmail()); err != nil {
		return &protouser.GetEmailCodeResponse{
			Status: NewBadRequestStatus(err),
		}, nil
	}

	return &protouser.GetEmailCodeResponse{
		Status: NewOkStatus(),
	}, nil
}

func (r *UsersController) GetUser(ctx context.Context, req *protouser.GetUserRequest) (*protouser.GetUserResponse, error) {
	userID := req.GetUserId()
	if userID == "" {
		requiredUserId, _ := facades.Lang(ctx).Get("required.user_id")

		return &protouser.GetUserResponse{
			Status: NewBadRequestStatus(errors.New(requiredUserId)),
		}, nil
	}

	//var user models.User
	//if err := facades.Auth(httpCtx).User(&user); err != nil {
	//	return &protouser.GetUserByTokenResponse{
	//		Status: NewBadRequestStatus(err),
	//	}, nil
	//}

	return &protouser.GetUserResponse{
		Status: NewOkStatus(),
		User: &protouser.User{
			Id:   userID,
			Name: "test",
		},
	}, nil
}

func (r *UsersController) GetUserByToken(ctx context.Context, req *protouser.GetUserByTokenRequest) (*protouser.GetUserByTokenResponse, error) {
	token := req.GetToken()
	if token == "" {
		requiredToken, _ := facades.Lang(ctx).Get("required.token")

		return &protouser.GetUserByTokenResponse{
			Status: NewBadRequestStatus(errors.New(requiredToken)),
		}, nil
	}

	httpCtx := http.Background()
	if _, err := facades.Auth(httpCtx).Parse(token); err != nil {
		return &protouser.GetUserByTokenResponse{
			Status: NewBadRequestStatus(err),
		}, nil
	}

	//var user models.User
	//if err := facades.Auth(httpCtx).User(&user); err != nil {
	//	return &protouser.GetUserByTokenResponse{
	//		Status: NewBadRequestStatus(err),
	//	}, nil
	//}

	return &protouser.GetUserByTokenResponse{
		Status: NewOkStatus(),
		User: &protouser.User{
			Id:   "uuid",
			Name: "test",
		},
	}, nil
}
