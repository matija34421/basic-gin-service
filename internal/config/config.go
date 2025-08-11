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

var AppConfig *Config

func LoadConfig() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(".env file not found, using system environment variables")
	}

	AppConfig = &Config{
		SERVER_PORT: getEnv("SERVER_PORT", "8080"),
		DB_PORT:     getEnv("DB_PORT", "5432"),
		DB_HOST:     getEnv("DB_HOST", "localhost"),
		DB_USER:     getEnv("DB_USER", ""),
		DB_PASS:     getEnv("DB_PASSWORD", ""),
		DB_NAME:     getEnv("DB_NAME", ""),
	}
}

func getEnv(key, fallback string) string {
	if v, exists := os.LookupEnv(key); exists {
		return v
	}

	return fallback
}
