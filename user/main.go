package main

import (
	"github.com/goravel/framework/facades"

	"github.com/goravel-ecosystem/market-backend/user/bootstrap"
)

func main() {
	// This bootstraps the framework and gets it ready for use.
	bootstrap.Boot()

	//Start http server by facades.Route().
	go func() {
		if err := facades.Grpc().Run(); err != nil {
			facades.Log().Errorf("Grpc run error: %v", err)
		}
	}()

	// Start queue server by facades.Queue().
	go func() {
		if err := facades.Queue().Worker(nil).Run(); err != nil {
			facades.Log().Errorf("Queue run error: %v", err)
		}
	}()

	select {}
}
