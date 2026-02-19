package esocial

type Plansaude struct {
	Cnpjoper    string       `xml:"cnpjOper"`
	Regans      []string     `xml:"regANS,omitempty"`
	Vlrsaudetit float64      `xml:"vlrSaudeTit"`
	Infodepsau  []Infodepsau `xml:"infoDepSau"`
}
