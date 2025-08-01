package tipos_test

import (
	"avelinobego/esocial/tipos"
	"encoding/xml"
	"fmt"
	"testing"
	"time"
)

func TestCreateIdePeriodo(t *testing.T) {
	hoje := time.Now()
	ide := tipos.IdePeriodo(hoje)
	t.Log(ide)

	value, _ := xml.MarshalIndent(ide, "", " ")
	t.Log(string(value))

	xml_str := "<IdePeriodo>2024-01</IdePeriodo>"
	var retorno tipos.IdePeriodo
	xml.Unmarshal([]byte(xml_str), &retorno)
	fmt.Println(retorno)
}
