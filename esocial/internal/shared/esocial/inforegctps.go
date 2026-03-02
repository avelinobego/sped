package esocial

type Inforegctps struct {
	Cbocargo   string   `xml:"CBOCargo"`
	Vrsalfx    float64  `xml:"vrSalFx"`
	Undsalfixo int64    `xml:"undSalFixo"`
	Tpcontr    int64    `xml:"tpContr"`
	Dtterm     []string `xml:"dtTerm,omitempty"`
}
