package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"sped/esocial/internal/handlers"
	_ "sped/esocial/pkg/config"
)

func main() {

	server := &http.Server{
		Handler:      handlers.Handlers(),
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	serverErrors := make(chan error, 1)

	go func(e chan<- error) {
		slog.Info("Starting eSocial application...")
		e <- server.ListenAndServe()
	}(serverErrors)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGHUP)

	select {
	case err := <-serverErrors:
		if err != nil && err != http.ErrServerClosed {
			slog.Error("Failed to start server", "error", err)
		}
	case s := <-stop:
		slog.Info("Shutting down eSocial application...", "signal", s)
		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			slog.Error("Failed to shutdown server", "error", err)
			if err := server.Close(); err != nil {
				slog.Error("Failed to close server", "error", err)
			}
		} else {
			slog.Info("Server shutdown successfully")
		}
	}

}
