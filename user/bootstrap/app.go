package bootstrap

import (
	"github.com/goravel/framework/foundation"

	"github.com/goravel-ecosystem/market-backend/user/config"
)

func Boot() {
	app := foundation.NewApplication()

	//Bootstrap the application
	app.Boot()

	//Bootstrap the config.
	config.Boot()
}
