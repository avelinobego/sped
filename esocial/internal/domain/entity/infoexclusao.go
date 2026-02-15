package entity

type Infoexclusao struct {
	Idetrabalhador []Idetrabalhador `xml:"ideTrabalhador,omitempty"`
	Idefolhapagto  []Idefolhapagto  `xml:"ideFolhaPagto,omitempty"`
	Tpevento       string           `xml:"tpEvento"`
	Nrrecevt       string           `xml:"nrRecEvt"`
	Ideproctrab    Ideproctrab      `xml:"ideProcTrab"`
}
