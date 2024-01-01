# Gateway

The gateway is a microservice that provides a unified entry point for all clients. It is responsible for request
forwarding, authentication, monitoring, and other functions. It has the following features:

1. Define the HTTP routing for all microservices;
2. Check JWT token, get user information from UserService and put it into GRPC request: user_id, user_name;

## Run In Local

1. Initialize .env File

The following configuration is required:

```
APP_NAME=Goravel
APP_ENV=local
APP_KEY=
APP_DEBUG=true
APP_URL=http://localhost
APP_HOST=127.0.0.1
APP_PORT=3000

# The host and port of the user microservice.
GRPC_USER_HOST=127.0.0.1
GRPC_USER_PORT=3010

# The host and port of gateway, for detail: https://github.com/goravel/gateway
GATEWAY_HOST=127.0.0.1
GATEWAY_PORT=3001
```
