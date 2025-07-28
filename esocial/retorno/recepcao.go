package retorno

import (
	"avelinobego/esocial/tipos"
	"time"
)

type Recepcao struct {
	TipoAmbiente   tipos.TipoAmbiente `xml:"tpAmb"`
	DhRecepcao     time.Time          `xml:"dhRecepcao"`
	VersaoApp      string             `xml:"versaoAppRecepcao"`
	ProtocoloEnvio string             `xml:"protocoloEnvioLote,omitempty"`
	Ocorrencias    []Ocorrencia       `xml:"ocorrencia,omitempty"`
	// Recibo Recibo `xml:"recibo,omitempty"``
}
