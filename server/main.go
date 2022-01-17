package main

import (
	"fmt"
	"log"
	"net"

	"github.com/arganaphangquestian/gerpisi/server/data"
	"github.com/arganaphangquestian/gerpisi/server/service"
	"google.golang.org/grpc"
)

func main() {
	network, err := net.Listen("tcp", "0.0.0.0:8000")
	if err != nil {
		log.Fatal("Make sure GRPC Server is Running")
	}
	s := service.Server{}
	grpcServer := grpc.NewServer()
	data.RegisterCalculateServer(grpcServer, &s)

	fmt.Println("GRPC Server running at http://localhost:8000")
	if err := grpcServer.Serve(network); err != nil {
		log.Fatal("Failed to connect")
	}
}
