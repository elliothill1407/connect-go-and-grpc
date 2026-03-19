# `gRPC-go` vs `Connect-go` Demo
This project demonstrates the implementation of a simple API using two different frameworks: gRPC and Connect-go. Both implementations provide a basic greeting service that accepts a name as input and returns a personalized greeting message.

## Table of Contents
- [`gRPC-go` vs `Connect-go` Demo](#grpc-go-vs-connect-go-demo)
  - [Table of Contents](#table-of-contents)
  - [Introduction](#introduction)
  - [Comparison](#comparison)
  - [Features](#features)
  - [Setup](#setup)
  - [Usage](#usage)
    - [Running the Servers](#running-the-servers)
    - [Making Requests](#making-requests)
    - [Protocol Interoperability](#protocol-interoperability)
    - [Additional Information](#additional-information)

## Introduction
The purpose of this project is to compare the usage and performance of two popular frameworks for building APIs: `gRPC-go` and `Connect-go`. Both frameworks offer features such as type-safe communication, code generation, and support for multiple programming languages.

## Comparison

|  | gRPC-go | Connect-go |
|---|---|---|
| Transport | HTTP/2 only | HTTP/1.1, HTTP/2, HTTP/3 |
| Wire format | Protobuf binary | Protobuf binary or JSON |
| Browser compatible | No (requires gRPC-web proxy) | Yes (JSON over HTTP/1.1) |
| Testable with `curl` | No | Yes |
| Testable with `grpcurl` | Yes | Yes (speaks gRPC protocol too) |
| Standard `net/http` middleware | No | Yes |

Both frameworks add negligible overhead over raw HTTP/2 + Protobuf. The meaningful performance variable is wire format: Connect-go's JSON mode carries typical JSON serialization cost, but when both frameworks use Protobuf binary over HTTP/2 they are functionally equivalent in throughput and latency.

## Features
- gRPC Implementation: Uses gRPC-go to define and implement the greeting service.
- Connect-go Implementation: Uses Connect-go to define and implement the greeting service.
- Client-Server Communication: Demonstrates how clients can communicate with servers using both HTTP/cURL and gRPC commands.

## Setup
Before running the examples, ensure you have the following prerequisites installed:
- `Go` (at least one of the last two major releases)
- `cURL` (for HTTP requests)
- `grpcurl` (for gRPC requests — `brew install grpcurl` or see [fullstorydev/grpcurl](https://github.com/fullstorydev/grpcurl))
- Protocol Buffers compiler (`protoc`) and the Go plugins for Protocol Buffers (`protoc-gen-go`)

## Usage

### Running the Servers
gRPC Server:
```bash
go run ./grpc-go-demo/cmd/server/main.go
```
Connect-go Server:
```bash
go run ./connect-go-demo/cmd/server/main.go
```

### Making Requests
HTTP/cURL Requests
```bash
# gRPC Server (Replace 'Jane' with the desired name)
grpcurl -plaintext -d '{"name": "Jane"}' localhost:8081 greet.v1.GreetService/Greet

# Connect-go Server (Replace 'Jane' with the desired name)
curl -X POST -H "Content-Type: application/json" -d '{"name": "Jane"}' http://localhost:8080/greet.v1.GreetService/Greet
```
Go Client Requests
```bash
# Connect-go Server (Replace 'Jane' with the desired name)
go run ./connect-go-demo/cmd/client/main.go Jane

# gRPC Server (Replace 'Jane' with the desired name)
go run ./grpc-go-demo/cmd/client/main.go Jane
```

### Protocol Interoperability

Connect-go speaks the gRPC wire protocol natively, so the gRPC-go client works against the Connect-go server without any changes. Start the Connect-go server, then point the gRPC client at it:

```bash
go run ./grpc-go-demo/cmd/client/main.go Jane localhost:8080
# 2026/03/19 19:35:12 Greeting: Hello, Jane!
```

This is the key practical implication of Connect-go: a single server handles gRPC clients, browser clients over JSON, and Connect clients simultaneously on the same port — no proxy required.

### Regenerating Protobuf Code

The two demos intentionally use different code generation toolchains.

Connect-go uses [Buf](https://buf.build):
```bash
cd connect-go-demo && buf generate
```

gRPC-go uses `protoc` directly:
```bash
cd grpc-go-demo
protoc --go_out=gen --go_opt=paths=source_relative \
       --go-grpc_out=gen --go-grpc_opt=paths=source_relative \
       greet/v1/greet.proto
```

### Additional Information
- Both implementations use the same Protocol Buffers schema to define the service.
- The client programs demonstrate how to interact with the servers using different frameworks.
- For more details on each implementation, refer to the respective directories.
