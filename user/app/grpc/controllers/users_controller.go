package controllers

import (
	"context"
	"errors"

	"github.com/goravel/framework/facades"
	"github.com/goravel/framework/http"

	"github.com/goravel-ecosystem/market-backend/proto/users"
	"github.com/goravel-ecosystem/market-backend/users/services"
)

type UsersController struct {
	users.UnimplementedUsersServiceServer
	notificationService services.Notification
}

func NewUsersController() *UsersController {
	return &UsersController{
		notificationService: services.NewNotification(),
	}
}

func (r *UsersController) EmailLogin(ctx context.Context, req *users.EmailLoginRequest) (*users.EmailLoginResponse, error) {
	if err := validateEmailLoginRequest(ctx, req); err != nil {
		return &users.EmailLoginResponse{
			Status: NewBadRequestStatus(err),
		}, nil
	}

	token, err := facades.Auth().LoginUsingID(http.Background(), "uuid")
	if err != nil {
		return &users.EmailLoginResponse{
			Status: NewBadRequestStatus(err),
		}, nil
	}

	return &users.EmailLoginResponse{
		Status: NewOkStatus(),
		User: &users.User{
			Id:   "uuid",
			Name: req.GetEmail(),
		},
		Token: "Bearer " + token,
	}, nil
}

func (r *UsersController) EmailRegister(ctx context.Context, req *users.EmailRegisterRequest) (*users.EmailRegisterResponse, error) {
	if err := validateEmailRegisterRequest(ctx, req); err != nil {
		return &users.EmailRegisterResponse{
			Status: NewBadRequestStatus(err),
		}, nil
	}

	if !r.notificationService.VerifyRegisterEmailCode(req.GetEmail(), req.GetCode()) {
		err, _ := facades.Lang(ctx).Get("invalid.code")

		return &users.EmailRegisterResponse{
			Status: NewBadRequestStatus(errors.New(err)),
		}, nil
	}

	return &users.EmailRegisterResponse{
		Status: NewOkStatus(),
		User: &users.User{
			Id:   "uuid",
			Name: req.GetEmail(),
		},
	}, nil
}

func (r *UsersController) GetEmailCode(ctx context.Context, req *users.GetEmailCodeRequest) (*users.GetEmailCodeResponse, error) {
	if req.GetEmail() == "" {
		err, _ := facades.Lang(ctx).Get("required_email")

		return &users.GetEmailCodeResponse{
			Status: NewBadRequestStatus(errors.New(err)),
		}, nil
	}

	if err := r.notificationService.SendRegisterEmailCode(ctx, req.GetEmail()); err != nil {
		return &users.GetEmailCodeResponse{
			Status: NewBadRequestStatus(err),
		}, nil
	}

	return &users.GetEmailCodeResponse{
		Status: NewOkStatus(),
	}, nil
}

func (r *UsersController) GetUser(ctx context.Context, req *users.GetUserRequest) (*users.GetUserResponse, error) {
	userID := req.GetUserId()
	if userID == "" {
		err, _ := facades.Lang(ctx).Get("required.user_id")

		return &users.GetUserResponse{
			Status: NewBadRequestStatus(errors.New(err)),
		}, nil
	}

	//var user models.User
	//if err := facades.Auth().User(httpCtx, &user); err != nil {
	//	return &users.GetUserByTokenResponse{
	//		Status: NewBadRequestStatus(err),
	//	}, nil
	//}

	return &users.GetUserResponse{
		Status: NewOkStatus(),
		User: &users.User{
			Id:   userID,
			Name: "test",
		},
	}, nil
}

func (r *UsersController) GetUserByToken(ctx context.Context, req *users.GetUserByTokenRequest) (*users.GetUserByTokenResponse, error) {
	token := req.GetToken()
	if token == "" {
		err, _ := facades.Lang(ctx).Get("required.token")

		return &users.GetUserByTokenResponse{
			Status: NewBadRequestStatus(errors.New(err)),
		}, nil
	}

	httpCtx := http.Background()
	if _, err := facades.Auth().Parse(httpCtx, token); err != nil {
		return &users.GetUserByTokenResponse{
			Status: NewBadRequestStatus(err),
		}, nil
	}

	//var user models.User
	//if err := facades.Auth().User(httpCtx, &user); err != nil {
	//	return &users.GetUserByTokenResponse{
	//		Status: NewBadRequestStatus(err),
	//	}, nil
	//}

	return &users.GetUserByTokenResponse{
		Status: NewOkStatus(),
		User: &users.User{
			Id:   "uuid",
			Name: "test",
		},
	}, nil
}
