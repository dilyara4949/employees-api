# Employee API

## Introduction

This API provides comprehensive CRUD (Create, Read, Update, Delete) functionality for managing employees and positions within an organization.

### Database

Command to run container for project's database:

```
docker run --name postgres --network network_name /
 -e POSTGRES_PASSWORD=12345 -p 5432:5432 -d postgres
```

or

```
docker run -d --name mongo -p 27017:27017 -e MONGO_INITDB_ROOT_USERNAME=mongo -e /
MONGO_INITDB_ROOT_PASSWORD=12345 mongo
```

To run this commands, you will need official postgres and mongo images.


#### migration
To work on with PostgreSQL database, make sure to migrate up, change DB_URL connection variable in makefile as your postgres configuration and then run:

```make migrate-up```

### Configs

Do not forget to set all needed configuration variables, for example: 

```
export JWT_TOKEN_SECRET=my_secret_key
export REST_PORT=8080
export GRPC_PORT=50052
export ADDRESS=0.0.0.0
export DB_HOST=localhost
export DB_PORT=5432
export DB_USER=postgres
export DB_PASSWORD=12345
export DB_NAME=postgres
export DB_TIMEOUT=5
export DB_MAX_CONNECTIONS=5
export REDIS_HOST=localhost
export POSTGRES_DB=postgres
export POSTGRES_PASSWORD=12345
export POSTGRES_USER=postgres
```

After initializing all the necessary dependencies, you can run project:
```
go run cmd/main.go
```