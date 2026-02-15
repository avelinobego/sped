package entity

import "encoding/xml"

type Evtreativben struct {
	XMLName       xml.Name      `xml:"evtReativBen"`
	Id            string        `xml:"Id"`
	Ideevento     Ideevento     `xml:"ideEvento"`
	Ideempregador Ideempregador `xml:"ideEmpregador"`
	Idebeneficio  Idebeneficio  `xml:"ideBeneficio"`
	Inforeativ    Inforeativ    `xml:"infoReativ"`
}
