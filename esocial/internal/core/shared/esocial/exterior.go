package esocial

type Exterior struct {
	Paisresid   string   `xml:"paisResid"`
	Dsclograd   string   `xml:"dscLograd"`
	Nrlograd    string   `xml:"nrLograd"`
	Complemento []string `xml:"complemento,omitempty"`
	Bairro      []string `xml:"bairro,omitempty"`
	Nmcid       string   `xml:"nmCid"`
	Codpostal   []string `xml:"codPostal,omitempty"`
}
