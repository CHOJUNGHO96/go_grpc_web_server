package test

import (
	pb "sequence_game_server/api/v1/test"
	dbInstance "sequence_game_server/core/db"
	grpcServer "sequence_game_server/pkg"
	testRepository "sequence_game_server/pkg/test/repository"
	testUsecase "sequence_game_server/pkg/test/usecase"
)

type InitFunc func(*dbInstance.PostgresDB)

func InitService(db *dbInstance.PostgresDB) {
	// DB와 연결된 repository 생성
	repo := testRepository.NewRepository(db)

	// 서비스 생성 및 서버에 등록
	testService := testUsecase.NewTestService(repo)
	pb.RegisterTestServiceServer(grpcServer.GrpcServer, testService)
}
