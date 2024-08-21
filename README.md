gRPC Calculator in Golang
This repository contains a simple calculator service implemented using gRPC in Golang. The service currently supports basic addition operations.

Table of Contents
Introduction
Prerequisites
Installation
Usage
Proto File
Testing the Service
Contributing
License
Introduction
This project demonstrates the implementation of a gRPC server in Golang, providing a calculator service that can perform addition of two numbers. gRPC is a high-performance, open-source framework that can efficiently connect services in and across data centers with pluggable support for load balancing, tracing, health checking, and authentication.

Prerequisites
Go 1.19 or later installed on your machine.
protoc (Protocol Buffers compiler) installed and added to your PATH.
gRPC and Protocol Buffers Go plugins installed:
bash
Copy code
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
Installation
Clone the repository:

bash
Copy code
git clone https://github.com/yourusername/grpc-calculator.git
cd grpc-calculator
Compile the Protocol Buffers definition:

bash
Copy code
protoc --go_out=. --go-grpc_out=. proto/calculator.proto
Build the server:

bash
Copy code
go build -o server cmd/server/main.go
Build the client:

bash
Copy code
go build -o client cmd/client/main.go
Usage
Running the Server
Start the gRPC server:

bash
Copy code
./server
The server will start and listen for incoming gRPC requests on localhost:50051.

Running the Client
In another terminal, run the client to perform an addition operation:

bash
Copy code
./client
You will be prompted to input two numbers, and the client will send the request to the server and print the result.
