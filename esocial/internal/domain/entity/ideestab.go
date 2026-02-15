package entity

type Ideestab struct {
	Tpinsc   int64    `xml:"tpInsc"`
	Nrinsc   string   `xml:"nrInsc"`
	Inivalid string   `xml:"iniValid"`
	Fimvalid []string `xml:"fimValid,omitempty"`
}
