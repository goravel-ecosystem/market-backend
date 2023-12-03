package middleware

import (
	"context"
	"errors"
	"time"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"

	"github.com/goravel-ecosystem/market-backend/proto/users"
)

func Jwt() http.Middleware {
	return func(ctx http.Context) {
		token := ctx.Request().Header("Authorization", "")
		if token == "" {
			ctx.Request().AbortWithStatus(http.StatusUnauthorized)
			return
		}

		user, err := getUser(ctx, token)
		if err != nil {
			facades.Log().Request(ctx.Request()).Errorf("get user err: %+v", err)
			ctx.Request().AbortWithStatus(http.StatusUnauthorized)
			return
		}

		query := ctx.Request().Origin().URL.Query()
		query.Add("user_id", user.GetId())
		query.Add("user_name", user.GetName())
		ctx.Request().Origin().URL.RawQuery = query.Encode()
		ctx.Request().Next()
	}
}

func getUser(ctx context.Context, token string) (*users.User, error) {
	client, err := facades.Grpc().Client(ctx, "users")
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	usersService := users.NewUsersServiceClient(client)
	response, err := usersService.GetUserByToken(ctx, &users.GetUserByTokenRequest{
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
