package entity

type Infoperref struct {
	Perref        string          `xml:"perRef"`
	Ideadc        []Ideadc        `xml:"ideADC"`
	Detinfoperref []Detinfoperref `xml:"detInfoPerRef"`
}
