package routes

import (
	"github.com/goravel/framework/facades"
	"github.com/goravel/gateway"

	"github.com/goravel-ecosystem/market-backend/gateway/app/http/middleware"
	"github.com/goravel-ecosystem/market-backend/gateway/app/services"
)

func Users() {
	userService := services.NewUserImpl()

	facades.Route().Post("/users/email/login", gateway.Post)
	facades.Route().Post("/users/email/register", gateway.Post)
	facades.Route().Post("/users/email/code", gateway.Post)
	facades.Route().Middleware(middleware.Jwt(userService)).Get("/users/self", gateway.Get)
}
