package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/duongptryu/excersice_unary/calculate/calculatepb"
	"google.golang.org/grpc"
)

type server struct{}

func (sv *server) Sum(ctx context.Context, req *calculatepb.CalculateRequest) (*calculatepb.CalculateResponse, error) {
	firstNumber := req.GetCalculate().GetFirstNumber()
	secondNumber := req.GetCalculate().GetSecondNumber()
	result := firstNumber + secondNumber
	res := &calculatepb.CalculateResponse{
		Result: result,
	}
	return res, nil
}

func main() {
	fmt.Println("Start server")

	lis, err := net.Listen("tcp", "0.0.0.0:9180")
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	s := grpc.NewServer()
	calculatepb.RegisterCalculateServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatal("Failed to serve")
	}
}
