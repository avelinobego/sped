package tipos_test

import (
	"testing"
	"time"

	"github.com/avelinobego/esocial/internal/tipos"
)

func TestId_NewId(t *testing.T) {
	cpfCnpj := tipos.CPF("37232795033")
	dataHora := time.Date(2024, 6, 20, 15, 30, 45, 0, time.UTC)
	sequencia := 1

	id, err := tipos.NewId(cpfCnpj, dataHora, sequencia)
	if err != nil {
		t.Fatalf("failed to create new ID: %v", err)
	}

	expected := "ID2000372327950332024062015304500001"

	if string(id) != expected {
		t.Errorf("expected %s, got %s", expected, string(id))
	}
}
