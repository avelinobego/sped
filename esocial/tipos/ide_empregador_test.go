package tipos_test

import (
	"avelinobego/esocial/tipos"
	"encoding/xml"
	"testing"
)

func TestIdeEmpregadorValidar(t *testing.T) {
	cnpj := tipos.CNPJ("12345678000195")
	idee := tipos.IdentificadorPorDocumento{
		Tipo: &cnpj, // Valid CNPJ
	}
	if err := idee.Validar(); err != nil {
		t.Errorf("Erro ao validar IdeEmpregador: %v", err)
	}
}

func TestIdeEmpregadorXML(t *testing.T) {
	var cnpj tipos.CNPJ = "12345678000195"
	idee := tipos.IdentificadorPorDocumento{
		Tipo: &cnpj, // Valid CNPJtab color
	}

	data, err := xml.Marshal(idee)
	if err != nil {
		t.Fatalf("Erro ao serializar IdeEmpregador: %v", err)
	}

	expected := `<ideEmpregador><nrInsc>12345678000195</nrInsc><tpInsc>1</tpInsc></ideEmpregador>`
	if string(data) != expected {
		t.Errorf("Dados XML inesperados: got %s, want %s", data, expected)
	}
}
