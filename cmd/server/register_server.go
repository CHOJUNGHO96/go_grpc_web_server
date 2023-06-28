package server

import (
	dbInstance "sequence_game_server/core/db"
	testService "sequence_game_server/pkg/test"
	test2Service "sequence_game_server/pkg/test2"
)

// RegisterServer gRPC 서버에 서비스를 등록하는 함수
func RegisterServer(db *dbInstance.PostgresDB) {
	// 각 서비스 초기화하고 gRPC 서버에 등록합니다.
	testService.InitService(db)
	test2Service.InitService(db)
}
