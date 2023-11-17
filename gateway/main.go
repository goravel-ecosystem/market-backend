package main

import (
	"github.com/goravel/framework/facades"
	gatewayfacades "github.com/goravel/gateway/facades"

	"goravel/bootstrap"
)

func main() {
	// This bootstraps the framework and gets it ready for use.
	bootstrap.Boot()

	//Start http server by facades.Route().
	go func() {
		if err := facades.Route().Run(); err != nil {
			facades.Log().Errorf("Route run error: %v", err)
		}
	}()

	go func() {
		if err := gatewayfacades.Gateway().Run(); err != nil {
			facades.Log().Errorf("Gateway run error: %v", err)
		}
		//connections := make(map[string]*grpc.ClientConn)
		//gwmux := runtime.NewServeMux()
		//grpcConfig := facades.Config().Get("gateway.grpc").([]util.Grpc)
		//for _, item := range grpcConfig {
		//	if _, exist := connections[item.Name]; !exist {
		//		connection, err := facades.Grpc().Client(context.Background(), item.Name)
		//		if err != nil {
		//			facades.Log().Errorf("Failed to init %s client: %v", item.Name, err)
		//			continue
		//		}
		//
		//		connections[item.Name] = connection
		//	}
		//
		//	if err := item.Handler(context.Background(), gwmux, connections[item.Name]); err != nil {
		//		facades.Log().Errorf("Failed to register %s handler: %v", item.Name, err)
		//	}
		//}
		//
		//gwServer := &http.Server{
		//	Addr:    fmt.Sprintf("%s:%s", facades.Config().GetString("gateway.host"), facades.Config().GetString("gateway.port")),
		//	Handler: gwmux,
		//}
		//if err := gwServer.ListenAndServe(); err != nil {
		//	facades.Log().Errorf("Gateway run error: %v", err)
		//}
	}()

	select {}
}
