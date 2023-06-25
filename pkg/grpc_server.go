package server

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var GrpcServer *grpc.Server

func init() {
	GrpcServer = grpc.NewServer()
	reflection.Register(GrpcServer)
}
