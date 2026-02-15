package entity

type Infodep struct {
	Cpfdep   string   `xml:"cpfDep"`
	Dtnascto []string `xml:"dtNascto,omitempty"`
	Nome     []string `xml:"nome,omitempty"`
	Depirrf  []string `xml:"depIRRF,omitempty"`
	Tpdep    []string `xml:"tpDep,omitempty"`
	Descrdep []string `xml:"descrDep,omitempty"`
}
