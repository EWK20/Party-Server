## Running the application

To start the application run:

```
make docker-up
```

This command spins up 2 containers, 1 running a MySQL 5.7 database and the other running the server

Port mapping was changed from 3000:3000 to 8080:4000 because i had some traffic going to port 3000 already

## Application Specs

- Gin web framework - https://github.com/gin-gonic/gin
- GORM(Object Relational Mapping library for Golang); provides CRUD operations and can also be used for the initial migration and creation of the database schema - https://gorm.io/
- Auto Migration to create databases based on model struct

## Documentation

To see the documentation for this application

- Install godocs onto your machine using the following command:

```
go install golang.org/x/tools/cmd/godoc@latest
```

- Ensure that your go bin folder is in your PATH
- Serve the Go documentation on your localhost and specify a free port, run:

```
godoc -http :PORT
```

In the future I would add swagger API documentation as well to properly document each endpoint.
