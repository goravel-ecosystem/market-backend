FROM alpine:latest

WORKDIR /www

COPY ./main /www/
COPY ./database/ /www/database/
COPY ./public/ /www/public/
COPY ./storage/ /www/storage/
COPY ./resources/ /www/resources/

ENV APP_NAME=Gateway \
    APP_ENV=${APP_ENV}  \
    APP_KEY=${APP_KEY} \
    APP_DEBUG=${APP_DEBUG} \
    APP_HOST=0.0.0.0 \
    APP_PORT=${APP_PORT} \
    GRPC_USER_HOST=0.0.0.0 \
    GRPC_USER_PORT=${GRPC_USER_PORT} \
    GATEWAY_HOST=0.0.0.0 \
    GATEWAY_PORT=${GATEWAY_PORT}

ENTRYPOINT ["/www/main"]
