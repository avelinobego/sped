package esocial

type Aso struct {
	Dtaso  string   `xml:"dtAso"`
	Resaso []string `xml:"resAso,omitempty"`
	Exame  []Exame  `xml:"exame"`
	Medico Medico   `xml:"medico"`
}
