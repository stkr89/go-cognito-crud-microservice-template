# go-crud-microservice-template

A minimalist template to get you started quickly with building a pre-implemented CRUD microservice in Go. 
Just add your model specific fields, and you are all set :rocket:.

## Overview

- Based on [go-kit](https://github.com/go-kit/kit) framework
- Supports all CRUD operations
- Provides `http` and `gRPC` server
- Supports [validate](https://github.com/go-playground/validator) and [conform](https://github.com/leebenson/conform) for user input
- Uses [gorm](https://github.com/go-gorm/gorm) to interact with postgres DB
- Supports Docker

## Getting Started

- Rename module in `go.mod`
- Rename Docker tag in `Makefile`
- Rename model in `models/models.go` and add fields
- Populate fields for request and response objects in `types/types.go`
- Populate fields for request and response objects in `pb/model.proto`
- Generate gRPC client and server code by running the following:
```shell
$ cd pb
$ make generate
```
- Add fields to crud methods in `service/model_service.go`
- Update routes in `transports/http.go`
- Add following environment variables:
  - `DB_USERNAME`
  - `DB_PASSWORD`
  - `DB_HOST`
  - `DB_PORT`
  - `DB_NAME`
  - `DB_CLUSTER`
  - `GRPC_PORT`
  - `HTTP_PORT`