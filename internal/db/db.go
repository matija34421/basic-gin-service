package db

import (
	"basic-gin/internal/config"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Init() {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.AppConfig.DB_HOST,
		config.AppConfig.DB_PORT,
		config.AppConfig.DB_USER,
		config.AppConfig.DB_PASS,
		config.AppConfig.DB_NAME,
	)

	var err error

	DB, err = sql.Open("postgres", connStr)

	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	log.Print("✅ Connected to database!")
}
