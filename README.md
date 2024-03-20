# Grocery Delivery Service

This is a grocery delivery service application that allows customers to add groceries, search for groceries, place orders, and view their order history.

## Prerequisites

- Go installed on your machine
- PostgreSQL database

## Getting Started

1. **Clone the repository:**
 
- git clone https://github.com/Avinashanakal/grocery-delivery-service.git

2. **Run App Locally:**
```bash
export ENV=dev
go run cmd/api/main.go
```

## Database Schema

The database schema is defined in the [`docs/table.sql`](docs/table.sql) file.

You can find a visual representation of the schema on [dbdiagram.io](https://dbdiagram.io/d/654d5ad07d8bbd6465e227f1).

## Relationships between Tables

- **A customer can place multiple orders.**
- **An order can contain multiple grocery items.**
- **A grocery item can be included in multiple orders.**

## Swagger Documentaion

* Swagger Documentaion can be found in [`docs/swagger.yaml`](docs/swagger.yaml)

## Postman Collection

* Postman Collection can be found in [`docs/postman/grocery-service.postman_collection.json`](docs/postman/grocery-service.postman_collection.json)

## 
# Build the docker image
```bash
    docker build .
```

# Tag the image
```bash
   docker tag grocery-delivery-service avinash/grocery-delivery-service:1.0.0
```

# Login to docker with your docker Id
```bash
    docker login
```

# Push the image to docker hub
```bash
    docker push avinash/grocery-delivery-service:1.0.0
```