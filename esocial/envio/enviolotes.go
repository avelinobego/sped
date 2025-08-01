package envio

import "avelinobego/esocial/tipos"

type EnvioLoteEventos struct {
	IdEmpregador  tipos.IdentificadorPorDocumento `xml:"ideEmpregador"`
	IdTransmissor tipos.IdentificadorPorDocumento `xml:"ideTransmissor"`
}
