package entity

import "encoding/xml"

type Evtadmissao struct {
	XMLName       xml.Name      `xml:"evtAdmissao"`
	Id            string        `xml:"Id"`
	Ideevento     Ideevento     `xml:"ideEvento"`
	Ideempregador Ideempregador `xml:"ideEmpregador"`
	Trabalhador   Trabalhador   `xml:"trabalhador"`
	Vinculo       Vinculo       `xml:"vinculo"`
}
