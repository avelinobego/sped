package entity

type Infocrirrf struct {
	Vrcr13      []string      `xml:"vrCR13,omitempty"`
	Infoir      []Infoir      `xml:"infoIR,omitempty"`
	Inforra     []Inforra     `xml:"infoRRA,omitempty"`
	Deddepen    []Deddepen    `xml:"dedDepen"`
	Penalim     []Penalim     `xml:"penAlim"`
	Infoprocret []Infoprocret `xml:"infoProcRet"`
	Tpcr        int64         `xml:"tpCR"`
	Vrcr        float64       `xml:"vrCR"`
}
