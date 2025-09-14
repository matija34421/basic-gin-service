package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	SERVER_PORT string
	DB_PORT     string
	DB_USER     string
	DB_PASS     string
	DB_HOST     string
	DB_NAME     string
}

var AppConfig Config

func LoadConfig() {
	if err := godotenv.Load(); err != nil {
		log.Println(".env file not found, using system environment variables")
	}

	pass := firstNonEmpty(os.Getenv("DB_PASS"), os.Getenv("DB_PASSWORD"))

	AppConfig = Config{
		SERVER_PORT: getEnv("SERVER_PORT", "8080"),
		DB_PORT:     getEnv("DB_PORT", "5432"),
		DB_HOST:     getEnv("DB_HOST", "localhost"),
		DB_USER:     getEnv("DB_USER", "postgres"),
		DB_PASS:     pass,
		DB_NAME:     getEnv("DB_NAME", "basic_gin"),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func firstNonEmpty(vals ...string) string {
	for _, v := range vals {
		if v != "" {
			return v
		}
	}
	return ""
}
