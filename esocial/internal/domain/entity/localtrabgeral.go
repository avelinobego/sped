package entity

type Localtrabgeral struct {
	Tpinsc   int64    `xml:"tpInsc"`
	Nrinsc   string   `xml:"nrInsc"`
	Desccomp []string `xml:"descComp,omitempty"`
}
