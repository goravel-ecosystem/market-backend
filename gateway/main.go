package main

import (
	"github.com/goravel/framework/facades"
	gatewayfacades "github.com/goravel/gateway/facades"

	"github.com/goravel-ecosystem/market-backend/gateway/bootstrap"
)

func main() {
	// This bootstraps the framework and gets it ready for use.
	bootstrap.Boot()

	//Start http server by facades.Route().
	go func() {
		if err := facades.Route().Run(); err != nil {
			facades.Log().Errorf("Route run error: %v", err)
		}
	}()

	go func() {
		if err := gatewayfacades.Gateway().Run(); err != nil {
			facades.Log().Errorf("Gateway run error: %v", err)
		}
	}()

	select {}
}
