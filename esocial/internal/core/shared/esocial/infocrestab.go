package esocial

type Infocrestab struct {
	Tpcr     int64    `xml:"tpCR"`
	Vrcr     float64  `xml:"vrCR"`
	Vrsuspcr []string `xml:"vrSuspCR,omitempty"`
}
