package esocial

type Infofgts struct {
	Vrbcfgtsproctrab float64    `xml:"vrBcFGTSProcTrab"`
	Vrbcfgtssefip    []string   `xml:"vrBcFGTSSefip,omitempty"`
	Vrbcfgtsdecant   []string   `xml:"vrBcFGTSDecAnt,omitempty"`
	Dtvenc           []string   `xml:"dtVenc,omitempty"`
	Classtrib        []string   `xml:"classTrib,omitempty"`
	Nrrecarqbase     string     `xml:"nrRecArqBase"`
	Indexistinfo     int64      `xml:"indExistInfo"`
	Ideestab         []Ideestab `xml:"ideEstab"`
}
