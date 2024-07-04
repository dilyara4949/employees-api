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
JWT_TOKEN_SECRET=my_secret_key

REST_PORT=8080
GRPC_PORT=50051

ADDRESS=0.0.0.0

DATABASE_TYPE=postgres

MONGO_HOST=localhost
MONGO_PORT=27017
MONGO_USER=mongo
MONGO_PASSWORD=12345
MONGO_NAME=mongo
MONGO_TIMEOUT=30
POSITIONS_COLLECTION=positions
EMPLOYEES_COLLECTION=employees

POSTGRES_HOST=localhost
POSTGRES_PORT=5432
POSTGRES_USER=postgres
POSTGRES_PASSWORD=12345
POSTGRES_DB=postgres
POSTGRES_TIMEOUT=30
POSTGRES_MAX_CONNECTIONS=20


REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_PASSWORD=12345
REDIS_TIMEOUT=10
REDIS_TTL=5
REDIS_DATABASE=0
REDIS_POOL_SIZE=10

```



If you are running project on docker compose, all the env vars should be in .env file, also do not forget to change database hosts to their service names, like: 
```
REDIS_HOST=redis
POSTGRES_HOST=db-postgres
MONGO_HOST=db-mongo
```

After initializing all the necessary dependencies, you can run project:
```
go run cmd/main.go
```



### Testing

To test unit tests, use:
```
make testall 
```

to test integration tests:
``` 
make testintegration
```

Don't forget to change database names for integration tests
Before running integration tests, make sure to export all the necessary variables