package entity

type Infocadastro struct {
	Classtrib            string                 `xml:"classTrib"`
	Indcoop              []string               `xml:"indCoop,omitempty"`
	Indconstr            []string               `xml:"indConstr,omitempty"`
	Inddesfolha          int64                  `xml:"indDesFolha"`
	Indopccp             []string               `xml:"indOpcCP,omitempty"`
	Indporte             []string               `xml:"indPorte,omitempty"`
	Indoptregeletron     int64                  `xml:"indOptRegEletron"`
	Cnpjefr              []string               `xml:"cnpjEFR,omitempty"`
	Dttrans11096         []string               `xml:"dtTrans11096,omitempty"`
	Indtribfolhapispasep []string               `xml:"indTribFolhaPisPasep,omitempty"`
	Indpertirrf          []string               `xml:"indPertIRRF,omitempty"`
	Dadosisencao         []Dadosisencao         `xml:"dadosIsencao,omitempty"`
	Infoorginternacional []Infoorginternacional `xml:"infoOrgInternacional,omitempty"`
}
