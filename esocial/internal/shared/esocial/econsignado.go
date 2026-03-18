package esocial

type Econsignado struct {
	Instfinanc    string  `xml:"instFinanc"`
	Nrcontrato    string  `xml:"nrContrato"`
	Vreconsignado float64 `xml:"vreConsignado"`
}
