# Go Micro

Udemy: Working with microservices in Go

See the project/Makefile for targets to build the suite of docker containers and start/stop the front end.

The front end is a basic go application that serves a HTML page.

The database is populated manually at this point (i.e. use a Postgres client and import the SQL in the included users.sql)

## Dependencies

### Protocol Buffers

Go tools
- go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.27
- go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

Protocol Buffer Compiler
- https://grpc.io/docs/protoc-installation/
- apt install -y protobuf-compiler

