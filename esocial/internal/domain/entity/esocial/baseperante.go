package esocial

type Baseperante struct {
	Remfgtse    float64       `xml:"remFGTSE"`
	Dpsfgtse    []string      `xml:"dpsFGTSE,omitempty"`
	Detrubrsusp []Detrubrsusp `xml:"detRubrSusp"`
	Tpvalore    int64         `xml:"tpValorE"`
	Indincide   int64         `xml:"indIncidE"`
	Basefgtse   float64       `xml:"baseFGTSE"`
	Vrfgtse     []string      `xml:"vrFGTSE,omitempty"`
}
