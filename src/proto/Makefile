base:
	protoc -I=./ \
		--go_out=../go/proto --go_opt=paths=source_relative \
		--go-grpc_out=../go/proto --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=../go/proto --grpc-gateway_opt=paths=source_relative,generate_unbound_methods=true \
		base/base.proto

user:
	protoc -I=./ \
		--go_out=../go/proto --go_opt=paths=source_relative \
		--go-grpc_out=../go/proto --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=../go/proto --grpc-gateway_opt=paths=source_relative,generate_unbound_methods=true \
		user/user.proto

package:
	protoc -I=./ \
		--go_out=../go/proto --go_opt=paths=source_relative \
		--go-grpc_out=../go/proto --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=../go/proto --grpc-gateway_opt=paths=source_relative,generate_unbound_methods=true \
		package/package.proto

all: base user package