package esocial

import "encoding/xml"

type Evtpgtos struct {
	XMLName       xml.Name      `xml:"evtPgtos"`
	Id            string        `xml:"Id"`
	Ideevento     Ideevento     `xml:"ideEvento"`
	Ideempregador Ideempregador `xml:"ideEmpregador"`
	Idebenef      Idebenef      `xml:"ideBenef"`
}
