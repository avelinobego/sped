package entity

type Trabtemporario struct {
	Hipleg             int64                `xml:"hipLeg"`
	Justcontr          string               `xml:"justContr"`
	Ideestabvinc       Ideestabvinc         `xml:"ideEstabVinc"`
	Idetrabsubstituido []Idetrabsubstituido `xml:"ideTrabSubstituido"`
	Justprorr          string               `xml:"justProrr"`
}
