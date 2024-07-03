DB_URL=postgres://postgres:12345@localhost:5432/postgres?sslmode=disable
TEST_DB_URL=postgres://postgres:12345@localhost:5432/testpostgres?sslmode=disable

.PHONY: migrate-up migrate-down create-migration proto

proto:
	protoc --proto_path=protobuf --go_out=proto --go_opt=paths=source_relative --go-grpc_out=proto --go-grpc_opt=paths=source_relative protobuf/*.proto

run:
	go run cmd/main.go

lint:
	 golangci-lint run --enable-all

testall:
	go test -v ./...

htmlcov:
	go tool cover -html=coverage.out

testcov:
	go test -v -covermode=count -coverprofile=coverage.out ./...

migrate-up:
	migrate -database $(DB_URL) -path internal/database/postgres/migrations up

migrate-down:
	migrate -database $(DB_URL) -path internal/database/postgres/migrations down

migrate-test-up:
	migrate -database $(TEST_DB_URL) -path internal/database/postgres/migrations up

migrate-docker-down:
	docker-compose run app migrate -path ./internal/database/postgres/migrations -database "postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@db:5432/${POSTGRES_DB}?sslmode=disable" down

migrate-docker-up:
	docker-compose run app migrate -path ./internal/database/postgres/migrations -database "postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@db:5432/${POSTGRES_DB}?sslmode=disable" up

migrate-test-down:
	migrate -database $(TEST_DB_URL) -path internal/database/postgres/migrations down

create-migration:
	@read -p "migration name: " name; \
	migrate create -ext sql -dir internal/database/postgres/migrations -seq $$name
