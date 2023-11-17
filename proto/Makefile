base:
	protoc -I=./ \
		--go_out=./ --go_opt=paths=source_relative \
		--go-grpc_out=./ --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=./ --grpc-gateway_opt=paths=source_relative \
		base/base.proto

users:
	protoc -I=./ \
		--go_out=./ --go_opt=paths=source_relative \
		--go-grpc_out=./ --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=./ --grpc-gateway_opt=paths=source_relative \
		users/users.proto

business:
	protoc -I=./ \
    		--go_out=./ --go_opt=paths=source_relative \
    		--go-grpc_out=./ --go-grpc_opt=paths=source_relative \
    		--grpc-gateway_out=./ --grpc-gateway_opt=paths=source_relative \
    		business/business.proto

all: base users business