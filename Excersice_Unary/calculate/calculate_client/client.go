package main

import (
	"context"
	"fmt"
	"log"

	"github.com/duongptryu/excersice_unary/calculate/calculatepb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("start client")

	cc, err := grpc.Dial("localhost:9180", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	c := calculatepb.NewCalculateServiceClient(cc)

	doSum(c)
}

func doSum(c calculatepb.CalculateServiceClient) {
	req := &calculatepb.CalculateRequest{
		Calculate: &calculatepb.Calculate{
			FirstNumber:  3,
			SecondNumber: 10,
		},
	}

	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Sum RPC: %v", err)
	}
	log.Printf("Result: %v", res.Result)
}
