package entity

import "encoding/xml"

type Evttsvtermino struct {
	XMLName           xml.Name          `xml:"evtTSVTermino"`
	Id                string            `xml:"Id"`
	Ideevento         Ideevento         `xml:"ideEvento"`
	Ideempregador     Ideempregador     `xml:"ideEmpregador"`
	Idetrabsemvinculo Idetrabsemvinculo `xml:"ideTrabSemVinculo"`
	Infotsvtermino    Infotsvtermino    `xml:"infoTSVTermino"`
}
