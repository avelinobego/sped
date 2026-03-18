package esocial

import "encoding/xml"

type Evtcdbenefalt struct {
	XMLName       xml.Name      `xml:"evtCdBenefAlt"`
	Id            string        `xml:"Id"`
	Ideevento     Ideevento     `xml:"ideEvento"`
	Ideempregador Ideempregador `xml:"ideEmpregador"`
	Idebenef      Idebenef      `xml:"ideBenef"`
	Alteracao     Alteracao     `xml:"alteracao"`
}
