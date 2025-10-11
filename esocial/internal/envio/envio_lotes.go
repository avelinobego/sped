package envio

import "github.com/avelinobego/esocial/internal/tipos"

type EnvioLoteEventos struct {
	IdEmpregador  tipos.IdentificadorPorDocumento `xml:"ideEmpregador"`
	IdTransmissor tipos.IdentificadorPorDocumento `xml:"ideTransmissor"`
}
