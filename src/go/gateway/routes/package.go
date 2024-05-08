package routes

import (
	"github.com/goravel/framework/facades"
	"github.com/goravel/gateway"

	"market.goravel.dev/gateway/app/http/middleware"
	"market.goravel.dev/gateway/app/services"
)

func Packages() {
	userService := services.NewUserImpl()

	facades.Route().Get("/packages", gateway.Get)
	facades.Route().Get("/packages/tags", gateway.Get)
	facades.Route().Middleware(middleware.Jwt(userService)).Get("/packages/{id}", gateway.Get)
	facades.Route().Middleware(middleware.Jwt(userService)).Put("/packages/{id}", gateway.Put)
	facades.Route().Middleware(middleware.Jwt(userService)).Post("/packages", gateway.Post)
}
