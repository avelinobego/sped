package eventos_test

import (
	"encoding/xml"
	"testing"
	"time"

	"github.com/avelinobego/esocial/internal/eventos"
	"github.com/avelinobego/esocial/internal/tipos"
)

func TestEsocial(t *testing.T) {

	ip := tipos.IdePeriodo(time.Now())

	es := eventos.ESocial[tipos.IdePeriodo]{
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
