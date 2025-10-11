package tipos_test

import (
	"encoding/xml"
	"testing"

	"github.com/avelinobego/esocial/internal/tipos"
)

func TestNumRecibo_MarshalXML(t *testing.T) {
	ele := tipos.NumRecibo{
		Ambiente:         tipos.Testes,
		NumeroSequencial: 1,
	}

	expected := "<NumRecibo>1.8.0000000000000000001</NumRecibo>"

	value, err := xml.Marshal(&ele)
	if err != nil {
		t.Fatalf("Erro ao serializar NumRecibo: %v", err)
	}

	t.Log(string(value))
	if string(value) != expected {
		t.Errorf("Esperado %s, mas obteve %s", expected, string(value))
	}
}

func TestNumRecibo_UnmarshalXML(t *testing.T) {
	data := "<NumRecibo>1.7.0000000000000000001</NumRecibo>"
	ele := tipos.NumRecibo{
		Ambiente:         tipos.Validacao,
		NumeroSequencial: 1,
	}

	err := xml.Unmarshal([]byte(data), &ele)
	if err != nil {
		t.Fatalf("Erro ao desserializar NumRecibo: %v", err)
	}

	if ele.Ambiente != tipos.Validacao || ele.NumeroSequencial != 1 {
		t.Errorf("Dados incorretos: Ambiente=%d, NumeroSequencial=%d", ele.Ambiente, ele.NumeroSequencial)
	}
}
