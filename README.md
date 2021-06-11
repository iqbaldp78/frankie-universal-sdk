# Boilerplate API

---
|Branch|Build Status|Coverage Report|
|:-------|:----------|:----|
|__Master__|[![build status](https://gitlab.com/datanest-engine/dce/gin-boilerplate/badges/master/build.svg)](https://gitlab.com/datanest-engine/dce/gin-boilerplate/commits/master)|[![coverage report](https://gitlab.com/datanest-engine/dce/gin-boilerplate/badges/master/coverage.svg)](https://gitlab.com/datanest-engine/dce/gin-boilerplate/commits/master)|
|__Develop__|[![build status](https://gitlab.com/datanest-engine/dce/gin-boilerplate/badges/develop/build.svg)](https://gitlab.com/datanest-engine/dce/gin-boilerplate/commits/develop)|[![coverage report](https://gitlab.com/datanest-engine/dce/gin-boilerplate/badges/develop/coverage.svg)](https://gitlab.com/datanest-engine/dce/gin-boilerplate/commits/develop)|

---

## Objective

This repo is starter pack to build REST API service ([gin base](https://github.com/gin-gonic/gin)) and communicate with postgreSQL. The goal is develop REST API service faster because no need to build from scratch.

## Requirement

- [Go 1.12 or higher](https://golang.org/dl/)
- [PostgreSQL 10 or higher](https://www.postgresql.org/download/)

## Installation

### Local environment

- Clone this repo

```bash
git clone https://gitlab.com/datanest-engine/dce/gin-boilerplate.git
```

- Put outside `$GOPATH`
- Run this service

```bash
cd $REPO_DIR
go run main.go
```

### Docker environment

- Login to registry.gitlab.com registry

```bash
docker login registry.gitlab.com
```

- Run docker

```bash
docker run -p 80:80 registry.gitlab.com/datanest-engine/dce/gin-boilerplate:latest
```

## Development

If you are __developer__ please see this readme first and information below may give help to you

### Project Structure

This is general project structure. If you add/change file the structure please change structure below

```ANSI
├── module                              //package for module file
|   ├── crud                            //package for crud operation
|   |   ├── sample_test.go              //unit test for sample.go file
|   |   └── sample.go                   //crud operation for sample
|   ├── sample                          //package for handling logic sample
|   |   ├── models                      //package for handling logic for model
|   |   |   ├── add_test.go             //unit test for add_sample.go
|   |   |   ├── add.go                  //logic for handling add sample operation
|   |   |   ├── change_test.go          //unit test for change_sample.go
|   |   |   ├── change.go               //logic for handling change sample operation
|   |   |   ├── get_test.go             //unit test for get_sample.go
|   |   |   ├── get.go                  //logic for handling get sample operation
|   |   |   ├── remove_test.go          //unit test for remove_sample.go
|   |   |   ├── remove.go               //logic for handling remove sample operation
|   |   |   ├── search_test.go          //unit test for search_sample.go
|   |   |   └── search.go               //logic for handling search sample operation
|   |   └── sample.go                   //logic for controller
|   ├── shared                          //package for shared object, function, or library
|   |   ├── lib                         //package for library collection
|   |   |   ├── jose                    //jose(Javascript Object Signing and Encryption) library
|   |   |   └── postgres                //postgresql library
|   |   ├── object                      //package for object collection
|   |   └── tools                       //package for function collection
|   ├── blueprints.go                   //logic for define route blueprint
|   ├── init.go                         //logic for initiate service
|   └── middleware.go                   //logic for define middleware to inject on route
├── vendor                              //vendor directory
├── .gitlab-ci.yml                      //script for run gitlab ci
├── apidoc.json                         //configuration file for apidocs generator
├── config.json                         //configuration file for this service
├── Dockerfile                          //dockerfile for build docker image
├── go.mod                              //gomod
├── go.sum                              //generated file from gomod
├── main.go                             //main function to run project
└── README.md                           //this file
```

### Environment Variables

This is changeable environment variables. If you add/change more variable on config file please change information below

|Variable|Description|Value|
|:-------|:----------|:----|
|DCE_SERVER_PORT|Server port listen|Default `80`|
|DCE_DB_HOST|Database host IP or URL|Default `localhost`|
|DCE_DB_PORT|Database port|Default `5432`|
|DCE_DB_USER|Database user|Default `postgres`|
|DCE_DB_PASS|Database password|Default blank|
|DCE_DB_NAME|Database name|Default `postgres`|
|DCE_SECRET_KEY|Secret key to encrypt jwt string|Default `8BF5B017E1C6560647DA276C1BF6391A4FF958911D1056B14C0D1165689985CF`|
|DCE_INGRIDIENT|Apikey to consume internal service on stack environment|Default `B3e6Tjas1k2DwCU4WuOfUseIEcgOgqbH`|
|DCE_TOKEN_ISSUER|Token issuer for matching issuer section on jwt|Default `api.boilerplate.com`|
|DCE_APP|Service debug mode. change value to `PRD` if you want to disable debug|Default blank|
