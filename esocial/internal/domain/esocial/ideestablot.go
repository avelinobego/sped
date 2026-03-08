package esocial

type Ideestablot struct {
	Qtddiasav      []string         `xml:"qtdDiasAv,omitempty"`
	Remunperapur   []Remunperapur   `xml:"remunPerApur"`
	Remunperant    []Remunperant    `xml:"remunPerAnt"`
	Infoagnocivo   []Infoagnocivo   `xml:"infoAgNocivo,omitempty"`
	Detverbas      []Detverbas      `xml:"detVerbas"`
	Infosimples    []Infosimples    `xml:"infoSimples,omitempty"`
	Tpinsc         int64            `xml:"tpInsc"`
	Nrinsc         string           `xml:"nrInsc"`
	Codlotacao     string           `xml:"codLotacao"`
	Infocategincid []Infocategincid `xml:"infoCategIncid"`
}
