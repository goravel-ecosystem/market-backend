package config

import (
	"github.com/goravel/framework/facades"
	"github.com/goravel/gateway"

	"github.com/goravel-ecosystem/market-backend/proto/users"
)

func init() {
	config := facades.Config()
	config.Add("grpc", map[string]any{
		// Configure your server host
		"host": config.Env("GRPC_HOST", ""),

		// Configure your server port
		"port": config.Env("GRPC_PORT", ""),

		// Configure your client host and interceptors.
		// Interceptors can be the group name of UnaryClientInterceptorGroups in app/grpc/kernel.go.
		"clients": map[string]any{
			"users": map[string]any{
				"host":         config.Env("GRPC_USER_HOST", ""),
				"port":         config.Env("GRPC_USER_PORT", ""),
				"handlers":     []gateway.Handler{users.RegisterUsersServiceHandler},
				"interceptors": []string{},
			},
			//"business": map[string]any{
			//	"host":         config.Env("GRPC_BUSINESS_HOST", ""),
			//	"port":         config.Env("GRPC_BUSINESS_PORT", ""),
			//	"handlers":     []gateway.Handler{business.RegisterBusinessServiceHandler},
			//	"interceptors": []string{},
			//},
		},
	})
}
