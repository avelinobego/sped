package esocial

type Baseperapur struct {
	Remfgts     float64       `xml:"remFGTS"`
	Dpsfgts     []string      `xml:"dpsFGTS,omitempty"`
	Detrubrsusp []Detrubrsusp `xml:"detRubrSusp"`
	Tpvalor     int64         `xml:"tpValor"`
	Indincid    int64         `xml:"indIncid"`
	Basefgts    float64       `xml:"baseFGTS"`
	Vrfgts      []string      `xml:"vrFGTS,omitempty"`
	Notaft      []string      `xml:"notAFT,omitempty"`
	Natrubr     []string      `xml:"natRubr,omitempty"`
}
