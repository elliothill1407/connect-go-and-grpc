package main

import (
	"context"
	"log"

	greetv1 "connect-go-and-grpc/grpc-go-demo/gen/greet/v1"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := greetv1.NewGreetServiceClient(conn)
	res, err := client.Greet(context.Background(), &greetv1.GreetRequest{Name: "Jane"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", res.Greeting)
}
