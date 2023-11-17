package routes

import (
	"github.com/goravel/framework/facades"
	"github.com/goravel/gateway"
)

func Users() {
	facades.Route().Post("/users/register/email", gateway.Post)
}
