# `gRPC-go` vs `Connect-go` Demo
This project demonstrates the implementation of a simple API using two different frameworks: gRPC and Connect-go. Both implementations provide a basic greeting service that accepts a name as input and returns a personalized greeting message.

## Table of Contents
- [`gRPC-go` vs `Connect-go` Demo](#grpc-go-vs-connect-go-demo)
  - [Table of Contents](#table-of-contents)
  - [Introduction](#introduction)
  - [Features](#features)
  - [Setup](#setup)
  - [Usage](#usage)
    - [Running the Servers](#running-the-servers)
    - [Making Requests](#making-requests)
    - [Additional Information](#additional-information)
  - [Contributing](#contributing)

## Introduction
The purpose of this project is to compare the usage and performance of two popular frameworks for building APIs: `gRPC-go` and `Connect-go`. Both frameworks offer features such as type-safe communication, code generation, and support for multiple programming languages.

## Features
- gRPC Implementation: Uses gRPC-go to define and implement the greeting service.
- Connect-go Implementation: Uses Connect-go to define and implement the greeting service.
- Client-Server Communication: Demonstrates how clients can communicate with servers using both HTTP/cURL and gRPC commands.

## Setup
Before running the examples, ensure you have the following prerequisites installed:
- `Go` (at least one of the last two major releases)
- `cURL` (for HTTP requests)
- Protocol Buffers compiler (`protoc`) and the Go plugins for Protocol Buffers (`protoc-gen-go`)

## Usage

### Running the Servers
gRPC Server:
```bash
go run ./grpc_server/main.go
```
Connect-go Server:
```bash
go run ./connect_server/main.go
```

### Making Requests
HTTP/cURL Requests
```bash
# gRPC Server (Replace 'Jane' with the desired name)
grpcurl -plaintext -d '{"name": "Jane"}' localhost:8080 greet.v1.GreetService/Greet

# Connect-go Server (Replace 'Jane' with the desired name)
curl -X POST -H "Content-Type: application/json" -d '{"name": "Jane"}' http://localhost:8080/greet.v1.GreetService/Greet
```
gRPC Requests
```bash
# gRPC Server (Replace 'Jane' with the desired name)
go run ./grpc_client/main.go Jane
```

### Additional Information
- Both implementations use Protocol Buffers for defining the service schema.
- The client programs demonstrate how to interact with the servers using different frameworks.
- For more details on each implementation, refer to the respective directories.

## Contributing
Contributions are welcome! If you have any suggestions, improvements, or feature requests, feel free to open an issue or create a pull request.