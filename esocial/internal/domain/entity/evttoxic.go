package entity

import "encoding/xml"

type Evttoxic struct {
	XMLName       xml.Name      `xml:"evtToxic"`
	Id            string        `xml:"Id"`
	Ideevento     Ideevento     `xml:"ideEvento"`
	Ideempregador Ideempregador `xml:"ideEmpregador"`
	Idevinculo    Idevinculo    `xml:"ideVinculo"`
	Toxicologico  Toxicologico  `xml:"toxicologico"`
}
