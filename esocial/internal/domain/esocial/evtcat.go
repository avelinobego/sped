package esocial

import "encoding/xml"

type Evtcat struct {
	XMLName       xml.Name      `xml:"evtCAT"`
	Id            string        `xml:"Id"`
	Ideevento     Ideevento     `xml:"ideEvento"`
	Ideempregador Ideempregador `xml:"ideEmpregador"`
	Idevinculo    Idevinculo    `xml:"ideVinculo"`
	Cat           Cat           `xml:"cat"`
}
