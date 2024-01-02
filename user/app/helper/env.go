package helper

import (
	"github.com/goravel/framework/facades"
)

func IsProduction() bool {
	return facades.Config().GetString("app.env") == "production"
}

func IsDevelopment() bool {
	return facades.Config().GetString("app.env") == "development"
}
