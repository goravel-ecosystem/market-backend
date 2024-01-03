package helper

import (
	"github.com/goravel/framework/facades"
)

func IsLocal() bool {
	return facades.Config().GetString("app.env") != "develop" && facades.Config().GetString("app.env") != "production"
}
