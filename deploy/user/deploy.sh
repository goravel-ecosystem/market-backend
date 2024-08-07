docker run -dit \
  -p $INPUT_USER_GRPC_PORT:$INPUT_USER_GRPC_PORT \
  -e APP_ENV=$INPUT_APP_ENV \
  -e APP_KEY=$INPUT_USER_APP_KEY \
  -e APP_DEBUG=true \
  -e GRPC_HOST=0.0.0.0 \
  -e GRPC_PORT=$INPUT_USER_GRPC_PORT \
  -e JWT_SECRET=$INPUT_USER_JWT_SECRET \
  -e DB_HOST=$INPUT_DB_HOST \
  -e DB_PORT=$INPUT_DB_PORT \
  -e DB_DATABASE=$INPUT_USER_DB_DATABASE \
  -e DB_USERNAME=$INPUT_DB_USERNAME \
  -e DB_PASSWORD=$INPUT_DB_PASSWORD \
  -e REDIS_HOST=$INPUT_REDIS_HOST \
  -e REDIS_PORT=$INPUT_REDIS_PORT \
  -e MAIL_HOST=$INPUT_MAIL_HOST \
  -e MAIL_PORT=$INPUT_MAIL_PORT \
  -e MAIL_USERNAME=$INPUT_MAIL_USERNAME \
  -e MAIL_PASSWORD=$INPUT_MAIL_PASSWORD \
  --network $INPUT_APP_ENV \
  --network-alias goravel-market-$INPUT_APP_NAME \
  --name goravel-market-$INPUT_APP_ENV-$INPUT_APP_NAME \
  $INPUT_IMAGE