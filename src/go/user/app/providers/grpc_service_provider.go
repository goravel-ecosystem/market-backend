package providers

import (
	"github.com/goravel/framework/contracts/foundation"
	"github.com/goravel/framework/facades"

	"market.goravel.dev/user/app/grpc"
	"market.goravel.dev/user/routes"
)

type GrpcServiceProvider struct {
}

func (receiver *GrpcServiceProvider) Register(app foundation.Application) {
	//Add Grpc interceptors
	kernel := grpc.Kernel{}
	facades.Grpc().UnaryServerInterceptors(kernel.UnaryServerInterceptors())
	facades.Grpc().UnaryClientInterceptorGroups(kernel.UnaryClientInterceptorGroups())
}

func (receiver *GrpcServiceProvider) Boot(app foundation.Application) {
	//Add routes
	routes.Grpc()
}
