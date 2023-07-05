package util

import (
	"log"
	"os"
	"path/filepath"
	path "sequence_game_server/core/config"
)

var (
	Logger *log.Logger
)

// InitLogger 전역으로 사용하는 로거를 초기화하는 함수
func InitLogger() {
	LogFileName := "sequence_game.log"

	// 로그 디렉토리 생성
	if _, err := os.Stat(path.LogDir); os.IsNotExist(err) {
		err := os.MkdirAll(path.LogDir, 0755)
		if err != nil {
			log.Fatalf("Failed to create log directory: %v", err)
		}
	}

	file, err := os.OpenFile(filepath.Join(path.LogDir, LogFileName), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}

	Logger = log.New(file, "", log.LstdFlags)
}

func init() {
	InitLogger()
}
