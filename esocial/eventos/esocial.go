package eventos

import (
	"encoding/xml"
)

type ESocial struct {
	XMLName xml.Name `xml:"eSocial"`
	XMLNS   string   `xml:"xmlns,attr"`
	Evento  interface{}
}

func (es ESocial) String() string {
	return "ESocial"
}
