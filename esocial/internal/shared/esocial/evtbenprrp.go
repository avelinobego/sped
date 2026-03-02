package esocial

import "encoding/xml"

type Evtbenprrp struct {
	XMLName       xml.Name      `xml:"evtBenPrRP"`
	Id            string        `xml:"Id"`
	Ideevento     Ideevento     `xml:"ideEvento"`
	Ideempregador Ideempregador `xml:"ideEmpregador"`
	Idebenef      Idebenef      `xml:"ideBenef"`
	Dmdev         []Dmdev       `xml:"dmDev"`
}
