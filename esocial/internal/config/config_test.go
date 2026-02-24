package config_test

import (
	"log/slog"
	"os"
	"testing"
)

func TestPropertiesonfig(t *testing.T) {
	slog.Info("Environment", "DATABASE_URL", os.Getenv("DATABASE_URL"))
}
