package entity

type Infoircr struct {
	Tpcr        string        `xml:"tpCR"`
	Deddepen    []Deddepen    `xml:"dedDepen"`
	Penalim     []Penalim     `xml:"penAlim"`
	Previdcompl []Previdcompl `xml:"previdCompl"`
	Infoprocret []Infoprocret `xml:"infoProcRet"`
}
