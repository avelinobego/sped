package entity

type Infobaseperante struct {
	Perref      string        `xml:"perRef"`
	Tpacconv    string        `xml:"tpAcConv"`
	Baseperante []Baseperante `xml:"basePerAntE"`
}
