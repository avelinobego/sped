package retorno

import "avelinobego/esocial/tipos"

type Recibo struct {
	NumeroRecibo tipos.NumRecibo `xml:"nrRecibo"`
	Hash         string          `xml:"hash"`
	Contrato     Contrato        `xml:"contrato,omitempty"`
}
