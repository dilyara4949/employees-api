run:
	JWT_TOKEN_SECRET=my_secret_key  PORT=8080 ADDRESS=0.0.0.0 go run cmd/main.go
exp:
	export JWTTokenSecret=my_secret_key
test:
	go test -v
testall:
	go test -v ./...
htmlcov:
	go tool cover -html=coverage.out
testcov:
	go test -v -covermode=count -coverprofile=coverage.out ./...
