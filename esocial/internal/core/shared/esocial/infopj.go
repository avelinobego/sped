package esocial

type Infopj struct {
	Indcoop              []string     `xml:"indCoop,omitempty"`
	Indconstr            int64        `xml:"indConstr"`
	Indsubstpatr         []string     `xml:"indSubstPatr,omitempty"`
	Percredcontrib       []string     `xml:"percRedContrib,omitempty"`
	Perctransf           []string     `xml:"percTransf,omitempty"`
	Indtribfolhapispasep []string     `xml:"indTribFolhaPisPasep,omitempty"`
	Infoatconc           []Infoatconc `xml:"infoAtConc,omitempty"`
}
