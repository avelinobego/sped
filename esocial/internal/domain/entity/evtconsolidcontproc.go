package entity

import "encoding/xml"

type Evtconsolidcontproc struct {
	XMLName       xml.Name      `xml:"evtConsolidContProc"`
	Id            string        `xml:"Id"`
	Ideevento     Ideevento     `xml:"ideEvento"`
	Ideempregador Ideempregador `xml:"ideEmpregador"`
	Ideproc       Ideproc       `xml:"ideProc"`
}
