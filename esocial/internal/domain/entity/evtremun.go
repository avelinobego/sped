package entity

import "encoding/xml"

type Evtremun struct {
	XMLName        xml.Name       `xml:"evtRemun"`
	Id             string         `xml:"Id"`
	Ideevento      Ideevento      `xml:"ideEvento"`
	Ideempregador  Ideempregador  `xml:"ideEmpregador"`
	Idetrabalhador Idetrabalhador `xml:"ideTrabalhador"`
	Dmdev          []Dmdev        `xml:"dmDev"`
}
