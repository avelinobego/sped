package esocial

type Inforetif struct {
	Origretif int64    `xml:"origRetif"`
	Tpproc    []string `xml:"tpProc,omitempty"`
	Nrproc    []string `xml:"nrProc,omitempty"`
}
