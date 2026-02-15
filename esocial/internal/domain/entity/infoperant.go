package entity

type Infoperant struct {
	Remunorgsuc string       `xml:"remunOrgSuc"`
	Ideperiodo  []Ideperiodo `xml:"idePeriodo"`
	Ideadc      []Ideadc     `xml:"ideADC"`
}
