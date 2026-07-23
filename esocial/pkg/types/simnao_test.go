package types_test

import (
	"encoding/xml"
	"sped/esocial/pkg/types"
	"testing"
)

func TestSimNao_MarshalUnmarshal(t *testing.T) {
	t.Run("Marshal Válido", func(t *testing.T) {
		w := WrapperAll{SimNao: types.Sim}
		output, err := xml.Marshal(w)
		if err != nil {
			t.Fatalf("erro inesperado: %v", err)
		}
		expected := "<wrapper><simNao>1</simNao></wrapper>"
		if string(output) != expected {
			t.Errorf("got %s, want %s", string(output), expected)
		}
	})

	t.Run("Marshal Inválido", func(t *testing.T) {
		w := WrapperAll{SimNao: types.SimNao(99)}
		_, err := xml.Marshal(w)
		if err == nil {
			t.Error("esperava erro para valor inválido, got nil")
		}
	})

	t.Run("Unmarshal Válido", func(t *testing.T) {
		data := []byte(`<wrapper><simNao>2</simNao></wrapper>`)
		var w WrapperAll
		if err := xml.Unmarshal(data, &w); err != nil {
			t.Fatalf("erro inesperado: %v", err)
		}
		if w.SimNao != types.Nao {
			t.Errorf("got %v, want %v", w.SimNao, types.Nao)
		}
	})
}
