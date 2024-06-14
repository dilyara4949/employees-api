DB_URL=postgres://postgres:12345@localhost:5432/postgres?sslmode=disable
TEST_DB_URL=postgres://postgres:12345@localhost:5432/testpostgres?sslmode=disable

.PHONY: migrate-up migrate-down create-migration proto

proto:
	protoc --proto_path=protobuf --go_out=proto --go_opt=paths=source_relative --go-grpc_out=proto --go-grpc_opt=paths=source_relative protobuf/*.proto

run:
	JWT_TOKEN_SECRET=my_secret_key REST_PORT=8080 GRPC_PORT=50052 ADDRESS=0.0.0.0 DB_HOST=localhost DB_PORT=5432 DB_USER=postgres DB_PASSWORD=12345 DB_NAME=postgres go run cmd/main.go

lint:
	 golangci-lint run --enable-all

testall:
	go test -v ./...

htmlcov:
	go tool cover -html=coverage.out

testcov:
	go test -v -covermode=count -coverprofile=coverage.out ./...

dockrunpq:
	docker run -it --rm --network mynetwork postgres psql -h postgre -U postgres

migrate-up:
	migrate -database $(DB_URL) -path internal/database/migrations up

migrate-down:
	migrate -database $(DB_URL) -path internal/database/migrations down

migrate-test-up:
	migrate -database $(TEST_DB_URL) -path internal/database/migrations up

migrate-test-down:
	migrate -database $(TEST_DB_URL) -path internal/database/migrations down

create-migration:
	@read -p "migration name: " name; \
	migrate create -ext sql -dir internal/database/migrations -seq $$name

create-cont:
	 docker run --name postgre --network mynetwork -e POSTGRES_PASSWORD=12345 -p 5432:5432 -d postgres

#add-u-dock:
#	sudo usermod -aG docker $USER
#	newgrp docker
