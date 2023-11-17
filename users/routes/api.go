package routes

import (
	"github.com/goravel/framework/facades"

	"users/app/http/controllers"
)

func Api() {
	userController := controllers.NewUserController()
	facades.Route().Get("/users/{id}", userController.Show)
}
