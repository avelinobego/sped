package entity

type Detreembtit struct {
	Tpinsc      int64    `xml:"tpInsc"`
	Nrinsc      string   `xml:"nrInsc"`
	Vlrreemb    []string `xml:"vlrReemb,omitempty"`
	Vlrreembant []string `xml:"vlrReembAnt,omitempty"`
}
