package entity

type Infoestagiario struct {
	Natestagio        string              `xml:"natEstagio"`
	Nivestagio        []string            `xml:"nivEstagio,omitempty"`
	Areaatuacao       []string            `xml:"areaAtuacao,omitempty"`
	Nrapol            []string            `xml:"nrApol,omitempty"`
	Dtprevterm        string              `xml:"dtPrevTerm"`
	Instensino        Instensino          `xml:"instEnsino"`
	Ageintegracao     []Ageintegracao     `xml:"ageIntegracao,omitempty"`
	Supervisorestagio []Supervisorestagio `xml:"supervisorEstagio,omitempty"`
}
