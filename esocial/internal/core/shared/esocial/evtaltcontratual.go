package esocial

import "encoding/xml"

type Evtaltcontratual struct {
	XMLName       xml.Name      `xml:"evtAltContratual"`
	Id            string        `xml:"Id"`
	Ideevento     Ideevento     `xml:"ideEvento"`
	Ideempregador Ideempregador `xml:"ideEmpregador"`
	Idevinculo    Idevinculo    `xml:"ideVinculo"`
	Altcontratual Altcontratual `xml:"altContratual"`
}
