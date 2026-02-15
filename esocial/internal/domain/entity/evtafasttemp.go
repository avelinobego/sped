package entity

import "encoding/xml"

type Evtafasttemp struct {
	XMLName         xml.Name        `xml:"evtAfastTemp"`
	Id              string          `xml:"Id"`
	Ideevento       Ideevento       `xml:"ideEvento"`
	Ideempregador   Ideempregador   `xml:"ideEmpregador"`
	Idevinculo      Idevinculo      `xml:"ideVinculo"`
	Infoafastamento Infoafastamento `xml:"infoAfastamento"`
}
