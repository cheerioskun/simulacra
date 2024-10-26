package main

import (
	pb "llm-simulation/proto"
	"llm-simulation/world"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create and start world server
	worldServer := world.NewWorldServer()
	if err := worldServer.Start(); err != nil {
		log.Fatalf("failed to start world server: %v", err)
	}

	// Create gRPC server
	grpcServer := grpc.NewServer()
	pb.RegisterWorldServerServer(grpcServer, world.NewGRPCServer(worldServer))

	log.Printf("server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
