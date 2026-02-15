package entity

type Instensino struct {
	Cnpjinstensino []string `xml:"cnpjInstEnsino,omitempty"`
	Nmrazao        []string `xml:"nmRazao,omitempty"`
	Dsclograd      []string `xml:"dscLograd,omitempty"`
	Nrlograd       []string `xml:"nrLograd,omitempty"`
	Bairro         []string `xml:"bairro,omitempty"`
	Cep            []string `xml:"cep,omitempty"`
	Codmunic       []string `xml:"codMunic,omitempty"`
	Uf             []string `xml:"uf,omitempty"`
}
