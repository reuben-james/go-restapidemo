# go-restapidemo
An exemplar CRUD API implemented in Go

## Quickstart 

### Setting DB Environment Vars for local development

Set the following Environment variables before creating the stack: 
```
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgrespass
DB_NAME=apidb
```

## Stand up the DB

```
docker compose up &
```

## Run the app locally

```
go run cmd/main.go
```

On first run, the application will attempt to connect to the Database, create the `articles` Table, then initialise the Table with the mock articles defined in `pkg/mocks/article.go`

You can populate the Articles structure with additional initial objects by adding them in here. 


## Planned Improvements

* Tests
* Container-ise the API, add Nginx frontend
* A simple React UI with CRUD functionality
* Healthcheck endpoints
* An exemplar CLI
* External Authentication Integration
* Authorization Model/RBAC
* TLS/Certs