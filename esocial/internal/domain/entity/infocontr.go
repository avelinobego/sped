package entity

type Infocontr struct {
	Tpcontr      int64          `xml:"tpContr"`
	Indcontr     string         `xml:"indContr"`
	Dtadmorig    []string       `xml:"dtAdmOrig,omitempty"`
	Indreint     []string       `xml:"indReint,omitempty"`
	Indcateg     string         `xml:"indCateg"`
	Indnatativ   string         `xml:"indNatAtiv"`
	Indmotdeslig string         `xml:"indMotDeslig"`
	Matricula    []string       `xml:"matricula,omitempty"`
	Codcateg     []string       `xml:"codCateg,omitempty"`
	Dtinicio     []string       `xml:"dtInicio,omitempty"`
	Infocompl    []Infocompl    `xml:"infoCompl,omitempty"`
	Mudcategativ []Mudcategativ `xml:"mudCategAtiv"`
	Uniccontr    []Uniccontr    `xml:"unicContr"`
	Ideestab     Ideestab       `xml:"ideEstab"`
}
