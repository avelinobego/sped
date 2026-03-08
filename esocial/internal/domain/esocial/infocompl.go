package esocial

type Infocompl struct {
	Codcbo        []string        `xml:"codCBO,omitempty"`
	Natatividade  []string        `xml:"natAtividade,omitempty"`
	Remuneracao   []Remuneracao   `xml:"remuneracao"`
	Infovinc      []Infovinc      `xml:"infoVinc,omitempty"`
	Infoterm      []Infoterm      `xml:"infoTerm,omitempty"`
	Sucessaovinc  []Sucessaovinc  `xml:"sucessaoVinc,omitempty"`
	Infointerm    []Infointerm    `xml:"infoInterm"`
	Infocomplcont []Infocomplcont `xml:"infoComplCont"`
}
