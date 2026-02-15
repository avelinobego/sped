package entity

type Infotrabfgts struct {
	Tpregtrab        []string         `xml:"tpRegTrab,omitempty"`
	Remunsuc         []string         `xml:"remunSuc,omitempty"`
	Dtdeslig         []string         `xml:"dtDeslig,omitempty"`
	Mtvdeslig        []string         `xml:"mtvDeslig,omitempty"`
	Dtterm           []string         `xml:"dtTerm,omitempty"`
	Mtvdesligtsv     []string         `xml:"mtvDesligTSV,omitempty"`
	Sucessaovinc     []Sucessaovinc   `xml:"sucessaoVinc,omitempty"`
	Infobasefgts     []Infobasefgts   `xml:"infoBaseFGTS,omitempty"`
	Proccs           []Proccs         `xml:"procCS,omitempty"`
	Econsignado      []Econsignado    `xml:"eConsignado"`
	Matricula        []string         `xml:"matricula,omitempty"`
	Codcateg         []string         `xml:"codCateg,omitempty"`
	Categorig        []string         `xml:"categOrig,omitempty"`
	Infofgtsproctrab Infofgtsproctrab `xml:"infoFGTSProcTrab"`
}
