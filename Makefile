proto:
	protoc protobuf/employee.proto --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative

rest:
	JWT_TOKEN_SECRET=my_secret_key REST_PORT=8080 GRPC_PORT=50052 ADDRESS=0.0.0.0 go run cmd/main.go

grpc:
	JWT_TOKEN_SECRET=my_secret_key REST_PORT=8080 GRPC_PORT=50052 ADDRESS=0.0.0.0 go run cmd/grpc/main.go

test:
	go test ./...

both:
	JWT_TOKEN_SECRET=my_secret_key REST_PORT=8080 GRPC_PORT=50052 ADDRESS=0.0.0.0 go run cmd/rest_grpc/main.go