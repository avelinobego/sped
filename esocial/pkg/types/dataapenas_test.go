package types_test

import (
	"encoding/xml"
	"sped/esocial/pkg/types"
	"testing"
	"time"
)

func TestDataApenas_MarshalUnmarshal(t *testing.T) {
	t.Run("Marshal", func(t *testing.T) {
		fixedTime := time.Date(2026, 6, 15, 0, 0, 0, 0, time.UTC)
		w := WrapperAll{DataApenas: types.DataApenas(fixedTime)}
		output, err := xml.Marshal(w)
		if err != nil {
			t.Fatalf("erro inesperado: %v", err)
		}
		expected := "<wrapper><dataApenas>2026-06-15</dataApenas></wrapper>"
		if string(output) != expected {
			t.Errorf("got %s, want %s", string(output), expected)
		}
	})

	t.Run("Unmarshal", func(t *testing.T) {
		data := []byte(`<wrapper><dataApenas>2026-06-15</dataApenas></wrapper>`)
		var w WrapperAll
		if err := xml.Unmarshal(data, &w); err != nil {
			t.Fatalf("erro inesperado: %v", err)
		}
		expected := time.Date(2026, 6, 15, 0, 0, 0, 0, time.UTC)
		if !time.Time(w.DataApenas).Equal(expected) {
			t.Errorf("got %v, want %v", time.Time(w.DataApenas), expected)
		}
	})
}
