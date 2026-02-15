package entity

import "encoding/xml"

type Evtexcproctrab struct {
	XMLName       xml.Name      `xml:"evtExcProcTrab"`
	Id            string        `xml:"Id"`
	Ideevento     Ideevento     `xml:"ideEvento"`
	Ideempregador Ideempregador `xml:"ideEmpregador"`
	Infoexclusao  Infoexclusao  `xml:"infoExclusao"`
}
