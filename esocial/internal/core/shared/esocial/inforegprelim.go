package esocial

type Inforegprelim struct {
	Cpftrab      string        `xml:"cpfTrab"`
	Dtnascto     string        `xml:"dtNascto"`
	Dtadm        string        `xml:"dtAdm"`
	Matricula    string        `xml:"matricula"`
	Codcateg     int64         `xml:"codCateg"`
	Natatividade []string      `xml:"natAtividade,omitempty"`
	Inforegctps  []Inforegctps `xml:"infoRegCTPS,omitempty"`
}
