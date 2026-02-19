package esocial

import "encoding/xml"

type Evtcontratavnp struct {
	XMLName       xml.Name      `xml:"evtContratAvNP"`
	Id            string        `xml:"Id"`
	Ideevento     Ideevento     `xml:"ideEvento"`
	Ideempregador Ideempregador `xml:"ideEmpregador"`
	Remunavnp     []Remunavnp   `xml:"remunAvNP"`
}
