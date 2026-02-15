package entity

type Infofgtsproctrab struct {
	Totalfgts float64    `xml:"totalFGTS"`
	Ideestab  []Ideestab `xml:"ideEstab,omitempty"`
}
