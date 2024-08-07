docker run -dit \
  -p $INPUT_GATEWAY_HTTP_PORT:$INPUT_GATEWAY_HTTP_PORT \
  -e APP_ENV=$INPUT_APP_ENV \
  -e APP_KEY=$INPUT_GATEWAY_APP_KEY \
  -e APP_DEBUG=true \
  -e APP_HOST=0.0.0.0 \
  -e APP_PORT=$INPUT_GATEWAY_HTTP_PORT \
  -e GRPC_USER_HOST=$INPUT_USER_GRPC_HOST \
  -e GRPC_USER_PORT=$INPUT_USER_GRPC_PORT \
  -e GRPC_PACKAGE_HOST=$INPUT_PACKAGE_GRPC_HOST \
  -e GRPC_PACKAGE_PORT=$INPUT_PACKAGE_GRPC_PORT \
  -e GATEWAY_HOST=0.0.0.0 \
  -e GATEWAY_PORT=$INPUT_GATEWAY_GATEWAY_PORT \
  --network $INPUT_APP_ENV \
  --network-alias goravel-market-$INPUT_APP_NAME \
  --name goravel-market-$INPUT_APP_ENV-$INPUT_APP_NAME \
  $INPUT_IMAGE