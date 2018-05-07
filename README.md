# Services & APIs Tech Challenge #1

A basic modern RESTful API web service providing CRUD operations and search functionality for a directory of hard-coded company employees stored in an in-memory database

## Usage

### Prerequisites

You need [Go](https://golang.org/dl/) to install this project


### Installing

To download and install this project:

```
$ go get github.com/chimauwah/services-api-tech-challenge
```

To build and run:

```
$ cd services-api-tech-challenge
$ go build
$ ./services-api-tech-challenge
```

## API

The API specification can be found at : http://localhost:8080/api/docs


## Structure
```
├── db               
│   ├── database.go  // db repository layer
│   ├── schema.sql   // ddl for company employee data
│   └── seed.sql     // dml for company employee seed data
├── handler          
│   └── handlers.go  // API core handler / service layer
├── helper
│   └── helpers.go   // 
├── middleware
│   └── redoc.go     // serves doc site for swagger spec
├── model
│   └── models.go    // models for application
│   └── swagg.go     // models for swagger specs
└── server.go        // initializes application and starts server
```

