package entity

type Dmdev struct {
	Infocomplcont []Infocomplcont `xml:"infoComplCont,omitempty"`
	Nrbeneficio   string          `xml:"nrBeneficio"`
	Indrra        []string        `xml:"indRRA,omitempty"`
	Notaft        []string        `xml:"notAFT,omitempty"`
	Infoperapur   []Infoperapur   `xml:"infoPerApur,omitempty"`
	Infoperant    []Infoperant    `xml:"infoPerAnt,omitempty"`
	Idedmdev      string          `xml:"ideDmDev"`
	Tppgto        int64           `xml:"tpPgto"`
	Dtpgto        string          `xml:"dtPgto"`
	Codcateg      int64           `xml:"codCateg"`
	Infoir        []Infoir        `xml:"infoIR"`
	Totapurmen    []Totapurmen    `xml:"totApurMen"`
	Totapurdia    []Totapurdia    `xml:"totApurDia"`
	Inforra       []Inforra       `xml:"infoRRA,omitempty"`
	Infopgtoext   []Infopgtoext   `xml:"infoPgtoExt,omitempty"`
}
