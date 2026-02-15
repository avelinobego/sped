package entity

import "encoding/xml"

type Evtfgts struct {
	XMLName       xml.Name      `xml:"evtFGTS"`
	Id            string        `xml:"Id"`
	Ideevento     Ideevento     `xml:"ideEvento"`
	Ideempregador Ideempregador `xml:"ideEmpregador"`
	Infofgts      Infofgts      `xml:"infoFGTS"`
}
