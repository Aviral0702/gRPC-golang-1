package main

import (
	"context"
	"fmt"
	"golang-gRPC/golang-gRPC/calculatorpb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	calculatorpb.UnimplementedCalculatorServer
}

func (*server) Add(ctx context.Context, req *calculatorpb.AddRequest) (*calculatorpb.AddResponse, error) {
	a := req.GetA()
	b := req.GetB()
	result := a + b
	res := &calculatorpb.AddResponse{
		Result: result,
	}
	return res, nil
}

func main() {
	fmt.Println("Starting the gRPC server")
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen:  %v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
