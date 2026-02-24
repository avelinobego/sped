package esocial

type Infointerm struct {
	Dia     int64    `xml:"dia"`
	Hrstrab []string `xml:"hrsTrab,omitempty"`
}
