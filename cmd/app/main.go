package main

import (
	"log"

	"basic-gin/internal/server"
)

func main() {
	log.Println("Starting application...")

	if err := server.Start(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
