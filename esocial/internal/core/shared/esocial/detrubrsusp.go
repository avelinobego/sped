package esocial

type Detrubrsusp struct {
	Codrubr         string            `xml:"codRubr"`
	Idetabrubr      string            `xml:"ideTabRubr"`
	Vrrubr          float64           `xml:"vrRubr"`
	Ideprocessofgts []Ideprocessofgts `xml:"ideProcessoFGTS"`
}
