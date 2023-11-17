package routes

import (
	"github.com/goravel/framework/facades"

	"github.com/goravel-ecosystem/market-backend/proto/users"
	"github.com/goravel-ecosystem/market-backend/users/app/grpc/controllers"
)

func Grpc() {
	users.RegisterUsersServiceServer(facades.Grpc().Server(), controllers.NewUsersController())
}
