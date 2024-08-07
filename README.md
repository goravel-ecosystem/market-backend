# Market Backend

Providers a backend for Goravel Market. In facilitate deployment and testing, we put all microservices in this project.

## Code Documentation

[Link](src/README.md)

## API Documentation

[Link](https://htmlpreview.github.io/?https://github.com/goravel-ecosystem/market-backend/blob/master/src/doc/index.html#string)

## Deploy

We are using [the Github action](.github/workflows/build.yml) to deploy the Staging environment. The 
action will build a docker image and deploy it to the Staging server automatically when you [create a new tag](https://github.com/goravel-ecosystem/market-backend/releases/new) 
in the repository. 

There is a rule when you create a new tag, the name should consist of the service name and version, for example: 
`gateway-0.0.1`, the `gateway` is the folder name of `src/go/gateway`.

Once you create a new tag, please check the deployment process [here](https://github.com/goravel-ecosystem/market-backend/actions), 
to ensure the deployment is successful.

You can also deploy the staging environment manually, open [this page](https://github.com/goravel-ecosystem/market-backend/actions/workflows/deploy.yml) 
and click the `Run workflow` button, then select the branch, environment and tag.

### Add A New Service 

1. Create a new project in the `src/go` folder, the folder name should be the service name;
2. Configure deploy parameters in the [deploy/config.yml](deploy/config.yml) file;
3. Create the docker running command in the `deploy` folder: `deploy/{SERVICE_NAME}/deploy.sh`, add the environment 
   variables according to the `.env` file;
4. Optimize the `Dockerfile` in the `src/go/{SERVICE_NAME}` folder;
5. Add the environment variables to Github secrets, such as: `PACKAGE_APP_KEY`;
6. Add the environment variables to the `.github/workflows/deploy.yml` file;
7. If the new service providers external GRPC endpoints, add the GRPC host and port to `deploy/gateway/deploy.sh`;
8. [create a new tag](https://github.com/goravel-ecosystem/market-backend/releases/new) in Github to test the 
   service is deployed successfully;
9. [A reference PR](https://github.com/goravel-ecosystem/market-backend/pull/55);

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
   `string user_name = 2` to your GRPC request, the fields will be filled by gateway automatically; 

## Required Tools

- [mockery@v2.42.1](https://vektra.github.io/mockery/latest/installation/#github-release): Generate mock files for testing.
  - check the version by `mockery --version`
- [protoc@libprotoc 25.1](https://grpc.io/docs/protoc-installation/): Generate proto files.
  - check the version by `protoc --version`
- [protoc-gen-go@v1.33.0](https://grpc.io/docs/languages/go/quickstart/#prerequisites): Generate go files from proto files.
  - `go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.33.0`
  - check the version by `protoc-gen-go --version`
- [protoc-gen-go-grpc@1.3.0](https://grpc.io/docs/languages/go/quickstart/#prerequisites): Generate go grpc files from proto 
  files.
  - `go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0`
  - check the version by `protoc-gen-go-grpc --version`
- [protoc-gen-grpc-gateway@v2.19.1](https://github.com/grpc-ecosystem/grpc-gateway): Generate grpc gateway files from proto files.
  - `go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.19.1`
  - check the version by `protoc-gen-grpc-gateway --version`, but it will print: `Version dev, commit unknown, built at unknown`
- [protoc-gen-doc@v1.5.1](https://github.com/pseudomuto/protoc-gen-doc): Generate documentation from proto files.
  - `go install github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc@v1.5.1`
  - check the version by `protoc-gen-doc --version`
