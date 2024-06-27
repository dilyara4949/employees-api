# Employee API

## Introduction

This API provides comprehensive CRUD (Create, Read, Update, Delete) functionality for managing employees and positions within an organization.

### Database

Command to run container for project's database:

```docker run --name postgre --network mynetwork -e POSTGRES_PASSWORD=12345 -p 5432:5432 -d postgres```

To run this command, you will need official postgres image.