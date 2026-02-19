package esocial

import "encoding/xml"

type Evtcessao struct {
	XMLName       xml.Name      `xml:"evtCessao"`
	Id            string        `xml:"Id"`
	Ideevento     Ideevento     `xml:"ideEvento"`
	Ideempregador Ideempregador `xml:"ideEmpregador"`
	Idevinculo    Idevinculo    `xml:"ideVinculo"`
	Infocessao    string        `xml:"infoCessao"`
}
