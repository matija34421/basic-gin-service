package server

import (
	"basic-gin/internal/config"
	"basic-gin/internal/db"
	"fmt"
	"log"
	"net/http"
)

func Start() error {
	config.LoadConfig()

	db.Init()

	port := config.AppConfig.SERVER_PORT

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, " Server is running")
	})

	log.Printf("🚀 Server listening on port %s", port)

	return http.ListenAndServe(":"+port, nil)
}
