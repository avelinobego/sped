package entity

type Idedep struct {
	Cpfdep   string   `xml:"cpfDep"`
	Depirrf  []string `xml:"depIRRF,omitempty"`
	Dtnascto []string `xml:"dtNascto,omitempty"`
	Nome     []string `xml:"nome,omitempty"`
	Tpdep    []string `xml:"tpDep,omitempty"`
	Descrdep []string `xml:"descrDep,omitempty"`
}
