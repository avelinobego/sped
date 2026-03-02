package esocial

type Dadosbeneficio struct {
	Inddecjud    []string       `xml:"indDecJud,omitempty"`
	Infohomolog  []Infohomolog  `xml:"infoHomolog,omitempty"`
	Tpbeneficio  string         `xml:"tpBeneficio"`
	Tpplanrp     int64          `xml:"tpPlanRP"`
	Dsc          []string       `xml:"dsc,omitempty"`
	Indsuspensao string         `xml:"indSuspensao"`
	Infopenmorte []Infopenmorte `xml:"infoPenMorte,omitempty"`
	Suspensao    []Suspensao    `xml:"suspensao,omitempty"`
}
