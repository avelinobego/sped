package entity

type Nfs struct {
	Serie       []string `xml:"serie,omitempty"`
	Nrdocto     string   `xml:"nrDocto"`
	Dtemisnf    string   `xml:"dtEmisNF"`
	Vlrbruto    float64  `xml:"vlrBruto"`
	Vrcpdescpr  float64  `xml:"vrCPDescPR"`
	Vrratdescpr float64  `xml:"vrRatDescPR"`
	Vrsenardesc float64  `xml:"vrSenarDesc"`
}
