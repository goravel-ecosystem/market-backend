package routes

import (
	"github.com/goravel/framework/facades"

	"github.com/goravel-ecosystem/market-backend/users/app/http/controllers"
)

func Api() {
	userController := controllers.NewUserController()
	facades.Route().Get("/users/{id}", userController.Show)
}
