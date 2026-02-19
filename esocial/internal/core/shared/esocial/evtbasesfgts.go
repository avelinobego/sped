package esocial

import "encoding/xml"

type Evtbasesfgts struct {
	XMLName        xml.Name       `xml:"evtBasesFGTS"`
	Id             string         `xml:"Id"`
	Ideevento      Ideevento      `xml:"ideEvento"`
	Ideempregador  Ideempregador  `xml:"ideEmpregador"`
	Idetrabalhador Idetrabalhador `xml:"ideTrabalhador"`
	Infofgts       Infofgts       `xml:"infoFGTS"`
}
