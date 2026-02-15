package entity

type Infocomplcont struct {
	Codcbo       string   `xml:"codCBO"`
	Natatividade []string `xml:"natAtividade,omitempty"`
	Qtddiastrab  []string `xml:"qtdDiasTrab,omitempty"`
}
