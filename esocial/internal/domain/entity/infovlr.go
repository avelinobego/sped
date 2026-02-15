package entity

type Infovlr struct {
	Compini    string       `xml:"compIni"`
	Compfim    string       `xml:"compFim"`
	Indreperc  int64        `xml:"indReperc"`
	Indensd    []string     `xml:"indenSD,omitempty"`
	Indenabono []string     `xml:"indenAbono,omitempty"`
	Abono      []Abono      `xml:"abono"`
	Ideperiodo []Ideperiodo `xml:"idePeriodo"`
}
