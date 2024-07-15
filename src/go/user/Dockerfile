FROM alpine:latest

WORKDIR /www

COPY ./main /www/
COPY ./database/ /www/database/
COPY ./public/ /www/public/
COPY ./storage/ /www/storage/
COPY ./resources/ /www/resources/
COPY ./lang/ /www/lang/

ENV APP_NAME=user \
    GRPC_HOST=0.0.0.0

ENTRYPOINT ["/www/main"]
