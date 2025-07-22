package tipos_test

import (
	"avelinobego/esocial/tipos"
	"testing"
)

func TestCpf(t *testing.T) {
	t.Log("Testing Cpf...")
	cpf := tipos.Cpf("16944301890")
	err := cpf.Validar()
	if err != nil {
		t.Fatal(err)
	}
}
