package routes

import (
	"github.com/goravel/framework/facades"

	"market.goravel.dev/user/app/http/controllers"
)

func Api() {
	userController := controllers.NewUserController()
	facades.Route().Get("/users/{id}", userController.Show)
}
