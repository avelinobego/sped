package entity

import "encoding/xml"

type Evtcomprod struct {
	XMLName       xml.Name      `xml:"evtComProd"`
	Id            string        `xml:"Id"`
	Ideevento     Ideevento     `xml:"ideEvento"`
	Ideempregador Ideempregador `xml:"ideEmpregador"`
	Infocomprod   Infocomprod   `xml:"infoComProd"`
}
