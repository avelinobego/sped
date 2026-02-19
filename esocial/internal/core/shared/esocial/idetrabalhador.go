package esocial

type Idetrabalhador struct {
	Infomv        []Infomv        `xml:"infoMV,omitempty"`
	Infointerm    []Infointerm    `xml:"infoInterm"`
	Infocomplem   []Infocomplem   `xml:"infoComplem,omitempty"`
	Infocompl     []Infocompl     `xml:"infoCompl,omitempty"`
	Procjudtrab   []Procjudtrab   `xml:"procJudTrab"`
	Cpfbenef      string          `xml:"cpfBenef"`
	Dmdev         []Dmdev         `xml:"dmDev"`
	Totinfoir     []Totinfoir     `xml:"totInfoIR,omitempty"`
	Infoircomplem []Infoircomplem `xml:"infoIRComplem"`
	Cpftrab       string          `xml:"cpfTrab"`
}
