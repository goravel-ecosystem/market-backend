base:
	protoc -I=./ \
		--go_out=./ --go_opt=paths=source_relative \
		--go-grpc_out=./ --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=./ --grpc-gateway_opt=paths=source_relative \
		base/base.proto

user:
	protoc -I=./ \
		--go_out=./ --go_opt=paths=source_relative \
		--go-grpc_out=./ --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=./ --grpc-gateway_opt=paths=source_relative \
		user/user.proto

business:
	protoc -I=./ \
		--go_out=./ --go_opt=paths=source_relative \
		--go-grpc_out=./ --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=./ --grpc-gateway_opt=paths=source_relative \
		business/business.proto

all: base user business