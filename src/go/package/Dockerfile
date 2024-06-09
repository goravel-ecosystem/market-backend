FROM alpine:latest

WORKDIR /www

COPY ./main /www/
COPY ./database/ /www/database/
COPY ./public/ /www/public/
COPY ./storage/ /www/storage/
COPY ./resources/ /www/resources/
COPY ./lang/ /www/lang/

ENV APP_NAME=package \
    APP_ENV=${APP_ENV}  \
    APP_KEY=${APP_KEY} \
    APP_DEBUG=${APP_DEBUG} \
    GRPC_HOST=0.0.0.0 \
    GRPC_PORT=${GRPC_PORT} \
    DB_CONNECTION=${DB_CONNECTION} \
    DB_HOST=${DB_HOST} \
    DB_PORT=${DB_PORT} \
    DB_DATABASE=${DB_DATABASE} \
    DB_USERNAME=${DB_USERNAME} \
    DB_PASSWORD=${DB_PASSWORD} \
    REDIS_HOST=${REDIS_HOST} \
    REDIS_PASSWORD=${REDIS_PASSWORD} \
    REDIS_PORT=${REDIS_PORT}

ENTRYPOINT ["/www/main"]
