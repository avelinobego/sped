package entity

import "encoding/xml"

type Evttablotacao struct {
	XMLName       xml.Name      `xml:"evtTabLotacao"`
	Id            string        `xml:"Id"`
	Ideevento     Ideevento     `xml:"ideEvento"`
	Ideempregador Ideempregador `xml:"ideEmpregador"`
	Infolotacao   string        `xml:"infoLotacao"`
}
