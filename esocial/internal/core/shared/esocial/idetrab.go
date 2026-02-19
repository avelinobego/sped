package esocial

type Idetrab struct {
	Nmtrab        []string        `xml:"nmTrab,omitempty"`
	Dtnascto      []string        `xml:"dtNascto,omitempty"`
	Ideseqtrab    []string        `xml:"ideSeqTrab,omitempty"`
	Infocontr     []Infocontr     `xml:"infoContr"`
	Cpftrab       string          `xml:"cpfTrab"`
	Calctrib      []Calctrib      `xml:"calcTrib"`
	Infocrirrf    []Infocrirrf    `xml:"infoCRIRRF"`
	Infoircomplem []Infoircomplem `xml:"infoIRComplem,omitempty"`
}
