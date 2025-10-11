package eventos

import (
	"encoding/xml"
	"fmt"
)

type ESocial[T any] struct {
	XMLName xml.Name `xml:"eSocial"`
	XMLNS   string   `xml:"xmlns,attr"`
	Evento  T
}

func (es ESocial[T]) String() string {
	return fmt.Sprintf("ESocial: %v", es.Evento)
}
