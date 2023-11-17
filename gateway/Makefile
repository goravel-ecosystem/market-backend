proto-base:
	protoc -I=proto \
		--go_out=proto --go_opt=paths=source_relative \
		--go-grpc_out=proto --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=proto --grpc-gateway_opt=paths=source_relative \
		base/base.proto

proto-user:
	protoc -I=proto \
		--go_out=proto --go_opt=paths=source_relative \
		--go-grpc_out=proto --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=proto --grpc-gateway_opt=paths=source_relative \
		users/users.proto

proto-business:
	protoc -I=proto \
    		--go_out=proto --go_opt=paths=source_relative \
    		--go-grpc_out=proto --go-grpc_opt=paths=source_relative \
    		--grpc-gateway_out=proto --grpc-gateway_opt=paths=source_relative \
    		business/business.proto

proto: proto-base proto-user proto-business