package configs_test

import (
	"os"
	"testing"

	"github.com/avelinobego/esocial/configs"
)

func TestLogger(t *testing.T) {
	log := configs.NewLogger(os.Stdout, os.Stdout, os.Stdout)

	log.Info("Teste", "ID", 12345)
	log.Warn("Teste", "ID", 12345)
	log.Error("Teste", "ID", 12345)
	log.Debug("Teste", "ID", 12345)
	log.Fatal("Teste", "ID", 12345)
}
