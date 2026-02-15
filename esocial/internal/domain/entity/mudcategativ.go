package entity

type Mudcategativ struct {
	Codcateg       int64    `xml:"codCateg"`
	Natatividade   []string `xml:"natAtividade,omitempty"`
	Dtmudcategativ string   `xml:"dtMudCategAtiv"`
}
