package entity

import "encoding/xml"

type Evtcdbenalt struct {
	XMLName          xml.Name         `xml:"evtCdBenAlt"`
	Id               string           `xml:"Id"`
	Ideevento        Ideevento        `xml:"ideEvento"`
	Ideempregador    Ideempregador    `xml:"ideEmpregador"`
	Idebeneficio     Idebeneficio     `xml:"ideBeneficio"`
	Infobenalteracao Infobenalteracao `xml:"infoBenAlteracao"`
}
