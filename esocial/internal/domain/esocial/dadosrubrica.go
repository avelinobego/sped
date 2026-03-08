package esocial

type Dadosrubrica struct {
	Dscrubr             string                `xml:"dscRubr"`
	Natrubr             int64                 `xml:"natRubr"`
	Tprubr              int64                 `xml:"tpRubr"`
	Codinccp            string                `xml:"codIncCP"`
	Codincirrf          int64                 `xml:"codIncIRRF"`
	Codincfgts          string                `xml:"codIncFGTS"`
	Codinccprp          []string              `xml:"codIncCPRP,omitempty"`
	Codincpispasep      []string              `xml:"codIncPisPasep,omitempty"`
	Tetoremun           []string              `xml:"tetoRemun,omitempty"`
	Observacao          []string              `xml:"observacao,omitempty"`
	Ideprocessocp       []Ideprocessocp       `xml:"ideProcessoCP"`
	Ideprocessoirrf     []Ideprocessoirrf     `xml:"ideProcessoIRRF"`
	Ideprocessofgts     []Ideprocessofgts     `xml:"ideProcessoFGTS"`
	Ideprocessopispasep []Ideprocessopispasep `xml:"ideProcessoPisPasep"`
}
