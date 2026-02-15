package entity_test

import (
	"fmt"
	"testing"

	"github.com/avelinobego/esocial/internal/domain/entity"
)

func TestS1000(t *testing.T) {

	iem := entity.Infoempregador{
		Inclusao: &entity.Inclusao{
			Ideperiodo: &entity.Ideperiodo{
				Inivalid: "2026-02",
				Fimvalid: "2027-02",
			},
		},
	}
	evt := entity.Evtinfoempregador{
		Infoempregador: iem,
	}
	esocial := entity.ESocial{
		Versao: "v_01_03_00",
		Evento: evt,
	}

	bytes, err := esocial.ToXML()
	if err != nil {
		t.Errorf("Error marshaling XML: %v", err)
	}
	fmt.Println(bytes)
}
