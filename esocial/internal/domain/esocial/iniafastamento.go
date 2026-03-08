package esocial

type Iniafastamento struct {
	Dtiniafast     string         `xml:"dtIniAfast"`
	Codmotafast    string         `xml:"codMotAfast"`
	Infomesmomtv   []string       `xml:"infoMesmoMtv,omitempty"`
	Tpacidtransito []string       `xml:"tpAcidTransito,omitempty"`
	Observacao     []string       `xml:"observacao,omitempty"`
	Peraquis       []Peraquis     `xml:"perAquis,omitempty"`
	Infocessao     []Infocessao   `xml:"infoCessao,omitempty"`
	Infomandsind   []Infomandsind `xml:"infoMandSind,omitempty"`
	Infomandelet   []Infomandelet `xml:"infoMandElet,omitempty"`
}
