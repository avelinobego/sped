package entity

import "encoding/xml"

type Evtcdbenin struct {
	XMLName       xml.Name      `xml:"evtCdBenIn"`
	Id            string        `xml:"Id"`
	Ideevento     Ideevento     `xml:"ideEvento"`
	Ideempregador Ideempregador `xml:"ideEmpregador"`
	Beneficiario  Beneficiario  `xml:"beneficiario"`
	Infobeninicio Infobeninicio `xml:"infoBenInicio"`
}
