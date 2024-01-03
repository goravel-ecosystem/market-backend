package routes

import (
	"github.com/goravel/framework/facades"

	"market.goravel.dev/proto/user"

	"market.goravel.dev/user/app/grpc/controllers"
)

func Grpc() {
	user.RegisterUserServiceServer(facades.Grpc().Server(), controllers.NewUsersController())
}
