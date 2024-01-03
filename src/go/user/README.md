# User

Responsible for user registration, login, and other functions.

## Run In Local

1. Initialize .env File

The following configuration is required:

```
APP_NAME=Goravel
APP_ENV=local
APP_KEY=
APP_DEBUG=true

GRPC_HOST=127.0.0.1
GRPC_PORT=3010

JWT_SECRET=

DB_CONNECTION=postgresql
DB_HOST=127.0.0.1
DB_PORT=5433
DB_DATABASE=goravel
DB_USERNAME=goravel
DB_PASSWORD=goravel
```

2. Initialize Key

```bash
go run . artisan key:generate
go run . artisan jwt:secret
```

3. Run PostgreSQL Locally

```
docker-compose up -d
```

## Generate Mock Files

```bash
mockery
```
