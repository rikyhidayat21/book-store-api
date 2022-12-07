# book-store-api

## Application Structure

```
.
├── README.md
├── app
│   └── app.go
├── dto
│   └── bookDto
│       └── BookResponse.go
├── exception
│   └── errors.go
├── go.mod
├── go.sum
├── internal
│   ├── core
│   │   ├── domain
│   │   │   └── book.go
│   │   ├── ports
│   │   │   ├── repositories.go
│   │   │   └── services.go
│   │   └── services
│   │       └── booksvc
│   │           └── service.go
│   ├── handlers
│   │   └── bookhdl
│   │       └── http.go
│   └── repositories
│       └── booksrepo
│           └── book.go
├── logger
│   └── logger.go
├── main.go
└── shared
    └── writeResponse.go

16 directories, 15 files
```

## Definition of the Application structure
- app -> serve and wiring the application
- exception -> handling error
- internal
    - core -> All the core components (domain, services, ports) will be placed in the directory ./internal/core
    - domain -> All the domain models will be placed in the directory `./internal/core/domain`. It contains the `go struct` definition of each entity that is part of the domain problem, and can be used accross the application.
    - ports -> The ports will be placed in the directory `./internal/core/ports`. It contains the interfaces definition used to communicate with `actors`.
    - services -> The services are our entry points to the core, and each one of them implements the corresponding port. They will be placed in the directory `./internal/core/services`
    - Adapters -> implement the adapters so the application can interact with the actors
      - driver adapters -> All the driver adapters will be placed in directory `./internal/handlers`. This driver adapter must be capable of transform an http request into a service cal.
      - driven adapters -> All the driven adapters will be placed in directory `./internal/repositories`. This driven adapter must be satisfy the `ports`
- logger -> logging application
- shared -> helper
- main.go -> starting point the application

## Dependencies
- install router gorilla/mux -> go get -u github.com/gorilla/mux
- install logger using zap -> go get -u go.uber.org/zap
- install mysql driver -> go get -u github.com/go-sql-driver/mysql
- install sqlx -> go get -u github.com/jmoiron/sqlx

## How to develop
Always start from the core domain.

## Sources
- [src](https://medium.com/@matiasvarela/hexagonal-architecture-in-go-cfd4e436faa3)
- 