package entity

type Inforra struct {
	Tpprocrra   int64         `xml:"tpProcRRA"`
	Nrprocrra   []string      `xml:"nrProcRRA,omitempty"`
	Descrra     string        `xml:"descRRA"`
	Qtdmesesrra float64       `xml:"qtdMesesRRA"`
	Despprocjud []Despprocjud `xml:"despProcJud,omitempty"`
	Ideadv      []Ideadv      `xml:"ideAdv"`
}
