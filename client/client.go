package main

import (
	"context"
	"fmt"
	"golang-gRPC/golang-gRPC/calculatorpb"
	"log"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Connecting to the gRPC server")
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatal("Failed to connect: %v", err)
	}
	defer conn.Close()

	c := calculatorpb.NewCalculatorClient(conn)
	doAdd(c)
}

func doAdd(c calculatorpb.CalculatorClient) {
	req := &calculatorpb.AddRequest{
		A: 10,
		B: 20,
	}

	res, err := c.Add(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling Add RPC: %v", err)
	}

	fmt.Printf("Result of addition %d\n", res.GetResult())
}
