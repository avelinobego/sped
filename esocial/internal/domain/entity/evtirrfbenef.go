package entity

import "encoding/xml"

type Evtirrfbenef struct {
	XMLName        xml.Name       `xml:"evtIrrfBenef"`
	Id             string         `xml:"Id"`
	Ideevento      Ideevento      `xml:"ideEvento"`
	Ideempregador  Ideempregador  `xml:"ideEmpregador"`
	Idetrabalhador Idetrabalhador `xml:"ideTrabalhador"`
}
