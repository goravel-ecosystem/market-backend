docker run -dit \
  -p $INPUT_HTTP_PORT:$INPUT_HTTP_PORT \
  -e APP_ENV=$INPUT_APP_ENV \
  -e APP_KEY=$INPUT_APP_KEY \
  -e APP_DEBUG=true \
  -e APP_HOST=0.0.0.0 \
  -e APP_PORT=$INPUT_HTTP_PORT \
  -e GRPC_USER_HOST=$INPUT_GRPC_USER_HOST \
  -e GRPC_USER_PORT=$INPUT_GRPC_USER_PORT \
  -e GATEWAY_HOST=0.0.0.0 \
  -e GATEWAY_PORT=$INPUT_GATEWAY_PORT \
  --network goravel-market \
  --network-alias $INPUT_APP_NAME \
  --name goravel-market-$INPUT_APP_NAME \
  $INPUT_IMAGE