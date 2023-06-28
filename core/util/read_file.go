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

	file, err := os.ReadFile(filepath.Join(path, "PostgresqlConfig.json"))
	if err != nil {
		log.Fatalf("Unable to read config file: %v", err)
	}

	var config PostgresInformation
	err = json.Unmarshal(file, &config)
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
