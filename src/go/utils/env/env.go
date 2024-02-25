package env

import (
	"github.com/goravel/framework/facades"
)

func IsProduction() bool {
	return facades.Config().GetString("app.env") == "production"
}

func IsStaging() bool {
	return facades.Config().GetString("app.env") == "staging"
}
