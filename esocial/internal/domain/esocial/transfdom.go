package esocial

type Transfdom struct {
	Cpfsubstituido string   `xml:"cpfSubstituido"`
	Matricant      []string `xml:"matricAnt,omitempty"`
	Dttransf       string   `xml:"dtTransf"`
}
