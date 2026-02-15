package entity

type Infocs struct {
	Nrrecarqbase  string          `xml:"nrRecArqBase"`
	Indexistinfo  int64           `xml:"indExistInfo"`
	Infocpseg     []Infocpseg     `xml:"infoCPSeg,omitempty"`
	Infocontrib   Infocontrib     `xml:"infoContrib"`
	Ideestab      []Ideestab      `xml:"ideEstab"`
	Infocrcontrib []Infocrcontrib `xml:"infoCRContrib"`
}
