# Code

## Go

We move each `go.mod` of microservices to the `/src/go` folder, there are some benefits:

1. We can use `go mod tidy` to manage the dependencies of all microservices;
2. We can upgrade goravel/framework uniformly, imagine that we have 10 microservices, they are using different 
   versions of goravel/framework and there are some break changes, we have to use different interface in different 
   microservices, it's a nightmare to maintain them;
3. We can run Github Action to check the code quality of all microservices, it's very convenient to find the 
   problems;
4. It's easy to use the interface in the `proto` folder;

## Develop Locally

Each microservice has its own DB and dependent on other microservices. It's very complex if we develop one microservice 
locally and have to run other all dependent microservices. So we use WireGuard VPN to build a bridge between local and 
staging environment, we can focus on one microservice development locally and connect other Staging microservices 
directly.

Please check the document to see how to use WireGuard VPN: [Link](README_VPN.md).

Once you connect the Staging environment successfully through WireGuard VPN, you can set the host of dependent 
microservices to the Staging environment IP address, and set the port of dependent based on the [config/deploy.yml](../config/deploy.yml). 

For example, if your microservice depends on the `user` microservice, you can set `GRPC_USER_HOST=10.0.0.1` and 
`GRPC_USER_PORT=4011` in your `.env` file. The `10.0.0.1` is the IP address of the VPN server, and the `4011` is the 
value of the key `staging.user.grpc.port` of the `deploy.yml`. 

### Make GRPC Request Locally

You can make a GRPC request locally by following the steps:

1. Install the [grpcurl tool](https://github.com/fullstorydev/grpcurl);
2. Add `reflection.Register(facades.Grpc().Server())` before `facades.Grpc().Run()` of the microservice `main.go`, eg: [Link](https://github.com/goravel-ecosystem/market-backend/blob/3394494d845cce810e6498781d049ae51c3b52b3/src/go/user/main.go#L16);

Some useful commands:

```shell
# List all services, 127.0.0.1:3010 is the IP address of the microservice
grpcurl -plaintext 127.0.0.1:3010 list

# List all methods of a service
grpcurl -plaintext 127.0.0.1:3010 list user.UserService

# Describe a method
grpcurl -plaintext 127.0.0.1:3010 describe user.UserService.EmailLogin

# Call a method
grpcurl -d '{"email": "market@goravel.dev"}' -plaintext 127.0.0.1:3010 user.UserService.EmailLogin
```
