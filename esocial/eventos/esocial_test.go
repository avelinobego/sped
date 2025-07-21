package eventos_test

import (
	"avelinobego/esocial/eventos"
	"avelinobego/esocial/tipos"
	"encoding/xml"
	"testing"
	"time"
)

func TestEsocial(t *testing.T) {

	ip := tipos.IdePeriodo(time.Now())

	es := eventos.ESocial{
		XMLNS:  "http://www.esocial.gov.br/schema/evt/evtInfoEmpregador/vx_x_x",
		Evento: ip,
	}
	t.Log(es)

	value, err := xml.Marshal(es)
	if err != nil {
		t.Error(err)
	}
	t.Log(string(value))
}
