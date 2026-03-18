package esocial

import "encoding/xml"

type Evttabprocesso struct {
	XMLName       xml.Name      `xml:"evtTabProcesso"`
	Id            string        `xml:"Id"`
	Ideevento     Ideevento     `xml:"ideEvento"`
	Ideempregador Ideempregador `xml:"ideEmpregador"`
	Infoprocesso  string        `xml:"infoProcesso"`
}
