package entity

type Calctrib struct {
	Perref        string          `xml:"perRef"`
	Vrbccpmensal  float64         `xml:"vrBcCpMensal"`
	Vrbccp13      float64         `xml:"vrBcCp13"`
	Infocrcontrib []Infocrcontrib `xml:"infoCRContrib"`
}
