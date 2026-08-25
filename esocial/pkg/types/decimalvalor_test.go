package types_test

import (
	"encoding/xml"
	"sped/esocial/pkg/types"
	"testing"
)

func TestDecimalValor_Marshal(t *testing.T) {
	w := WrapperAll{
		Decimal: types.DecimalValor{Value: 1234.567, Casas: 2},
	}
	output, err := xml.Marshal(w)
	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}
	expected := "<wrapper><decimal>1234.57</decimal></wrapper>"
	if string(output) != expected {
		t.Errorf("got %s, want %s", string(output), expected)
	}
}
