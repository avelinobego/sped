package entity

type Dependente struct {
	Depsf     string   `xml:"depSF"`
	Inctrab   string   `xml:"incTrab"`
	Tpdep     []string `xml:"tpDep,omitempty"`
	Nmdep     string   `xml:"nmDep"`
	Dtnascto  string   `xml:"dtNascto"`
	Cpfdep    []string `xml:"cpfDep,omitempty"`
	Sexodep   string   `xml:"sexoDep"`
	Depirrf   string   `xml:"depIRRF"`
	Incfismen string   `xml:"incFisMen"`
	Descrdep  []string `xml:"descrDep,omitempty"`
}
