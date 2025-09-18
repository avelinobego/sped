package retorno

import "github.com/avelinobego/esocial/internal/tipos"

type Recibo struct {
	NumeroRecibo tipos.NumRecibo `xml:"nrRecibo"`
	Hash         string          `xml:"hash"`
	Contrato     Contrato        `xml:"contrato,omitempty"`
}
