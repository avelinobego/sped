package esocial

type Infobeninicio struct {
	Cadini         string           `xml:"cadIni"`
	Indsitbenef    []string         `xml:"indSitBenef,omitempty"`
	Nrbeneficio    string           `xml:"nrBeneficio"`
	Dtinibeneficio string           `xml:"dtIniBeneficio"`
	Dtpublic       []string         `xml:"dtPublic,omitempty"`
	Dadosbeneficio Dadosbeneficio   `xml:"dadosBeneficio"`
	Sucessaobenef  []Sucessaobenef  `xml:"sucessaoBenef,omitempty"`
	Mudancacpf     []Mudancacpf     `xml:"mudancaCPF,omitempty"`
	Infobentermino []Infobentermino `xml:"infoBenTermino,omitempty"`
}
