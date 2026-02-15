package entity

type Idetrabsemvinculo struct {
	Cpftrab   string   `xml:"cpfTrab"`
	Matricula []string `xml:"matricula,omitempty"`
	Codcateg  []string `xml:"codCateg,omitempty"`
}
