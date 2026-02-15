package entity

type Infovalores struct {
	Indapuracao  int64     `xml:"indApuracao"`
	Vlrnretido   []string  `xml:"vlrNRetido,omitempty"`
	Vlrdepjud    []string  `xml:"vlrDepJud,omitempty"`
	Vlrcmpanocal []string  `xml:"vlrCmpAnoCal,omitempty"`
	Vlrcmpanoant []string  `xml:"vlrCmpAnoAnt,omitempty"`
	Vlrrendsusp  []string  `xml:"vlrRendSusp,omitempty"`
	Dedsusp      []Dedsusp `xml:"dedSusp"`
}
