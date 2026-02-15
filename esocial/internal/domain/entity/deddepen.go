package entity

type Deddepen struct {
	Vlrdeducao float64 `xml:"vlrDeducao"`
	Tprend     int64   `xml:"tpRend"`
	Cpfdep     string  `xml:"cpfDep"`
	Vlrdeddep  float64 `xml:"vlrDedDep"`
}
