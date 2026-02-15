package entity

type Localtempdom struct {
	Tplograd    []string `xml:"tpLograd,omitempty"`
	Dsclograd   string   `xml:"dscLograd"`
	Nrlograd    string   `xml:"nrLograd"`
	Complemento []string `xml:"complemento,omitempty"`
	Bairro      []string `xml:"bairro,omitempty"`
	Cep         string   `xml:"cep"`
	Codmunic    int64    `xml:"codMunic"`
	Uf          string   `xml:"uf"`
}
