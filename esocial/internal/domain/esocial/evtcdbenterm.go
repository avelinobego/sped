package esocial

import "encoding/xml"

type Evtcdbenterm struct {
	XMLName        xml.Name       `xml:"evtCdBenTerm"`
	Id             string         `xml:"Id"`
	Ideevento      Ideevento      `xml:"ideEvento"`
	Ideempregador  Ideempregador  `xml:"ideEmpregador"`
	Idebeneficio   Idebeneficio   `xml:"ideBeneficio"`
	Infobentermino Infobentermino `xml:"infoBenTermino"`
}
