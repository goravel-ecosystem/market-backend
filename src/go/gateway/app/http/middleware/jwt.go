package middleware

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"

	"market.goravel.dev/gateway/app/services"
)

func Jwt(userService services.User) http.Middleware {
	return func(ctx http.Context) {
		token := ctx.Request().Header("Authorization", "")
		if token == "" {
			ctx.Request().AbortWithStatus(http.StatusUnauthorized)
			return
		}

		user, err := userService.GetUserByToken(ctx, token)
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
