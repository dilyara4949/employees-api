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

### redis

Create docker container for redis with following command:

```
docker run -d --name redis -p 6379:6379 redis redis-server --requirepass "12345"
```

```make migrate-up```

### Configs

Do not forget to set all needed configuration variables, for example: 

```
export JWT_TOKEN_SECRET=my_secret_key

export REST_PORT=8080
export GRPC_PORT=50051

export ADDRESS=0.0.0.0

export DATABASE_TYPE=postgres

export MONGO_HOST=localhost
export MONGO_PORT=27017
export MONGO_USER=mongo
export MONGO_PASSWORD=12345
export MONGO_NAME=mongo
export MONGO_TIMEOUT=30
export POSITIONS_COLLECTION=positions
export EMPLOYEES_COLLECTION=employees

export POSTGRES_HOST=localhost
export POSTGRES_PORT=5432
export POSTGRES_USER=postgres
export POSTGRES_PASSWORD=12345
export POSTGRES_NAME=postgres
export POSTGRES_TIMEOUT=30
export POSTGRES_MAX_CONNECTIONS=20


export REDIS_HOST=localhost
export REDIS_PORT=6379
export REDIS_PASSWORD=12345
export REDIS_TIMEOUT=10
export REDIS_TTL=5
export REDIS_DATABASE=0
export REDIS_POOL_SIZE=10

```

If you are running project on docker compose, change database hosts to their service names, like: 
```
REDIS_HOST=redis
POSTGRES_HOST=db-postgres
MONGO_HOST=db-mongo
```

After initializing all the necessary dependencies, you can run project:
```
go run cmd/main.go
```