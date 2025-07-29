package retorno

import "avelinobego/esocial/tipos"

type RetornoEvento struct {
	Id            tipos.IdeEvento     `xml:"Id,attr"`
	IdeEmpregador tipos.IdeEmpregador `xml:"ideEmpregador"`
	Recepcao      Recepcao            `xml:"recepcao"`
	Processamento Processamento       `xml:"processamento"`
	Recibo        Recibo              `xml:"recibo,omitempty"`
}
