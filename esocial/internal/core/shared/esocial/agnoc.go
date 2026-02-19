package esocial

type Agnoc struct {
	Codagnoc   string   `xml:"codAgNoc"`
	Dscagnoc   []string `xml:"dscAgNoc,omitempty"`
	Tpaval     []string `xml:"tpAval,omitempty"`
	Intconc    []string `xml:"intConc,omitempty"`
	Limtol     []string `xml:"limTol,omitempty"`
	Unmed      []string `xml:"unMed,omitempty"`
	Tecmedicao []string `xml:"tecMedicao,omitempty"`
	Nrprocjud  []string `xml:"nrProcJud,omitempty"`
	Epcepi     []Epcepi `xml:"epcEpi,omitempty"`
}
