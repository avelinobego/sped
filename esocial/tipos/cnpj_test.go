package tipos_test

import (
	"avelinobego/esocial/tipos"
	"errors"
	"testing"
)

func TestCNPJValidar(t *testing.T) {
	tests := []struct {
		cnpj     string
		expected error
	}{
		{"12345678000195", nil},
		{"12345678000196", errors.New("CNPJ inválido")},
		{"11111111000191", errors.New("CNPJ inválido")},
		{"194988750001930", errors.New("CNPJ deve ter 14 dígitos")},
	}

	for _, test := range tests {
		cnpj := tipos.CNPJ(test.cnpj)
		err := cnpj.Validar()

		if err == nil {
			continue
		}

		if test.expected == nil {
			continue
		}

		if err.Error() != test.expected.Error() {
			t.Errorf("CNPJ %s: expected %v, got %v", test.cnpj, test.expected, err)
		}
	}
}
