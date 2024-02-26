docker run -dit \
  -p $HTTP_PORT:$HTTP_PORT \
  -e APP_ENV=${{ inputs.env }} \
  -e APP_KEY=$APP_KEY \
  -e APP_DEBUG=true \
  -e APP_HOST=0.0.0.0 \
  -e APP_PORT=$HTTP_PORT \
  -e GRPC_USER_HOST=$GRPC_USER_HOST \
  -e GRPC_USER_PORT=$GRPC_USER_PORT \
  -e GATEWAY_HOST=0.0.0.0 \
  -e GATEWAY_PORT=$GATEWAY_PORT \
  --network goravel-market \
  --network-alias $NAME \
  --name goravel-market-$NAME \
  ${{ secrets.ALIYUN_ACR_REGISTRY }}:${{ inputs.tag }}