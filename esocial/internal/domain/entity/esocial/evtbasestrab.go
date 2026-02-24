package esocial

import "encoding/xml"

type Evtbasestrab struct {
	XMLName        xml.Name       `xml:"evtBasesTrab"`
	Id             string         `xml:"Id"`
	Ideevento      Ideevento      `xml:"ideEvento"`
	Ideempregador  Ideempregador  `xml:"ideEmpregador"`
	Idetrabalhador Idetrabalhador `xml:"ideTrabalhador"`
	Infocpcalc     []Infocpcalc   `xml:"infoCpCalc"`
	Infocp         []Infocp       `xml:"infoCp,omitempty"`
	Infopispasep   []Infopispasep `xml:"infoPisPasep,omitempty"`
}
