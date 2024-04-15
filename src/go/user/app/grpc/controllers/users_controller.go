package controllers

import (
	"context"
	"errors"

	"github.com/goravel/framework/database/orm"
	"github.com/goravel/framework/facades"
	"github.com/goravel/framework/http"

	protouser "market.goravel.dev/proto/user"
	"market.goravel.dev/user/app/models"
	"market.goravel.dev/user/app/services"
	utilserrors "market.goravel.dev/utils/errors"
	utilsresponse "market.goravel.dev/utils/response"
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
		return nil, utilserrors.NewBadRequest(facades.Lang(ctx).Get("invalid.password.error"))
	}

	token, err := facades.Auth(http.Background()).LoginUsingID(user.ID)
	if err != nil {
		return nil, err
	}

	return &protouser.EmailLoginResponse{
		Status: utilsresponse.NewOkStatus(),
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
		return nil, utilserrors.NewBadRequest(facades.Lang(ctx).Get("exist.email"))
	}

	if !r.notificationService.VerifyEmailRegisterCode(req.GetCodeKey(), req.GetCode()) {
		return nil, utilserrors.NewBadRequest(facades.Lang(ctx).Get("invalid.code"))
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
		Status: utilsresponse.NewOkStatus(),
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
		return nil, utilserrors.NewBadRequest(facades.Lang(ctx).Get("exist.email"))
	}

	key, err := r.notificationService.SendEmailRegisterCode(ctx, req.GetEmail())
	if err != nil {
		return nil, err
	}

	return &protouser.GetEmailRegisterCodeResponse{
		Status: utilsresponse.NewOkStatus(),
		Key:    key,
	}, nil
}

func (r *UsersController) GetUser(ctx context.Context, req *protouser.GetUserRequest) (*protouser.GetUserResponse, error) {
	userID := req.GetUserId()
	if userID == "" {
		return nil, utilserrors.NewBadRequest(facades.Lang(ctx).Get("required.user_id"))
	}

	user, err := r.userService.GetUserByID(userID)
	if err != nil {
		if errors.Is(err, orm.ErrRecordNotFound) {
			return nil, utilserrors.NewNotFound(facades.Lang(ctx).Get("not_exist.user"))
		}
		return nil, err
	}

	return &protouser.GetUserResponse{
		Status: utilsresponse.NewOkStatus(),
		User:   user.ToProto(),
	}, nil
}

func (r *UsersController) GetUserByToken(ctx context.Context, req *protouser.GetUserByTokenRequest) (*protouser.GetUserByTokenResponse, error) {
	token := req.GetToken()
	if token == "" {
		return nil, utilserrors.NewBadRequest(facades.Lang(ctx).Get("required.token"))
	}

	httpCtx := http.Background()
	if _, err := facades.Auth(httpCtx).Parse(token); err != nil {
		return nil, utilserrors.NewInternalServerError(err)
	}

	var user models.User
	if err := facades.Auth(httpCtx).User(&user); err != nil {
		return nil, utilserrors.NewInternalServerError(err)
	}

	return &protouser.GetUserByTokenResponse{
		Status: utilsresponse.NewOkStatus(),
		User:   user.ToProto(),
	}, nil
}
