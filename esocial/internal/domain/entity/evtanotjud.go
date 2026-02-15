package entity

import "encoding/xml"

type Evtanotjud struct {
	XMLName       xml.Name      `xml:"evtAnotJud"`
	Id            string        `xml:"Id"`
	Ideevento     Ideevento     `xml:"ideEvento"`
	Ideempregador Ideempregador `xml:"ideEmpregador"`
	Infoprocesso  Infoprocesso  `xml:"infoProcesso"`
	Infoanotjud   Infoanotjud   `xml:"infoAnotJud"`
}
