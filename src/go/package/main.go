package main

import (
	"market.goravel.dev/package/bootstrap"
)

func main() {
	// This bootstraps the framework and gets it ready for use.
	bootstrap.Boot()

	select {}
}
