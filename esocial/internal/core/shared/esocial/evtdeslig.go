package esocial

import "encoding/xml"

type Evtdeslig struct {
	XMLName       xml.Name      `xml:"evtDeslig"`
	Id            string        `xml:"Id"`
	Ideevento     Ideevento     `xml:"ideEvento"`
	Ideempregador Ideempregador `xml:"ideEmpregador"`
	Idevinculo    Idevinculo    `xml:"ideVinculo"`
	Infodeslig    Infodeslig    `xml:"infoDeslig"`
}
