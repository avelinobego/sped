package esocial

type Localacidente struct {
	Tplocal      int64          `xml:"tpLocal"`
	Dsclocal     []string       `xml:"dscLocal,omitempty"`
	Tplograd     []string       `xml:"tpLograd,omitempty"`
	Dsclograd    string         `xml:"dscLograd"`
	Nrlograd     string         `xml:"nrLograd"`
	Complemento  []string       `xml:"complemento,omitempty"`
	Bairro       []string       `xml:"bairro,omitempty"`
	Cep          []string       `xml:"cep,omitempty"`
	Codmunic     []string       `xml:"codMunic,omitempty"`
	Uf           []string       `xml:"uf,omitempty"`
	Pais         []string       `xml:"pais,omitempty"`
	Codpostal    []string       `xml:"codPostal,omitempty"`
	Idelocalacid []Idelocalacid `xml:"ideLocalAcid,omitempty"`
}
