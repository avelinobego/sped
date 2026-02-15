package entity

type Remuneracao struct {
	Dtremun    string   `xml:"dtRemun"`
	Vrsalfx    float64  `xml:"vrSalFx"`
	Undsalfixo int64    `xml:"undSalFixo"`
	Dscsalvar  []string `xml:"dscSalVar,omitempty"`
}
