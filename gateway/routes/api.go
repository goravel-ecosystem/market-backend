package routes

import (
	"github.com/goravel/framework/facades"
	"github.com/goravel/gateway"

	"goravel/app/http/controllers"
)

func Api() {
	userController := controllers.NewUserController()
	facades.Route().Get("/users/{id}", userController.Show)

	facades.Route().Get("/v1/user/{name}", gateway.Get)
	facades.Route().Post("/v1/example/echo", gateway.Post)
}
