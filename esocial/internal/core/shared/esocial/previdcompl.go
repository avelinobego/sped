package esocial

type Previdcompl struct {
	Tpprev          int64    `xml:"tpPrev"`
	Cnpjentidpc     string   `xml:"cnpjEntidPC"`
	Vlrdedpc        []string `xml:"vlrDedPC,omitempty"`
	Vlrdedpc13      []string `xml:"vlrDedPC13,omitempty"`
	Vlrpatrocfunp   []string `xml:"vlrPatrocFunp,omitempty"`
	Vlrpatrocfunp13 []string `xml:"vlrPatrocFunp13,omitempty"`
}
