package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatal("Listen Err: %v", err)
	}
	grpcServer := grpc.NewServer()

	if err := grpcServer.Server(lis); err != nil {
		log.Fatal("gRPC serve Err on 9000: %v", err)
	}
}
