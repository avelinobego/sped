package entity

import "encoding/xml"

type Evtcdbenefin struct {
	XMLName       xml.Name      `xml:"evtCdBenefIn"`
	Id            string        `xml:"Id"`
	Ideevento     Ideevento     `xml:"ideEvento"`
	Ideempregador Ideempregador `xml:"ideEmpregador"`
	Beneficiario  Beneficiario  `xml:"beneficiario"`
}
