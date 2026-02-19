package esocial_test

import (
	"fmt"
	"testing"

	"github.com/avelinobego/esocial/internal/core/shared/esocial"
)

func TestS1000(t *testing.T) {

	iem := esocial.Infoempregador{
		Inclusao: &esocial.Inclusao{
			Ideperiodo: &esocial.Ideperiodo{
				Inivalid: "2026-02",
				Fimvalid: "2027-02",
			},
		},
	}

	evt := esocial.Evtinfoempregador{
		Infoempregador: iem,
	}

	esocial := esocial.NewESocial("v_1_3_0", evt)

	for range 3 {
		value, err := esocial.ToXML()
		if err != nil {
			t.Errorf("Error marshaling XML: %v", err)
		}
		fmt.Println("---------------------------")
		fmt.Println(value)
	}

}
