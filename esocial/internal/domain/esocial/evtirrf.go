package esocial

import "encoding/xml"

type Evtirrf struct {
	XMLName       xml.Name      `xml:"evtIrrf"`
	Id            string        `xml:"Id"`
	Ideevento     Ideevento     `xml:"ideEvento"`
	Ideempregador Ideempregador `xml:"ideEmpregador"`
	Infoirrf      Infoirrf      `xml:"infoIRRF"`
}
