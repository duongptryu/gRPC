package main

import (
	"context"
	"fmt"
	"log"

	"github.com/duongptryu/start_grpc/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello, I'm a client")

	cc, err := grpc.Dial("localhost:9180", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Cannot connect to server", err)
	}

	c := greetpb.NewGreetServiceClient(cc)

	doUnary(c)
}

func doUnary(c greetpb.GreetServiceClient) {
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Duong",
			LastName:  "Trang",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatal("Error while calling RPC Greet - ", err)
	}
	log.Printf("Response from greet: %v", res.Result)
}
