package tipos_test

import (
	"encoding/xml"
	"fmt"
	"testing"
	"time"

	"github.com/avelinobego/esocial/internal/tipos"
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
