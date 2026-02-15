package entity

type Ideadv struct {
	Tpinsc int64    `xml:"tpInsc"`
	Nrinsc string   `xml:"nrInsc"`
	Vlradv []string `xml:"vlrAdv,omitempty"`
}
