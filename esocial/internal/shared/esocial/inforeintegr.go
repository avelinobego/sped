package esocial

type Inforeintegr struct {
	Tpreint       int64    `xml:"tpReint"`
	Nrprocjud     []string `xml:"nrProcJud,omitempty"`
	Nrleianistia  []string `xml:"nrLeiAnistia,omitempty"`
	Dtefetretorno string   `xml:"dtEfetRetorno"`
	Dtefeito      string   `xml:"dtEfeito"`
}
