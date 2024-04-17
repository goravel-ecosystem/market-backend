package grpc

import (
	"google.golang.org/grpc"

	"market.goravel.dev/utils/interceptors"
)

type Kernel struct {
}

// The application's global GRPC interceptor stack.
// These middleware are run during every request to your application.
func (kernel *Kernel) UnaryServerInterceptors() []grpc.UnaryServerInterceptor {
	return []grpc.UnaryServerInterceptor{
		interceptors.Response(),
	}
}

// The application's client interceptor groups.
func (kernel *Kernel) UnaryClientInterceptorGroups() map[string][]grpc.UnaryClientInterceptor {
	return map[string][]grpc.UnaryClientInterceptor{}
}
