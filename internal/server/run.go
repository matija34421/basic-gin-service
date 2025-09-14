package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"basic-gin/internal/config"
	"basic-gin/internal/db"
	"basic-gin/internal/handler"
	"basic-gin/internal/repository"
	"basic-gin/internal/service"
)

func Run(ctx context.Context) error {
	config.LoadConfig()

	port := config.AppConfig.SERVER_PORT
	if port == "" {
		port = "8080"
	}

	if err := db.Init(ctx); err != nil {
		return fmt.Errorf("db init: %w", err)
	}

	clientRepo := repository.NewClientRepository(db.DB)
	accountRepo := repository.NewAccountRepository(db.DB)

	clientSvc := service.NewClientService(clientRepo)
	accountSvc := service.NewAccountService(accountRepo, clientRepo)

	clientH := handler.NewClientHandler(clientSvc)
	accountH := handler.NewAccountHandler(accountSvc)

	r := NewRouter(clientH, accountH)

	srv := &http.Server{
		Addr:              ":" + port,
		Handler:           r,
		ReadTimeout:       5 * time.Second,
		ReadHeaderTimeout: 3 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	srv.RegisterOnShutdown(func() {
		if db.DB != nil {
			_ = db.DB.Close()
		}
	})

	errCh := make(chan error, 1)
	go func() {
		log.Println("🚀 Server listening on", srv.Addr)
		errCh <- srv.ListenAndServe()
	}()

	select {
	case <-ctx.Done():
		log.Println("shutdown signal received")
		shCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := srv.Shutdown(shCtx); err != nil && err != http.ErrServerClosed {
			log.Printf("server shutdown error: %v", err)
		}
		return nil
	case err := <-errCh:
		if err != nil && err != http.ErrServerClosed {
			return err
		}
		return nil
	}
}
