package esocial

type Infopgto struct {
	Dtpgto       string        `xml:"dtPgto"`
	Tppgto       int64         `xml:"tpPgto"`
	Perref       string        `xml:"perRef"`
	Idedmdev     string        `xml:"ideDmDev"`
	Vrliq        float64       `xml:"vrLiq"`
	Paisresidext []string      `xml:"paisResidExt,omitempty"`
	Infopgtoext  []Infopgtoext `xml:"infoPgtoExt,omitempty"`
}
