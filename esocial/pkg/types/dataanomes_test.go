package types_test

import (
	"encoding/xml"
	"sped/esocial/pkg/types"
	"testing"
	"time"
)

type WrapperDataAnoMes struct {
	XMLName    xml.Name         `xml:"wrapper"`
	DataAnoMes types.DataAnoMes `xml:"dataAnoMes"`
}

func TestDataAnoMes_MarshalXML(t *testing.T) {
	fixedTime := time.Date(2026, 6, 1, 0, 0, 0, 0, time.UTC)

	tests := []struct {
		name    string
		input   types.DataAnoMes
		wantXML string
		wantErr bool
	}{
		{
			name:    "DataAnoMes válida",
			input:   types.DataAnoMes(fixedTime),
			wantXML: "<wrapper><dataAnoMes>2026-06</dataAnoMes></wrapper>",
			wantErr: false,
		},
		{
			name:    "DataAnoMes zero",
			input:   types.DataAnoMes(time.Time{}),
			wantXML: "<wrapper></wrapper>",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := WrapperDataAnoMes{DataAnoMes: tt.input}
			output, err := xml.Marshal(w)

			if (err != nil) != tt.wantErr {
				t.Fatalf("MarshalXML() erro = %v, wantErr %v", err, tt.wantErr)
			}

			if string(output) != tt.wantXML {
				t.Errorf("MarshalXML() = %s, want %s", string(output), tt.wantXML)
			}
		})
	}
}

func TestDataAnoMes_UnmarshalXML(t *testing.T) {
	xmlData := []byte(`<wrapper><dataAnoMes>2026-06</dataAnoMes></wrapper>`)

	var w WrapperDataAnoMes
	err := xml.Unmarshal(xmlData, &w)
	if err != nil {
		t.Fatalf("UnmarshalXML() erro inesperado = %v", err)
	}

	expected := time.Date(2026, 6, 1, 0, 0, 0, 0, time.UTC)
	got := time.Time(w.DataAnoMes)

	if !got.Equal(expected) {
		t.Errorf("UnmarshalXML() = %v, want %v", got, expected)
	}
}
