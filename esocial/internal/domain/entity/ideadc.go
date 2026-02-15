package entity

type Ideadc struct {
	Ideperiodo []Ideperiodo `xml:"idePeriodo"`
	Dtacconv   []string     `xml:"dtAcConv,omitempty"`
	Tpacconv   string       `xml:"tpAcConv"`
	Dsc        string       `xml:"dsc"`
	Remunsuc   []string     `xml:"remunSuc,omitempty"`
}
