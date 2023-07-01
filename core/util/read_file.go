package util

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func ReadPgJson(path string) string {
	defer Recovery()

	type DatabaseInfo struct {
		DatabaseID       string `json:"database_id"`
		DatabaseIP       string `json:"database_ip"`
		DatabaseName     string `json:"database_name"`
		DatabasePassword string `json:"database_password"`
		DatabasePort     int    `json:"database_port"`
	}

	type PostgresInformation struct {
		Info DatabaseInfo `json:"postgres_information"`
	}

	// 환경변수에서 AES 키를 읽어옴.
	key := []byte(os.Getenv("AES_KEY"))
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		log.Fatal("Invalid AES key length. It should be 16, 24 or 32 bytes.")
	}

	encryptedData, err := os.ReadFile(filepath.Join(path, "PostgresqlConfig.json"))
	if err != nil {
		log.Fatalf("Unable to read config file: %v", err)
	}

	// 파일에서 읽은 데이터를 복호화.
	encryptedDataBytes, err := hexDecode(string(encryptedData))
	if err != nil {
		log.Fatalf("Failed to decode hex string: %v", err)
	}

	decryptedData := decrypt(encryptedDataBytes, key)

	var config PostgresInformation
	err = json.Unmarshal(decryptedData, &config)
	if err != nil {
		log.Fatalf("Unable to parse config file: %v", err)
	}

	pgInfo := fmt.Sprintf("dbname=%s user=%s password=%s host=%s port=%d sslmode=disable",
		config.Info.DatabaseName,
		config.Info.DatabaseID,
		config.Info.DatabasePassword,
		config.Info.DatabaseIP,
		config.Info.DatabasePort,
	)

	return pgInfo
}
