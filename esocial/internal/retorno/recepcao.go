package retorno

import (
	"time"

	"github.com/avelinobego/esocial/internal/tipos"
)

type Recepcao struct {
	TipoAmbiente   tipos.TipoAmbiente `xml:"tpAmb"`
	DhRecepcao     time.Time          `xml:"dhRecepcao"`
	VersaoApp      string             `xml:"versaoAppRecepcao"`
	ProtocoloEnvio string             `xml:"protocoloEnvioLote,omitempty"`
}
