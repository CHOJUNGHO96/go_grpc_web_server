package test

import (
	pb "sequence_game_server/api/v1/test2"
	dbInstance "sequence_game_server/core/db"
	grpcServer "sequence_game_server/pkg"
	test2Repository "sequence_game_server/pkg/test2/repository"
	test2Usecase "sequence_game_server/pkg/test2/usecase"
)

func InitService(db *dbInstance.PostgresDB) {
	// DB와 연결된 repository 생성
	repo := test2Repository.NewRepository(db)

	// 서비스 생성 및 서버에 등록
	testService := test2Usecase.NewTestService(repo)
	pb.RegisterTest2ServiceServer(grpcServer.GrpcServer, testService)
}
