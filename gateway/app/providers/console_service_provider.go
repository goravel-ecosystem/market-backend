package providers

import (
	"github.com/goravel/framework/contracts/foundation"
	"github.com/goravel/framework/facades"

	"github.com/goravel-ecosystem/market-backend/gateway/app/console"
)

type ConsoleServiceProvider struct {
}

func (receiver *ConsoleServiceProvider) Register(app foundation.Application) {
	kernel := console.Kernel{}
	facades.Schedule().Register(kernel.Schedule())
	facades.Artisan().Register(kernel.Commands())
}

func (receiver *ConsoleServiceProvider) Boot(app foundation.Application) {

}
