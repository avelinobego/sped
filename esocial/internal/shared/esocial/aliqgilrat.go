package esocial

type Aliqgilrat struct {
	Aliqrat       []string        `xml:"aliqRat,omitempty"`
	Fap           []string        `xml:"fap,omitempty"`
	Procadmjudrat []Procadmjudrat `xml:"procAdmJudRat,omitempty"`
	Procadmjudfap []Procadmjudfap `xml:"procAdmJudFap,omitempty"`
}
