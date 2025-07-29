package retorno

import (
	"avelinobego/esocial/tipos"
	"encoding/xml"
)

type Contrato struct {
	XMLName         xml.Name              `xml:"contrato"`
	IdeEmpregador   tipos.IdeEmpregador   `xml:"ideEpregador"`
	IdeTrabalhador  tipos.IdeTrabalhador  `xml:"trabalhador"`
	InfoDeficiencia tipos.InfoDeficiencia `xml:"infoDeficiencia,omitempty"`
	Vinculo         tipos.Vinculo         `xml:"vinculo"`
}
