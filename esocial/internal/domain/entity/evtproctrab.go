package entity

import "encoding/xml"

type Evtproctrab struct {
	XMLName       xml.Name      `xml:"evtProcTrab"`
	Id            string        `xml:"Id"`
	Ideevento     Ideevento     `xml:"ideEvento"`
	Ideempregador Ideempregador `xml:"ideEmpregador"`
	Infoprocesso  Infoprocesso  `xml:"infoProcesso"`
	Idetrab       Idetrab       `xml:"ideTrab"`
}
