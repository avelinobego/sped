package esocial

type Penalim struct {
	Vlrpensao     float64 `xml:"vlrPensao"`
	Tprend        int64   `xml:"tpRend"`
	Cpfdep        string  `xml:"cpfDep"`
	Vlrdedpenalim float64 `xml:"vlrDedPenAlim"`
}
