FROM alpine:latest

WORKDIR /www

COPY ./main /www/
COPY ./database/ /www/database/
COPY ./public/ /www/public/
COPY ./storage/ /www/storage/
COPY ./resources/ /www/resources/

ENV APP_NAME=gateway

ENTRYPOINT ["/www/main"]
