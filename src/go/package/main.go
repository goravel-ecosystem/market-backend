package main

import (
	"github.com/goravel/framework/facades"
	"google.golang.org/grpc/reflection"

	"market.goravel.dev/package/bootstrap"
)

func main() {
	// This bootstraps the framework and gets it ready for use.
	bootstrap.Boot()

	//Start http server by facades.Route().
	go func() {
		reflection.Register(facades.Grpc().Server())
		if err := facades.Grpc().Run(); err != nil {
			facades.Log().Errorf("Grpc run error: %v", err)
		}
	}()

	select {}
}
