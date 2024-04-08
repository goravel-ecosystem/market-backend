package routes

import (
	"github.com/goravel/framework/facades"
	"github.com/goravel/gateway"
)

func Packages() {
	facades.Route().Get("/packages/tags", gateway.Get)
}
