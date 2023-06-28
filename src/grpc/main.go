package main

import (
	"log"
	"net"

	summary_proto "task1/entity/summaryservice"
	"task1/src/grpc/server"

	g "google.golang.org/grpc"
	ref "google.golang.org/grpc/reflection"
)

func main() {
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}

	transportLayer := server.Server{}

	grpcServer := g.NewServer()
	ref.Register(grpcServer) // expose for easier query with tools

	summary_proto.RegisterSummaryServiceServer(grpcServer, transportLayer)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 9000 %v", err)
	}

}
