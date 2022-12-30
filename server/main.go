package main

import (
	"log"
	"net"

	pb "github.com/x-abgth/sample_grpc_1/proto"

	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to start the server - %v", err)
	}

	// Create new grpc server
	grpcServer := grpc.NewServer()
	pb.RegisterGreetServer(grpcServer, &helloServer{})
	log.Printf("Server started at %v", lis.Addr())
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to start - %v", err)
	}
}
