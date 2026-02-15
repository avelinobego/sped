package entity

import "encoding/xml"

type Evtfgtsproctrab struct {
	XMLName        xml.Name       `xml:"evtFGTSProcTrab"`
	Id             string         `xml:"Id"`
	Ideevento      Ideevento      `xml:"ideEvento"`
	Ideempregador  Ideempregador  `xml:"ideEmpregador"`
	Ideproc        Ideproc        `xml:"ideProc"`
	Idetrabalhador Idetrabalhador `xml:"ideTrabalhador"`
	Infotrabfgts   []Infotrabfgts `xml:"infoTrabFGTS"`
}
