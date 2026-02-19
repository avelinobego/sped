package esocial

type Infocrcontrib struct {
	Vrcrsusp []string `xml:"vrCRSusp,omitempty"`
	Tpcr     int64    `xml:"tpCR"`
	Vrcr     float64  `xml:"vrCR"`
}
