package esocial

import "encoding/xml"

type Evttabestab struct {
	XMLName       xml.Name      `xml:"evtTabEstab"`
	Id            string        `xml:"Id"`
	Ideevento     Ideevento     `xml:"ideEvento"`
	Ideempregador Ideempregador `xml:"ideEmpregador"`
	Infoestab     string        `xml:"infoEstab"`
}
