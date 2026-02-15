package entity

import "encoding/xml"

type Evttsvinicio struct {
	XMLName       xml.Name      `xml:"evtTSVInicio"`
	Id            string        `xml:"Id"`
	Ideevento     Ideevento     `xml:"ideEvento"`
	Ideempregador Ideempregador `xml:"ideEmpregador"`
	Trabalhador   Trabalhador   `xml:"trabalhador"`
	Infotsvinicio Infotsvinicio `xml:"infoTSVInicio"`
}
