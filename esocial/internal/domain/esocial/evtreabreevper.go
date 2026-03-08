package esocial

import "encoding/xml"

type Evtreabreevper struct {
	XMLName       xml.Name      `xml:"evtReabreEvPer"`
	Id            string        `xml:"Id"`
	Ideevento     Ideevento     `xml:"ideEvento"`
	Ideempregador Ideempregador `xml:"ideEmpregador"`
}
