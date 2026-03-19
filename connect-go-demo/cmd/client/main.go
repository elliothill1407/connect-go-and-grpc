package main

import (
	"context"
	"log"
	"net/http" // provides HTTP client and server implementations
	"os"

	"connectrpc.com/connect" // import the Connect package

	// import the generated protobuf types for the Greet service (by protoc-gen-go)
	greetv1 "connect-go-and-grpc/connect-go-demo/gen/greet/v1"

	// import the generated Connect-Go code for the Greet service (by protoc-gen-connect-go)
	"connect-go-and-grpc/connect-go-demo/gen/greet/v1/greetv1connect"
)

func main() {
	// create a new Greet service client
	client := greetv1connect.NewGreetServiceClient(
		http.DefaultClient, // use the default HTTP client
		"http://localhost:8080",
	)

	// use the name provided as a CLI argument, defaulting to "Jane"
	name := "Jane"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	req := connect.NewRequest(&greetv1.GreetRequest{Name: name})

	// call the Greet method on the client with a background context and the request
	res, err := client.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Greeting: %s", res.Msg.Greeting)
}
