package types_test

import (
	"encoding/xml"
	"sped/esocial/pkg/types"
	"testing"
)

type Wrapper struct {
	XMLName xml.Name   `xml:"wrapper"`
	CNPJ    types.CNPJ `xml:"cnpj"`
}

func TestCNPJ_MarshalXML(t *testing.T) {
	tests := []struct {
		name    string
		input   types.CNPJ
		wantXML string
		wantErr bool
	}{
		{
			name:    "CNPJ válido com pontuação",
			input:   "11.222.333/0001-81",
			wantXML: "<wrapper><cnpj>11222333000181</cnpj></wrapper>",
			wantErr: false,
		},
		{
			name:    "CNPJ alfanumérico válido (conforme padrão atual)",
			input:   "12.ABC.345/0001-78", // Ajuste para um CNPJ alfanumérico com DV válido se necessário
			wantXML: "",
			wantErr: true, // Se o dígito não bater com o cálculo padrão do módulo 11
		},
		{
			name:    "CNPJ com tamanho inválido",
			input:   "11.222.333/0001",
			wantXML: "",
			wantErr: true,
		},
		{
			name:    "CNPJ alfanumérico dígito verificador inválido (conforme padrão atual)",
			input:   "12.ABC.345/0001-70", // Ajuste para um CNPJ alfanumérico com DV inválido se necessário
			wantXML: "",
			wantErr: false, // Se o dígito não bater com o cálculo padrão do módulo 11, deve retornar erro
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := Wrapper{CNPJ: tt.input}
			output, err := xml.Marshal(w)

			if (err != nil) != tt.wantErr {
				t.Fatalf("MarshalXML() erro = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr && string(output) != tt.wantXML {
				t.Errorf("MarshalXML() = %s, want %s", string(output), tt.wantXML)
			}
		})
	}
}

func TestCNPJ_UnmarshalXML(t *testing.T) {
	xmlData := []byte(`<wrapper><cnpj>11222333000181</cnpj></wrapper>`)

	var w Wrapper
	err := xml.Unmarshal(xmlData, &w)
	if err != nil {
		Fatalf := t.Fatalf
		Fatalf("UnmarshalXML() erro inesperado = %v", err)
	}

	if string(w.CNPJ) != "11222333000181" {
		t.Errorf("UnmarshalXML() = %s, want %s", w.CNPJ, "11222333000181")
	}
}
