package esocial

type Infoprocret struct {
	Tpprocret   int64         `xml:"tpProcRet"`
	Nrprocret   string        `xml:"nrProcRet"`
	Codsusp     []string      `xml:"codSusp,omitempty"`
	Infovalores []Infovalores `xml:"infoValores"`
}
