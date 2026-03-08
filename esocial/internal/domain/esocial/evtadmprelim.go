package esocial

import "encoding/xml"

type Evtadmprelim struct {
	XMLName       xml.Name      `xml:"evtAdmPrelim"`
	Id            string        `xml:"Id"`
	Ideevento     Ideevento     `xml:"ideEvento"`
	Ideempregador Ideempregador `xml:"ideEmpregador"`
	Inforegprelim Inforegprelim `xml:"infoRegPrelim"`
}
