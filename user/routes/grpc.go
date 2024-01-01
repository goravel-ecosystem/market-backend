package routes

import (
	"github.com/goravel/framework/facades"

	"github.com/goravel-ecosystem/market-backend/proto/user"

	"github.com/goravel-ecosystem/market-backend/user/app/grpc/controllers"
)

func Grpc() {
	user.RegisterUserServiceServer(facades.Grpc().Server(), controllers.NewUsersController())
}
