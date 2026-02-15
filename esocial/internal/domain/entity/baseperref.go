package entity

type Baseperref struct {
	Perref          string   `xml:"perRef"`
	Codcateg        int64    `xml:"codCateg"`
	Tpvalorproctrab int64    `xml:"tpValorProcTrab"`
	Remfgtsproctrab float64  `xml:"remFGTSProcTrab"`
	Dpsfgtsproctrab []string `xml:"dpsFGTSProcTrab,omitempty"`
	Remfgtssefip    []string `xml:"remFGTSSefip,omitempty"`
	Dpsfgtssefip    []string `xml:"dpsFGTSSefip,omitempty"`
	Remfgtsdecant   []string `xml:"remFGTSDecAnt,omitempty"`
	Dpsfgtsdecant   []string `xml:"dpsFGTSDecAnt,omitempty"`
}
