package esocial

type Infotsvinicio struct {
	Cadini             string               `xml:"cadIni"`
	Matricula          []string             `xml:"matricula,omitempty"`
	Codcateg           int64                `xml:"codCateg"`
	Dtinicio           string               `xml:"dtInicio"`
	Nrproctrab         []string             `xml:"nrProcTrab,omitempty"`
	Natatividade       []string             `xml:"natAtividade,omitempty"`
	Infocomplementares []Infocomplementares `xml:"infoComplementares,omitempty"`
	Mudancacpf         []Mudancacpf         `xml:"mudancaCPF,omitempty"`
	Afastamento        []Afastamento        `xml:"afastamento,omitempty"`
	Termino            []Termino            `xml:"termino,omitempty"`
}
