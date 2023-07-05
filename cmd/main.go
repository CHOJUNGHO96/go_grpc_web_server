package main

import (
	"net"
	server "sequence_game_server/cmd/server"
	"sequence_game_server/core/config"
	dbInstance "sequence_game_server/core/db"
	"sequence_game_server/core/util"
	grpcServer "sequence_game_server/pkg"

	_ "github.com/lib/pq"
)

func main() {
	// TCP 네트워크 연결 수신 시작
	util.Logger.Println("start server")
	lis, err := net.Listen("tcp", ":4186")
	if err != nil {
		util.Logger.Fatalf("failed to listen: %v", err)
	}

	// DB 연결 변수 및 설정 읽기
	var (
		db  = new(dbInstance.PostgresDB)
		dns = util.ReadPgJson(config.Config)
	)

	defer db.Close()

	// DB 연결 시작
	if _, err = db.NewPgConnection(dns); err != nil {
		util.Logger.Fatal(err)
	}

	// gRPC 서비스 초기화
	server.RegisterServer(db)

	// gRPC 서버 시작
	if err := grpcServer.GrpcServer.Serve(lis); err != nil {
		util.Logger.Fatalf("failed to serve: %v", err)
	}
}
