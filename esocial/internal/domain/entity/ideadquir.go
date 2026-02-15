package entity

type Ideadquir struct {
	Tpinsc   int64   `xml:"tpInsc"`
	Nrinsc   string  `xml:"nrInsc"`
	Vrcomerc float64 `xml:"vrComerc"`
	Nfs      []Nfs   `xml:"nfs"`
}
