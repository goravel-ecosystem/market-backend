package routes

import (
	"github.com/goravel/framework/facades"

	packageproto "market.goravel.dev/proto/package"

	"market.goravel.dev/package/app/grpc/controllers"
)

func Grpc() {
	packageproto.RegisterPackageServiceServer(facades.Grpc().Server(), controllers.NewPackageController())
}
