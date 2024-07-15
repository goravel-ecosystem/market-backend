package middleware

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"github.com/goravel/gateway"

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

		gateway.Inject(ctx, "user_id", user.GetId())
		gateway.Inject(ctx, "user_name", user.GetName())

		ctx.Request().Next()
	}
}
