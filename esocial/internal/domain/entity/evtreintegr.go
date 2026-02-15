package entity

import "encoding/xml"

type Evtreintegr struct {
	XMLName       xml.Name      `xml:"evtReintegr"`
	Id            string        `xml:"Id"`
	Ideevento     Ideevento     `xml:"ideEvento"`
	Ideempregador Ideempregador `xml:"ideEmpregador"`
	Idevinculo    Idevinculo    `xml:"ideVinculo"`
	Inforeintegr  Inforeintegr  `xml:"infoReintegr"`
}
