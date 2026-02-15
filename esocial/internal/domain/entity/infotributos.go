package entity

type Infotributos struct {
	Perref        string          `xml:"perRef"`
	Infocrcontrib []Infocrcontrib `xml:"infoCRContrib"`
}
