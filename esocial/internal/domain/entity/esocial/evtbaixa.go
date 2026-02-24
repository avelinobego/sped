package esocial

import "encoding/xml"

type Evtbaixa struct {
	XMLName       xml.Name      `xml:"evtBaixa"`
	Id            string        `xml:"Id"`
	Ideevento     Ideevento     `xml:"ideEvento"`
	Ideempregador Ideempregador `xml:"ideEmpregador"`
	Idevinculo    Idevinculo    `xml:"ideVinculo"`
	Infobaixa     Infobaixa     `xml:"infoBaixa"`
}
