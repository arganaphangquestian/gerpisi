package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	network, err := net.Listen("tcp", "0.0.0.0:8000")
	if err != nil {
		log.Fatal("Make sure GRPC Server is Running")
	}
	grpcServer := grpc.NewServer()

	if err := grpcServer.Serve(network); err != nil {
		log.Fatal("Failed to connect")
	}
}
