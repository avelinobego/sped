package esocial

type Idebenef struct {
	Infopgto      []Infopgto      `xml:"infoPgto"`
	Infoircomplem []Infoircomplem `xml:"infoIRComplem"`
	Cpfbenef      string          `xml:"cpfBenef"`
}
