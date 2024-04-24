# Account Service #
[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)
[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/gomods/athens.svg)](https://bitbucket.org/rctiplus/tusd-server)

Account Service is a gRPC service for managing account creation, update, search and specific search

## System Requirement ##
* GO SDK +1.21
* ZapLogger
* Cobra CLI

### setup and installation ###

* Install GO SDK

Please refer to golang [installation](https://golang.org/doc/install) on how to install GO SDK match with your system (Mac, Windows, Linux)

* Configuration

Before application can be started, you must set configuration using `.env` file which the sample file can
be found in the root folder.

below is example configuration, change value inside curly braces to match with configuration in server.
```
GRPC_HOST={gRPC server host}
GRPC_PORT={gRPC server port}
GRPC_TLS={gRPC server TLS (true|false)}
GRPC_TIMEOUT={gRPC server timeout}
MONGO_HOST={MongoDB host}
MONGO_PORT={MongoDB port}
MONGO_DATABASE={MongoDB database name}
MONGO_USERNAME={MongoDB database username}
MONGO_PASSWORD={MongoDB database password}
```

note : _please fill configuration values with the correct one or match with your system._

* Dependencies

To download dependencies which are used by the app, you can use `go mod` tool, which is a standard tool for
dependency management.

manual download dependencies
```bash
$go mod tidy && go mod vendor
```

using `make` command
```bash
$make dep
```

* Build instructions

binary deployment using `go tool` command
```bash
$go build -o account_svc
$./account_svc grpc
```

build binary using `make` command
```bash
$make build
$./account_svc grpc
```

* Run instructions

You can run application directly without having to build the code, to directly run the app you can
use `go run` tool or using `make` command.

running directly using `go tool` command
```bash
$go run main.go serve
```

build binary using `make` command
```bash
$make run
```

* Deployment instructions

There are two type of deployment which are : binary and docker, you can choose what suit your need.
For binary deployment, you can use `rsync` command or `sftp` binary file to instance server.


### TODO ###

* Unit test
* Containerize (untested)



