package entity

type Horcontratual struct {
	Qtdhrssem  []string `xml:"qtdHrsSem,omitempty"`
	Tpjornada  int64    `xml:"tpJornada"`
	Tmpparc    int64    `xml:"tmpParc"`
	Hornoturno []string `xml:"horNoturno,omitempty"`
	Dscjorn    string   `xml:"dscJorn"`
}
