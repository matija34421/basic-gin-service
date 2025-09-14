package server

import (
	"context"
	"log"
	"net/http"
	"runtime/debug"
	"time"

	"basic-gin/internal/handler"

	"github.com/gorilla/mux"
)

func NewRouter(clientH *handler.ClientHandler, accountH *handler.AccountHandler) http.Handler {
	r := mux.NewRouter()

	r.Use(RecoverMiddleware)
	r.Use(RequestTimeout(2 * time.Second))

	r.HandleFunc("/clients", clientH.GetClients).Methods(http.MethodGet)
	r.HandleFunc("/clients/{id}", clientH.GetClientById).Methods(http.MethodGet)
	r.HandleFunc("/clients", clientH.CreateClient).Methods(http.MethodPost)
	r.HandleFunc("/clients", clientH.UpdateClient).Methods(http.MethodPut)
	r.HandleFunc("/clients/{id}", clientH.DeleteClient).Methods(http.MethodDelete)

	r.HandleFunc("/clients/{id}/accounts", accountH.GetAccountsByClientId).Methods(http.MethodGet)
	r.HandleFunc("/accounts/{id}", accountH.GetAccountById).Methods(http.MethodGet)
	r.HandleFunc("/accounts", accountH.CreateAccount).Methods(http.MethodPost)
	r.HandleFunc("/accounts", accountH.UpdateAccount).Methods(http.MethodPut)

	return r
}

func RequestTimeout(d time.Duration) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx, cancel := context.WithTimeout(r.Context(), d)
			defer cancel()
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func RecoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				log.Printf("panic: %v\n%s", rec, debug.Stack())
				http.Error(w, "internal server error", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
