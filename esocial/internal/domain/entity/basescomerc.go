package entity

type Basescomerc struct {
	Indcomerc   int64    `xml:"indComerc"`
	Vrbccompr   float64  `xml:"vrBcComPR"`
	Vrcpsusp    []string `xml:"vrCPSusp,omitempty"`
	Vrratsusp   []string `xml:"vrRatSusp,omitempty"`
	Vrsenarsusp []string `xml:"vrSenarSusp,omitempty"`
}
