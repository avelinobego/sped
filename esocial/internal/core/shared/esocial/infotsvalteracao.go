package esocial

type Infotsvalteracao struct {
	Dtalteracao        string               `xml:"dtAlteracao"`
	Natatividade       []string             `xml:"natAtividade,omitempty"`
	Infocomplementares []Infocomplementares `xml:"infoComplementares,omitempty"`
}
