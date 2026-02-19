package esocial

type Ideprocesso struct {
	Tpproc   int64    `xml:"tpProc"`
	Nrproc   string   `xml:"nrProc"`
	Inivalid string   `xml:"iniValid"`
	Fimvalid []string `xml:"fimValid,omitempty"`
}
