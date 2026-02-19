package esocial

type Basesremun struct {
	Indincid  int64       `xml:"indIncid"`
	Codcateg  int64       `xml:"codCateg"`
	Basescp   Basescp     `xml:"basesCp"`
	Basescp13 []Basescp13 `xml:"basesCp13,omitempty"`
}
