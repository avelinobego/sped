package esocial

type Infoir struct {
	Vrrendtrib        []string         `xml:"vrRendTrib,omitempty"`
	Vrrendtrib13      []string         `xml:"vrRendTrib13,omitempty"`
	Vrrendmolegrave   []string         `xml:"vrRendMoleGrave,omitempty"`
	Vrrendmolegrave13 []string         `xml:"vrRendMoleGrave13,omitempty"`
	Vrrendisen65      []string         `xml:"vrRendIsen65,omitempty"`
	Vrrendisen65dec   []string         `xml:"vrRendIsen65Dec,omitempty"`
	Vrjurosmora       []string         `xml:"vrJurosMora,omitempty"`
	Vrjurosmora13     []string         `xml:"vrJurosMora13,omitempty"`
	Vrrendisenntrib   []string         `xml:"vrRendIsenNTrib,omitempty"`
	Descisenntrib     []string         `xml:"descIsenNTrib,omitempty"`
	Vrprevoficial     []string         `xml:"vrPrevOficial,omitempty"`
	Vrprevoficial13   []string         `xml:"vrPrevOficial13,omitempty"`
	Rendisen0561      []Rendisen0561   `xml:"rendIsen0561,omitempty"`
	Tpinfoir          int64            `xml:"tpInfoIR"`
	Valor             float64          `xml:"valor"`
	Descrendimento    []string         `xml:"descRendimento,omitempty"`
	Infoprocjudrub    []Infoprocjudrub `xml:"infoProcJudRub"`
}
