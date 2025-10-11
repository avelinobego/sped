package tipos_test

import (
	"encoding/xml"
	"testing"

	"github.com/avelinobego/esocial/internal/tipos"
)

func TestCpf(t *testing.T) {
	t.Log("Testing Cpf...")
	cpf := tipos.CPF("16944301890")
	err := cpf.Validar()
	if err != nil {
		t.Fatal(err)
	}
}

func TestCpfXml(t *testing.T) {
	t.Log("Testing Cpf to XML...")
	cpf := tipos.CPF("16944301890")
	value, err := xml.Marshal(cpf)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("==>", string(value))
}

func TestXmlCpf(t *testing.T) {
	t.Log("Testing XML to Cpf...")
	value := "<Cpf>16944301890</Cpf>"
	var cpf tipos.CPF
	err := xml.Unmarshal([]byte(value), &cpf)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("==>", string(cpf))
}
