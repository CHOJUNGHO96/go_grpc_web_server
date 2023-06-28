package server

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// GrpcServer gRPC 서버의 인스턴스를 저장
var GrpcServer *grpc.Server

func init() {
	// 새로운 gRPC 서버 인스턴스를 생성하고 저장
	GrpcServer = grpc.NewServer()

	// 클라이언트가 런타임에 서버의 API 조회할 수 있도록 서버 인스턴스에 리플렉션 서비스를 등록 (ex : grpcurl 사용용도)
	reflection.Register(GrpcServer)
}
