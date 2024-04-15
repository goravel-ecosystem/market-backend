package services

import (
	"context"

	"github.com/goravel/framework/facades"
	"github.com/spf13/cast"

	"market.goravel.dev/proto/user"
	"market.goravel.dev/utils/errors"
)

var userInstance *UserImpl

type User interface {
	GetUser(ctx context.Context, userID uint64) (*user.User, error)
}

type UserImpl struct {
	client user.UserServiceClient
}

func NewUserImpl() *UserImpl {
	if userInstance != nil {
		return userInstance
	}

	client, err := facades.Grpc().Client(context.Background(), "user")
	if err != nil {
		facades.Log().Errorf("init UserService err: %+v", err)
		return nil
	}

	userInstance = &UserImpl{
		client: user.NewUserServiceClient(client),
	}

	return userInstance
}

func (r *UserImpl) GetUser(ctx context.Context, userID uint64) (*user.User, error) {
	resp, err := r.client.GetUser(ctx, &user.GetUserRequest{UserId: cast.ToString(userID)})
	if err := errors.NewResponse(resp, err); err != nil {
		return nil, err
	}

	return resp.GetUser(), nil
}
