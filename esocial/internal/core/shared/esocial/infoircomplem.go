package esocial

type Infoircomplem struct {
	Infodep      []Infodep      `xml:"infoDep"`
	Dtlaudo      []string       `xml:"dtLaudo,omitempty"`
	Perant       []Perant       `xml:"perAnt,omitempty"`
	Idedep       []Idedep       `xml:"ideDep"`
	Infoircr     []Infoircr     `xml:"infoIRCR"`
	Plansaude    []Plansaude    `xml:"planSaude"`
	Inforeembmed []Inforeembmed `xml:"infoReembMed"`
}
