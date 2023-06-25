package main

import (
	"log"
	"net"
	server "sequence_game_server/pkg"
	_ "sequence_game_server/pkg/test/usecase"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	if err := server.GrpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
