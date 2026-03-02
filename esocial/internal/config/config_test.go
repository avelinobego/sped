package config_test

import (
	"log/slog"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	slog.Info("Get environment variable from os env")
	m.Run()
}

func TestPropertiesonfig(t *testing.T) {
	slog.Info("Environment", "DATABASE_URL", os.Getenv("DATABASE_URL"))
}
