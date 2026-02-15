package entity

import "encoding/xml"

type Evttabrubrica struct {
	XMLName       xml.Name      `xml:"evtTabRubrica"`
	Id            string        `xml:"Id"`
	Ideevento     Ideevento     `xml:"ideEvento"`
	Ideempregador Ideempregador `xml:"ideEmpregador"`
	Inforubrica   string        `xml:"infoRubrica"`
}
