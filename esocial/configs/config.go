package configs

import (
	"log/slog"
	"os"
)

func init() {
	configureLogger()
}

func configureLogger() {
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})
	instance := slog.New(handler)

	// Set as the default logger for the slog package
	slog.SetDefault(instance)
}
