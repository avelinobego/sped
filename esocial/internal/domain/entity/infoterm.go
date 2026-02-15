package entity

type Infoterm struct {
	Dtterm       string   `xml:"dtTerm"`
	Mtvdesligtsv []string `xml:"mtvDesligTSV,omitempty"`
}
