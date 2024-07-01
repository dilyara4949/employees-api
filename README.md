# Employee API

## Introduction

This API provides comprehensive CRUD (Create, Read, Update, Delete) functionality for managing employees and positions within an organization.

### Database

Command to run container for project's database:

```docker run --name postgre --network mynetwork -e POSTGRES_PASSWORD=12345 -p 5432:5432 -d postgres```

or

```
docker run -d --name mongo -p 27017:27017 -e MONGO_INITDB_ROOT_USERNAME=mongo -e /
MONGO_INITDB_ROOT_PASSWORD=12345 mongo
```

To run this commands, you will need official postgres and mongo images.