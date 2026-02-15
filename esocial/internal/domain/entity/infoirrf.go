package entity

type Infoirrf struct {
	Nrrecarqbase string      `xml:"nrRecArqBase"`
	Indexistinfo int64       `xml:"indExistInfo"`
	Infocrmen    []Infocrmen `xml:"infoCRMen"`
	Infocrdia    []Infocrdia `xml:"infoCRDia"`
}
