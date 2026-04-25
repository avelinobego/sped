package config

import (
	"log/slog"
	"os"
)

func init() {

	opts := &slog.HandlerOptions{
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			// 1. Rename 'level' to 'status' or 'severity' for Datadog
			if a.Key == slog.LevelKey {
				return slog.Attr{Key: "status", Value: a.Value}
			}
			// 2. Rename 'msg' to 'message' (Datadog's default message key)
			if a.Key == slog.MessageKey {
				return slog.Attr{Key: "message", Value: a.Value}
			}
			return a
		},
	}

	// 1. Create a JSON handler that writes to Stdout
	handler := slog.NewJSONHandler(os.Stdout, opts)

	// 2. Create a new logger with that handler
	logger := slog.New(handler)

	// 3. (Optional) Set it as the global default
	slog.SetDefault(logger)
}
