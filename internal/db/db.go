package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"

	"basic-gin/internal/config"
)

var DB *sql.DB

func resolveDSN() string {
	if v := os.Getenv("DB_DSN"); v != "" {
		return v
	}
	if v := os.Getenv("DATABASE_URL"); v != "" {
		return v
	}
	if v := os.Getenv("DB_URL"); v != "" {
		return v
	}
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.AppConfig.DB_HOST,
		config.AppConfig.DB_PORT,
		config.AppConfig.DB_USER,
		config.AppConfig.DB_PASS,
		config.AppConfig.DB_NAME,
	)
}

func Init(ctx context.Context) error {
	dsn := resolveDSN()

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return fmt.Errorf("open db: %w", err)
	}

	db.SetMaxOpenConns(16)
	db.SetMaxIdleConns(16)
	db.SetConnMaxLifetime(30 * time.Minute)

	pctx, cancel := context.WithTimeout(ctx, 7*time.Second)
	defer cancel()
	if err := db.PingContext(pctx); err != nil {
		_ = db.Close()
		return fmt.Errorf("ping db: %w", err)
	}

	DB = db
	log.Print("✅ Connected to database!")
	return nil
}
