package main

import (
	"context"
	"fmt"
	"log"
	"net/http" // provides HTTP client and server implementations

	"connectrpc.com/connect"     // import the Connect package
	"golang.org/x/net/http2"     // import HTTP/2 package
	"golang.org/x/net/http2/h2c" // import HTTP/2 cleartext (h2c) package

	// import the generated gRPC code for the Greet service (by protoc-gen-go)
	greetv1 "connect-go-and-grpc/connect-go-demo/gen/greet/v1"
	// import the generated Connect-Go code for the Greet service (by protoc-gen-connect-go)
	"connect-go-and-grpc/connect-go-demo/gen/greet/v1/greetv1connect"
)

// define a struct to implement the Greet service
type GreetServer struct{}

// implement the Greet method for the GreetServer
func (s *GreetServer) Greet(
	ctx context.Context,
	req *connect.Request[greetv1.GreetRequest], // the request is of type connect.Request with GreetRequest as the message
) (*connect.Response[greetv1.GreetResponse], error) { // the response is of type connect.Response with GreetResponse as the message
	log.Println("Request headers: ", req.Header())

	// create a new response with a greeting message
	res := connect.NewResponse(&greetv1.GreetResponse{
		Greeting: fmt.Sprintf("Hello, %s!", req.Msg.Name),
	})

	// set a custom header in the response
	res.Header().Set("Greet-Version", "v1")
	return res, nil
}

func main() {
	// create a new instance of GreetServer
	greeter := &GreetServer{}

	// create a new HTTP request multiplexer
	mux := http.NewServeMux()

	// create a new handler for the Greet service
	path, handler := greetv1connect.NewGreetServiceHandler(greeter)

	// register the handler with the multiplexer
	mux.Handle(path, handler)

	// listen on localhost port 8080
	http.ListenAndServe(
		"localhost:8080",
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)
}
