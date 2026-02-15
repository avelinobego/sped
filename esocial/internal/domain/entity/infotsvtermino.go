package entity

type Infotsvtermino struct {
	Dtterm        string          `xml:"dtTerm"`
	Mtvdesligtsv  []string        `xml:"mtvDesligTSV,omitempty"`
	Pensalim      []string        `xml:"pensAlim,omitempty"`
	Percaliment   []string        `xml:"percAliment,omitempty"`
	Vralim        []string        `xml:"vrAlim,omitempty"`
	Nrproctrab    []string        `xml:"nrProcTrab,omitempty"`
	Mudancacpf    []Mudancacpf    `xml:"mudancaCPF,omitempty"`
	Verbasresc    []Verbasresc    `xml:"verbasResc,omitempty"`
	Remunaposterm []Remunaposterm `xml:"remunAposTerm,omitempty"`
}
