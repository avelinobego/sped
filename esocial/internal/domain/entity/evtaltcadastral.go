package entity

import "encoding/xml"

type Evtaltcadastral struct {
	XMLName        xml.Name       `xml:"evtAltCadastral"`
	Id             string         `xml:"Id"`
	Ideevento      Ideevento      `xml:"ideEvento"`
	Ideempregador  Ideempregador  `xml:"ideEmpregador"`
	Idetrabalhador Idetrabalhador `xml:"ideTrabalhador"`
	Alteracao      Alteracao      `xml:"alteracao"`
}
