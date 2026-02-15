package entity

import "encoding/xml"

type Evtrmnrpps struct {
	XMLName        xml.Name       `xml:"evtRmnRPPS"`
	Id             string         `xml:"Id"`
	Ideevento      Ideevento      `xml:"ideEvento"`
	Ideempregador  Ideempregador  `xml:"ideEmpregador"`
	Idetrabalhador Idetrabalhador `xml:"ideTrabalhador"`
	Dmdev          []Dmdev        `xml:"dmDev"`
}
