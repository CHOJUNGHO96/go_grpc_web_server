package server

import (
	"context"
	"sequence_game_server/core/util"

	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/reflection"
)

// GrpcServer gRPC 서버의 인스턴스를 저장
var GrpcServer *grpc.Server

// unaryInterceptor gRPC 서버를 인터셉터후 클라이언트 정보 로그기록
func unaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	p, ok := peer.FromContext(ctx)
	if ok {
		util.Logger.Printf("Request - Method:%s Addr:%s\n", info.FullMethod, p.Addr)
	}
	return handler(ctx, req)
}

func init() {
	// 새로운 gRPC 서버 인스턴스를 생성하고 저장
	GrpcServer = grpc.NewServer(grpc.UnaryInterceptor(unaryInterceptor))

	// 클라이언트가 런타임에 서버의 API 조회할 수 있도록 서버 인스턴스에 리플렉션 서비스를 등록 (ex : grpcurl 사용용도)
	reflection.Register(GrpcServer)
}
