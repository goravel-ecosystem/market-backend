package config

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

func init() {
	config := facades.Config()
	config.Add("gateway", map[string]any{
		"host": config.Env("GATEWAY_HOST", ""),
		"port": config.Env("GATEWAY_PORT", ""),
		"fallback": func(ctx http.Context, err error) http.Response {
			return ctx.Response().Success().Json(map[string]any{
				"status": map[string]any{
					"code":  http.StatusInternalServerError,
					"error": err.Error(),
				},
			})
		},
	})
}
