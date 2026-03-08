package esocial

import "encoding/xml"

type Evtinfocomplper struct {
	XMLName             xml.Name              `xml:"evtInfoComplPer"`
	Id                  string                `xml:"Id"`
	Ideevento           Ideevento             `xml:"ideEvento"`
	Ideempregador       Ideempregador         `xml:"ideEmpregador"`
	Infosubstpatr       []Infosubstpatr       `xml:"infoSubstPatr,omitempty"`
	Infosubstpatropport []Infosubstpatropport `xml:"infoSubstPatrOpPort"`
	Infoativconcom      []Infoativconcom      `xml:"infoAtivConcom,omitempty"`
	Infoperctransf11096 []Infoperctransf11096 `xml:"infoPercTransf11096,omitempty"`
}
