package services

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"

	protouser "github.com/goravel-ecosystem/market-backend/proto/user"
)

var userInstance *UserImpl
var userOnce sync.Once

type User interface {
	GetUserByToken(ctx context.Context, token string) (user *protouser.User, err error)
}

type UserImpl struct {
	client  protouser.UserServiceClient
	timeout int
}

func NewUserImpl() *UserImpl {
	userOnce.Do(func() {
		client, err := facades.Grpc().Client(context.Background(), "user")
		if err != nil {
			panic(fmt.Sprintf("init UserService err: %+v", err))
		}

		userInstance = &UserImpl{
			client:  protouser.NewUserServiceClient(client),
			timeout: 10,
		}
	})

	return userInstance
}

func (r *UserImpl) GetUserByToken(ctx context.Context, token string) (*protouser.User, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(r.timeout)*time.Second)
	defer cancel()

	response, err := r.client.GetUserByToken(ctx, &protouser.GetUserByTokenRequest{
		Token: token,
	})
	if err != nil {
		return nil, err
	}
	if response.Status.Code != http.StatusOK {
		return nil, errors.New(response.Status.Error)
	}

	return response.User, nil
}
