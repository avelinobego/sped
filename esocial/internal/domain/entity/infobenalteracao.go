package entity

type Infobenalteracao struct {
	Dtaltbeneficio string         `xml:"dtAltBeneficio"`
	Dadosbeneficio Dadosbeneficio `xml:"dadosBeneficio"`
}
