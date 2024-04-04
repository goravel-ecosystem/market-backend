package routes

import (
	"github.com/goravel/framework/facades"

	"market.goravel.dev/package/app/grpc/controllers"
	protopackage "market.goravel.dev/proto/package"
)

func Grpc() {
	protopackage.RegisterTagServiceServer(facades.Grpc().Server(), controllers.NewTagController())
}
