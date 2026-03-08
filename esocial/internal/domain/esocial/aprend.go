package esocial

type Aprend struct {
	Indaprend   int64    `xml:"indAprend"`
	Cnpjentqual []string `xml:"cnpjEntQual,omitempty"`
	Tpinsc      []string `xml:"tpInsc,omitempty"`
	Nrinsc      []string `xml:"nrInsc,omitempty"`
	Cnpjprat    []string `xml:"cnpjPrat,omitempty"`
}
