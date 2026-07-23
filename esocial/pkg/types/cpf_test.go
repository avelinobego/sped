package types_test

import (
	"encoding/xml"
	"testing"
)

func TestCPF_MarshalUnmarshal(t *testing.T) {
	t.Run("Marshal Válido", func(t *testing.T) {
		w := WrapperAll{CPF: "123.456.789-01"}
		output, err := xml.Marshal(w)
		if err != nil {
			t.Fatalf("erro inesperado: %v", err)
		}
		expected := "<wrapper><cpf>12345678901</cpf></wrapper>"
		if string(output) != expected {
			t.Errorf("got %s, want %s", string(output), expected)
		}
	})

	t.Run("Marshal Tamanho Inválido", func(t *testing.T) {
		w := WrapperAll{CPF: "123.456"}
		_, err := xml.Marshal(w)
		if err == nil {
			t.Error("esperava erro de tamanho inválido, got nil")
		}
	})

	t.Run("Unmarshal Válido", func(t *testing.T) {
		data := []byte(`<wrapper><cpf>12345678901</cpf></wrapper>`)
		var w WrapperAll
		if err := xml.Unmarshal(data, &w); err != nil {
			t.Fatalf("erro inesperado: %v", err)
		}
		if string(w.CPF) != "12345678901" {
			t.Errorf("got %s, want %s", w.CPF, "12345678901")
		}
	})
}
