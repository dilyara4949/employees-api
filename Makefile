DB_URL=postgres://postgres:12345@localhost:5432/postgres?sslmode=disable
TEST_DB_URL=postgres://postgres:12345@localhost:5432/testpostgres?sslmode=disable
GOLANGCILINT ?= docker run --rm -v $(shell pwd):/app -w /app golangci/golangci-lint:v1.57.2 golangci-lint


.PHONY: migrate-up migrate-down create-migration proto lint

proto:
	protoc --proto_path=protobuf --go_out=proto --go_opt=paths=source_relative --go-grpc_out=proto --go-grpc_opt=paths=source_relative protobuf/*.proto

run:
	go run cmd/main.go

lint:
	  $(GOLANGCILINT) run -v --enable-all

testall:
	go test -v ./...

htmlcover:
	go tool cover -html=coverage.out

testcover:
	go test -v -covermode=count -coverprofile=coverage.out ./...

migrate-up:
	migrate -database $(DB_URL) -path internal/database/postgres/migrations up

migrate-down:
	migrate -database $(DB_URL) -path internal/database/postgres/migrations down

migrate-test-up:
	migrate -database $(TEST_DB_URL) -path internal/database/postgres/migrations up

migrate-docker-down:
	docker-compose run app migrate -path ./internal/database/postgres/migrations -database "postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@db-postgres:5432/${POSTGRES_NAME}?sslmode=disable" down

migrate-docker-up:
	docker-compose run app migrate -path ./internal/database/postgres/migrations -database "postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@db-postgres:5432/${POSTGRES_NAME}?sslmode=disable" up

migrate-test-down:
	migrate -database $(TEST_DB_URL) -path internal/database/postgres/migrations down

create-migration:
	@read -p "migration name: " name; \
	migrate create -ext sql -dir internal/database/postgres/migrations -seq $$name
