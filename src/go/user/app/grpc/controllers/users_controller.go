package controllers

import (
	"context"
	"errors"

	"github.com/goravel/framework/facades"
	"github.com/goravel/framework/http"

	protouser "market.goravel.dev/proto/user"
	utilserrors "market.goravel.dev/utils/errors"

	"market.goravel.dev/user/app/services"
)

type UsersController struct {
	protouser.UnimplementedUserServiceServer
	notificationService services.Notification
	userService         services.User
}

func NewUsersController() *UsersController {
	return &UsersController{
		notificationService: services.NewNotificationImpl(),
		userService:         services.NewUserImpl(),
	}
}

func (r *UsersController) EmailLogin(ctx context.Context, req *protouser.EmailLoginRequest) (*protouser.EmailLoginResponse, error) {
	if err := validateEmailLoginRequest(ctx, req); err != nil {
		return nil, err
	}

	user, err := r.userService.GetUserByEmail(req.GetEmail())
	if err != nil {
		return nil, err
	}

	if !facades.Hash().Check(req.GetPassword(), user.Password) {
		return nil, utilserrors.NewValidate(facades.Lang(ctx).Get("invalid.password.error"))
	}

	token, err := facades.Auth(http.Background()).LoginUsingID(user.ID)
	if err != nil {
		return nil, err
	}

	return &protouser.EmailLoginResponse{
		Status: NewOkStatus(),
		User:   user.ToProto(),
		Token:  "Bearer " + token,
	}, nil
}

func (r *UsersController) EmailRegister(ctx context.Context, req *protouser.EmailRegisterRequest) (*protouser.EmailRegisterResponse, error) {
	if err := validateEmailRegisterRequest(ctx, req); err != nil {
		return nil, err
	}

	exist, err := r.userService.IsEmailExist(req.GetEmail())
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, utilserrors.NewValidate(facades.Lang(ctx).Get("exist.email"))
	}

	if !r.notificationService.VerifyEmailRegisterCode(req.GetCodeKey(), req.GetCode()) {
		return nil, utilserrors.NewValidate(facades.Lang(ctx).Get("invalid.code"))
	}

	user, err := r.userService.Register(req.GetName(), req.GetEmail(), req.GetPassword())
	if err != nil {
		return nil, err
	}

	token, err := facades.Auth(http.Background()).LoginUsingID(user.ID)
	if err != nil {
		return nil, err
	}

	return &protouser.EmailRegisterResponse{
		Status: NewOkStatus(),
		User:   user.ToProto(),
		Token:  "Bearer " + token,
	}, nil
}

func (r *UsersController) GetEmailRegisterCode(ctx context.Context, req *protouser.GetEmailRegisterCodeRequest) (*protouser.GetEmailRegisterCodeResponse, error) {
	if err := validateGetEmailRegisterCodeRequest(ctx, req); err != nil {
		return nil, err
	}

	exist, err := r.userService.IsEmailExist(req.GetEmail())
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, utilserrors.NewValidate(facades.Lang(ctx).Get("exist.email"))
	}

	key, err := r.notificationService.SendEmailRegisterCode(ctx, req.GetEmail())
	if err != nil {
		return nil, err
	}

	return &protouser.GetEmailRegisterCodeResponse{
		Status: NewOkStatus(),
		Key:    key,
	}, nil
}

func (r *UsersController) GetUser(ctx context.Context, req *protouser.GetUserRequest) (*protouser.GetUserResponse, error) {
	userID := req.GetUserId()
	if userID == "" {
		return &protouser.GetUserResponse{
			Status: NewBadRequestStatus(errors.New(facades.Lang(ctx).Get("required.user_id"))),
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
		return &protouser.GetUserByTokenResponse{
			Status: NewBadRequestStatus(errors.New(facades.Lang(ctx).Get("required.token"))),
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
