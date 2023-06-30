package main

import (
	"log"
	"net"
	"task1/src/calculation"
	"task1/src/repo"

	summary_proto "task1/entity/summaryservice"
	"task1/src/grpc/server"

	g "google.golang.org/grpc"
	ref "google.golang.org/grpc/reflection"
)

func main() {
	redisRepo := repo.NewRedisRepo(&repo.RedisOptions{
		Address: "127.0.0.1:6379",
	})

	ohlc := calculation.NewOHLCRecords(&calculation.OHLC{
		Store: redisRepo,
	})

	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}

	transportLayer := server.NewGRPCServiceHandler(&server.GRPCServiceHandler{
		OHLCUsecase: *ohlc,
	})

	grpcServer := g.NewServer()
	ref.Register(grpcServer) // expose for easier query with tools

	summary_proto.RegisterSummaryServiceServer(grpcServer, transportLayer)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 9000 %v", err)
	}

}
