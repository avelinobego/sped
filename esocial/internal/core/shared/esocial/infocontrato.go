package esocial

type Infocontrato struct {
	Dtingrcargo    []string         `xml:"dtIngrCargo,omitempty"`
	Nmcargo        []string         `xml:"nmCargo,omitempty"`
	Cbocargo       []string         `xml:"CBOCargo,omitempty"`
	Nmfuncao       []string         `xml:"nmFuncao,omitempty"`
	Cbofuncao      []string         `xml:"CBOFuncao,omitempty"`
	Acumcargo      []string         `xml:"acumCargo,omitempty"`
	Codcateg       int64            `xml:"codCateg"`
	Remuneracao    []Remuneracao    `xml:"remuneracao,omitempty"`
	Duracao        []Duracao        `xml:"duracao,omitempty"`
	Localtrabalho  Localtrabalho    `xml:"localTrabalho"`
	Horcontratual  []Horcontratual  `xml:"horContratual,omitempty"`
	Alvarajudicial []Alvarajudicial `xml:"alvaraJudicial,omitempty"`
	Observacoes    []Observacoes    `xml:"observacoes"`
	Treicap        []Treicap        `xml:"treiCap"`
}
