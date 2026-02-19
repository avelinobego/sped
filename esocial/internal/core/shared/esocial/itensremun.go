package esocial

type Itensremun struct {
	Descfolha  []Descfolha `xml:"descFolha,omitempty"`
	Codrubr    string      `xml:"codRubr"`
	Idetabrubr string      `xml:"ideTabRubr"`
	Qtdrubr    []string    `xml:"qtdRubr,omitempty"`
	Fatorrubr  []string    `xml:"fatorRubr,omitempty"`
	Vrrubr     float64     `xml:"vrRubr"`
	Indapurir  int64       `xml:"indApurIR"`
}
