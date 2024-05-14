package routes

import (
	"github.com/goravel/framework/facades"
	httpmiddleware "github.com/goravel/framework/http/middleware"
	"github.com/goravel/gateway"

	"market.goravel.dev/gateway/app/http/middleware"
	"market.goravel.dev/gateway/app/services"
)

func Users() {
	userService := services.NewUserImpl()

	facades.Route().Post("/users/email/login", gateway.Post)
	facades.Route().Post("/users/email/register", gateway.Post)
	facades.Route().Middleware(httpmiddleware.Throttle("VerifyCode")).Get("/users/email/register/code", gateway.Get)
	facades.Route().Middleware(middleware.Jwt(userService)).Get("/users/self", gateway.Get)
	facades.Route().Middleware(middleware.Jwt(userService)).Put("/users/{id}", gateway.Put)
}
