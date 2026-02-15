package entity

type Cat struct {
	Dtacid           string         `xml:"dtAcid"`
	Tpacid           int64          `xml:"tpAcid"`
	Hracid           []string       `xml:"hrAcid,omitempty"`
	Hrstrabantesacid []string       `xml:"hrsTrabAntesAcid,omitempty"`
	Tpcat            int64          `xml:"tpCat"`
	Indcatobito      string         `xml:"indCatObito"`
	Dtobito          []string       `xml:"dtObito,omitempty"`
	Indcomunpolicia  string         `xml:"indComunPolicia"`
	Codsitgeradora   int64          `xml:"codSitGeradora"`
	Iniciatcat       int64          `xml:"iniciatCAT"`
	Obscat           []string       `xml:"obsCAT,omitempty"`
	Ultdiatrab       []string       `xml:"ultDiaTrab,omitempty"`
	Houveafast       []string       `xml:"houveAfast,omitempty"`
	Localacidente    Localacidente  `xml:"localAcidente"`
	Parteatingida    Parteatingida  `xml:"parteAtingida"`
	Agentecausador   Agentecausador `xml:"agenteCausador"`
	Atestado         Atestado       `xml:"atestado"`
	Catorigem        []Catorigem    `xml:"catOrigem,omitempty"`
}
