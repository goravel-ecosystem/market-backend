FROM alpine:latest

WORKDIR /www

COPY ./main /www/
COPY ./database/ /www/database/
COPY ./public/ /www/public/
COPY ./storage/ /www/storage/
COPY ./resources/ /www/resources/

ENV APP_NAME=gateway \
    APP_ENV=${APP_ENV}  \
    APP_KEY=${APP_KEY} \
    APP_DEBUG=${APP_DEBUG} \
    APP_HOST=${APP_HOST} \
    APP_PORT=${APP_PORT} \
    GRPC_USER_HOST=${GRPC_USER_HOST} \
    GRPC_USER_PORT=${GRPC_USER_PORT} \
    GRPC_PACKAGE_HOST=${GRPC_PACKAGE_HOST} \
    GRPC_PACKAGE_PORT=${GRPC_PACKAGE_PORT} \
    GATEWAY_HOST=${GATEWAY_HOST} \
    GATEWAY_PORT=${GATEWAY_PORT}

ENTRYPOINT ["/www/main"]
