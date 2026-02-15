package entity

type Idevinculo struct {
	Codcateg  []string `xml:"codCateg,omitempty"`
	Cpftrab   string   `xml:"cpfTrab"`
	Matricula string   `xml:"matricula"`
}
