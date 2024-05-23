package main

import (
	"context"
	"log"

	// import the generated gRPC code for the Greet service
	greetv1 "connect-go-and-grpc/grpc-go-demo/gen/greet/v1"

	"google.golang.org/grpc"
)

func main() {
	// dial the gRPC server at localhost on port 8080 without any transport security
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// create a new Greet service client using the connection
	client := greetv1.NewGreetServiceClient(conn)

	// create a new GreetRequest with the name "Jane"
	req := &greetv1.GreetRequest{Name: "Jane"}

	// call the Greet method on the client with a background context and the request
	res, err := client.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Greeting: %s", res.Greeting)
}
