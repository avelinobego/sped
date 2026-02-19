package esocial

import "encoding/xml"

type Evtexprisco struct {
	XMLName       xml.Name      `xml:"evtExpRisco"`
	Id            string        `xml:"Id"`
	Ideevento     Ideevento     `xml:"ideEvento"`
	Ideempregador Ideempregador `xml:"ideEmpregador"`
	Idevinculo    Idevinculo    `xml:"ideVinculo"`
	Infoexprisco  Infoexprisco  `xml:"infoExpRisco"`
}
