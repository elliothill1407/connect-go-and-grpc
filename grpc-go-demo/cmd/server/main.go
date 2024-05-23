package main

import (
	"context"
	"fmt"
	"log"
	"net"

	// import the generated gRPC code for the Greet service
	greetv1 "connect-go-and-grpc/grpc-go-demo/gen/greet/v1"

	"google.golang.org/grpc"
)

// define a struct to implement the Greet service
type GreetServer struct {
	greetv1.UnimplementedGreetServiceServer
}

// implement the Greet method for the GreetServer
func (s *GreetServer) Greet(ctx context.Context, req *greetv1.GreetRequest) (*greetv1.GreetResponse, error) {
	log.Println("Received request for:", req.Name)
	res := &greetv1.GreetResponse{Greeting: fmt.Sprintf("Hello, %s!", req.Name)}
	return res, nil
}

func main() {
	// listen on TCP port 8080
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// create a new (unregistered) gRPC server
	s := grpc.NewServer()

	// register the GreetServer with the gRPC server
	greetv1.RegisterGreetServiceServer(s, &GreetServer{})
	log.Println("Server is running on port 8080...")

	// start serving requests on listener
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
