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

package:
	protoc -I=./ \
		--go_out=./ --go_opt=paths=source_relative \
		--go-grpc_out=./ --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=./ --grpc-gateway_opt=paths=source_relative \
		package/package.proto

all: base user package