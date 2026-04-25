package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"sped/esocial/internal/controller"
	_ "sped/esocial/pkg/config"
)

func main() {

	server := &http.Server{
		Handler:      controller.NewHandlers(),
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	go func(s *http.Server) {
		slog.Info("Starting eSocial application...")
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("Failed to start server", "error", err)
		}
	}(server)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop
	slog.Info("Shutting down eSocial application...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.ErrorContext(ctx, "Failed to shutdown server", "error", err)
	} else {
		slog.InfoContext(ctx, "Server shutdown successfully")
	}

}
