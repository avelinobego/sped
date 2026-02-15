package entity

type Remunperant struct {
	Indsimples   []string       `xml:"indSimples,omitempty"`
	Infoagnocivo []Infoagnocivo `xml:"infoAgNocivo,omitempty"`
	Matricula    []string       `xml:"matricula,omitempty"`
	Itensremun   []Itensremun   `xml:"itensRemun"`
}
