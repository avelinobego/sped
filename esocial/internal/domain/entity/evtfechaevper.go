package entity

import "encoding/xml"

type Evtfechaevper struct {
	XMLName       xml.Name      `xml:"evtFechaEvPer"`
	Id            string        `xml:"Id"`
	Ideevento     Ideevento     `xml:"ideEvento"`
	Ideempregador Ideempregador `xml:"ideEmpregador"`
	Infofech      Infofech      `xml:"infoFech"`
}
