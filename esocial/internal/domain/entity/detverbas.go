package entity

type Detverbas struct {
	Codrubr    string      `xml:"codRubr"`
	Idetabrubr string      `xml:"ideTabRubr"`
	Qtdrubr    []string    `xml:"qtdRubr,omitempty"`
	Fatorrubr  []string    `xml:"fatorRubr,omitempty"`
	Vrrubr     float64     `xml:"vrRubr"`
	Indapurir  []string    `xml:"indApurIR,omitempty"`
	Descfolha  []Descfolha `xml:"descFolha,omitempty"`
}
