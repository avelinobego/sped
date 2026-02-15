package entity

import "encoding/xml"

type Evtcontproc struct {
	XMLName       xml.Name      `xml:"evtContProc"`
	Id            string        `xml:"Id"`
	Ideevento     Ideevento     `xml:"ideEvento"`
	Ideempregador Ideempregador `xml:"ideEmpregador"`
	Ideproc       Ideproc       `xml:"ideProc"`
	Idetrab       []Idetrab     `xml:"ideTrab"`
}
