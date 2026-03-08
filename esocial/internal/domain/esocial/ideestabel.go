package esocial

type Ideestabel struct {
	Nrinscestabrural string     `xml:"nrInscEstabRural"`
	Tpcomerc         []Tpcomerc `xml:"tpComerc"`
}
