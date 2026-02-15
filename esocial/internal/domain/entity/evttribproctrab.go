package entity

import "encoding/xml"

type Evttribproctrab struct {
	XMLName       xml.Name      `xml:"evtTribProcTrab"`
	Id            string        `xml:"Id"`
	Ideevento     Ideevento     `xml:"ideEvento"`
	Ideempregador Ideempregador `xml:"ideEmpregador"`
	Ideproc       Ideproc       `xml:"ideProc"`
}
