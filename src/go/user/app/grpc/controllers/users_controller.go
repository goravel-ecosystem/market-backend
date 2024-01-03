package controllers

import (
	"context"
	"errors"

	"github.com/goravel/framework/facades"
	"github.com/goravel/framework/http"

	protouser "market.goravel.dev/proto/user"

	"market.goravel.dev/user/app/services"
)

type UsersController struct {
	protouser.UnimplementedUserServiceServer
	notificationService services.Notification
}

func NewUsersController() *UsersController {
	return &UsersController{
		notificationService: services.NewNotificationImpl(),
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

	if !r.notificationService.VerifyEmailRegisterCode(req.GetEmail(), req.GetCode()) {
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

func (r *UsersController) GetEmailRegisterCode(ctx context.Context, req *protouser.GetEmailRegisterCodeRequest) (*protouser.GetEmailRegisterCodeResponse, error) {
	if err := validateGetEmailRegisterCodeRequest(ctx, req); err != nil {
		return &protouser.GetEmailRegisterCodeResponse{
			Status: NewBadRequestStatus(err),
		}, nil
	}

	key, err := r.notificationService.SendEmailRegisterCode(ctx, req.GetEmail())
	if err != nil {
		return &protouser.GetEmailRegisterCodeResponse{
			Status: NewBadRequestStatus(err),
		}, nil
	}

	return &protouser.GetEmailRegisterCodeResponse{
		Status: NewOkStatus(),
		Key:    key,
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
