package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/duongptryu/start_grpc/greet/greetpb"
	"google.golang.org/grpc"
)

type server struct{}

func (server *server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Println("Greet function is called")
	firstName := req.GetGreeting().GetFirstName()
	result := "Hello " + firstName
	res := &greetpb.GreetResponse{
		Result: result,
	}
	return res, nil
}

func main() {
	fmt.Println("hello world - I'm a server")

	lis, err := net.Listen("tcp", "0.0.0.0:9180")
	//port 50051 for gRpc
	if err != nil {
		log.Fatal("Failed to listen", err)
	}
	//create gRPC server
	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatal("failed to serve", err)
	}
}
