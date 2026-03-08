package esocial

import "encoding/xml"

type Evtmonit struct {
	XMLName       xml.Name      `xml:"evtMonit"`
	Id            string        `xml:"Id"`
	Ideevento     Ideevento     `xml:"ideEvento"`
	Ideempregador Ideempregador `xml:"ideEmpregador"`
	Idevinculo    Idevinculo    `xml:"ideVinculo"`
	Exmedocup     Exmedocup     `xml:"exMedOcup"`
}
