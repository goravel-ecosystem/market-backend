# Market Backend

Providers a backend for Goravel Market. In facilitate deployment and testing, we put all microservices in this project.

## Folder Structure

- [gateway](gateway/README.md): The gateway microservice provides a unified entry point for all clients.
- [package](package/README.md): The package microservice is responsible for all kinds of operations on package.
- [proto](proto/README.md): Define the proto files for all microservices.
- [user](user/README.md): The user microservice is responsible for user registration, login, and other functions.

## The Request Process

1. A request from the client;
2. Nginx forward to gateway;
3. The HTTP server of gateway;
4. (If JWT token exists)Parse User information through UserService;
5. (If JWT token exists)Append `user_id`, `user_name` to the request queries;
6. The request forward to the corresponding HTTP interface that define in proto files(through [goravel/gateway]
(https://github.com/goravel/gateway));
7. The HTTP interface forward to the corresponding GRPC interface(through [grpc-ecosystem/grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway));

## Key Points

Based on the request process, we can get some key points:

1. All proto files are defined in the proto folder, and the GRPC endpoints that expose to the outside need to 
   contain the `google.api.http` annotation, for example: `option (google.api.http) = {get: "/v1/user"}`;
2. We need to define a corresponding HTTP routing based on the annotation in step 1, and all HTTP routing is 
   defined in the route folder of gateway;
3. All microservices only need to define the GRPC endpoints, and the gateway will automatically generate the 
   HTTP routing to the corresponding GRPC endpoint;
4. If you want to get the current login user information in a microservice, you can add `string user_id = 1`, 
   `string user_id = 2` to your GRPC request, the fields will be filled by gateway automatically; 

## Required Tools

- [mockery](https://vektra.github.io/mockery/latest/installation/#github-release): Generate mock files for testing.
- [protoc](https://grpc.io/docs/protoc-installation/): Generate proto files.
- [protoc-gen-go](https://grpc.io/docs/languages/go/quickstart/#prerequisites): Generate go files from proto files.
  - `go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`
- [protoc-gen-go-grpc](https://grpc.io/docs/languages/go/quickstart/#prerequisites): Generate go grpc files from proto 
  files.
  - `go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest`
- [protoc-gen-grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway): Generate grpc gateway files from proto files.
  - `go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2`
