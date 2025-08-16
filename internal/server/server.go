package server

import (
	"basic-gin/internal/config"
	"basic-gin/internal/db"
	"basic-gin/internal/handler"
	"basic-gin/internal/repository"
	"basic-gin/internal/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() error {
	config.LoadConfig()

	db.Init()

	clientRepo := repository.NewClientRepository(db.DB)
	clientService := service.NewClientService(clientRepo)
	clientHandler := handler.NewClientHandler(clientService)

	r := mux.NewRouter()

	r.HandleFunc("/clients", clientHandler.GetClients).Methods("GET")
	r.HandleFunc("/clients/{id}", clientHandler.GetClientById).Methods("GET")
	r.HandleFunc("/clients", clientHandler.CreateClient).Methods("POST")
	r.HandleFunc("/clients", clientHandler.UpdateClient).Methods("PUT")
	r.HandleFunc("/clients/{id}", clientHandler.DeleteClient).Methods("DELETE")

	port := config.AppConfig.SERVER_PORT
	log.Printf("🚀 Server listening on port %s", port)
	return http.ListenAndServe(":"+port, r)
}
