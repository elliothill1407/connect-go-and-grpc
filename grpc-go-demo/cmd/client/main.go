package main

import (
	"context"
	"log"
	"os"

	// import the generated gRPC code for the Greet service
	greetv1 "connect-go-and-grpc/grpc-go-demo/gen/greet/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// use the name and address provided as CLI arguments, with defaults
	name := "Jane"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	addr := "localhost:8081"
	if len(os.Args) > 2 {
		addr = os.Args[2]
	}

	// connect to the server without any transport security
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// create a new Greet service client using the connection
	client := greetv1.NewGreetServiceClient(conn)

	req := &greetv1.GreetRequest{Name: name}

	// call the Greet method on the client with a background context and the request
	res, err := client.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Greeting: %s", res.Greeting)
}
