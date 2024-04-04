package routes

import (
	"github.com/goravel/framework/facades"
	"github.com/goravel/gateway"

	"market.goravel.dev/gateway/app/http/middleware"
	"market.goravel.dev/gateway/app/services"
)

func Packages() {
	userService := services.NewUserImpl()

	facades.Route().Middleware(middleware.Jwt(userService)).Get("/packages/{id}", gateway.Get)
}
