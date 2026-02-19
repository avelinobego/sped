package esocial

import "encoding/xml"

type Evttsvaltcontr struct {
	XMLName           xml.Name          `xml:"evtTSVAltContr"`
	Id                string            `xml:"Id"`
	Ideevento         Ideevento         `xml:"ideEvento"`
	Ideempregador     Ideempregador     `xml:"ideEmpregador"`
	Idetrabsemvinculo Idetrabsemvinculo `xml:"ideTrabSemVinculo"`
	Infotsvalteracao  Infotsvalteracao  `xml:"infoTSVAlteracao"`
}
