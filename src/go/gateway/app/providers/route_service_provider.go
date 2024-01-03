package providers

import (
	"github.com/goravel/framework/contracts/foundation"
	contractshttp "github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"github.com/goravel/framework/http/limit"

	"market.goravel.dev/gateway/app/helper"
	"market.goravel.dev/gateway/app/http"
	"market.goravel.dev/gateway/routes"
)

type RouteServiceProvider struct {
}

func (receiver *RouteServiceProvider) Register(app foundation.Application) {
}

func (receiver *RouteServiceProvider) Boot(app foundation.Application) {
	//Add HTTP middleware
	facades.Route().GlobalMiddleware(http.Kernel{}.Middleware()...)

	receiver.configureRateLimiting()

	routes.Users()
}

func (receiver *RouteServiceProvider) configureRateLimiting() {
	facades.RateLimiter().For("VerifyCode", func(ctx contractshttp.Context) contractshttp.Limit {
		if helper.IsLocal() {
			return limit.PerMinute(100)
		} else {
			return limit.PerMinute(1)
		}
	})
}
