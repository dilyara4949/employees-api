.PHONY: proto

proto:
	protoc --proto_path=protobuf --go_out=proto --go_opt=paths=source_relative --go-grpc_out=proto --go-grpc_opt=paths=source_relative protobuf/*.proto

test:
	go test ./...

run:
	JWT_TOKEN_SECRET=my_secret_key REST_PORT=8080 GRPC_PORT=50052 ADDRESS=0.0.0.0 go run cmd/main.go

lint:
	 golangci-lint run --enable-all