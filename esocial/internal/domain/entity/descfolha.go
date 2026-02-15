package entity

type Descfolha struct {
	Tpdesc     int64    `xml:"tpDesc"`
	Instfinanc string   `xml:"instFinanc"`
	Nrdoc      string   `xml:"nrDoc"`
	Observacao []string `xml:"observacao,omitempty"`
}
