package esocial

import "encoding/xml"

type Evtexclusao struct {
	XMLName       xml.Name      `xml:"evtExclusao"`
	Id            string        `xml:"Id"`
	Ideevento     Ideevento     `xml:"ideEvento"`
	Ideempregador Ideempregador `xml:"ideEmpregador"`
	Infoexclusao  Infoexclusao  `xml:"infoExclusao"`
}
