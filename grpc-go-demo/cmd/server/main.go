package main

import (
	"context"
	"fmt"
	"log"
	"net"

	greetv1 "connect-go-and-grpc/grpc-go-demo/gen/greet/v1"

	"google.golang.org/grpc"
)

type GreetServer struct {
	greetv1.UnimplementedGreetServiceServer
}

func (s *GreetServer) Greet(ctx context.Context, req *greetv1.GreetRequest) (*greetv1.GreetResponse, error) {
	log.Println("Received request for:", req.Name)
	return &greetv1.GreetResponse{Greeting: fmt.Sprintf("Hello, %s!", req.Name)}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	greetv1.RegisterGreetServiceServer(s, &GreetServer{})
	log.Println("Server is running on port 8080...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
