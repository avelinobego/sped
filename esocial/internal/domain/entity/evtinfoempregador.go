package entity

import "encoding/xml"

type Evtinfoempregador struct {
	XMLName        xml.Name       `xml:"evtInfoEmpregador"`
	Id             string         `xml:"Id,attr"`
	Ideevento      Ideevento      `xml:"ideEvento"`
	Ideempregador  Ideempregador  `xml:"ideEmpregador"`
	Infoempregador Infoempregador `xml:"infoEmpregador"`
}
