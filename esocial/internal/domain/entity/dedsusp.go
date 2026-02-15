package entity

type Dedsusp struct {
	Indtpdeducao  int64      `xml:"indTpDeducao"`
	Vlrdedsusp    []string   `xml:"vlrDedSusp,omitempty"`
	Cnpjentidpc   []string   `xml:"cnpjEntidPC,omitempty"`
	Vlrpatrocfunp []string   `xml:"vlrPatrocFunp,omitempty"`
	Benefpen      []Benefpen `xml:"benefPen"`
}
