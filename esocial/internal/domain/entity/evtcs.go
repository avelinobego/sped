package entity

import "encoding/xml"

type Evtcs struct {
	XMLName       xml.Name      `xml:"evtCS"`
	Id            string        `xml:"Id"`
	Ideevento     Ideevento     `xml:"ideEvento"`
	Ideempregador Ideempregador `xml:"ideEmpregador"`
	Infocs        Infocs        `xml:"infoCS"`
}
