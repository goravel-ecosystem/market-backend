# Proto

Define the proto files for all microservices. We use `google/api/annotations.proto` to achieve HTTP to GRPC. If you 
want to expose a GRPC endpoint to the outside, you need to add the `google.api.http` annotation to the GRPC endpoint,
for example: `option (google.api.http) = {get: "/v1/user"}`, you can get more usages from [here](google/api/http.proto).

## Build Proto Files

You need to build proto files after you modify them, you can use the following commands to build them:

```bash
# Build all microservices proto files.
make all

# Build UserService
make user

# Build GatewayService
make gateway

# Build BusinessService
make business
```
