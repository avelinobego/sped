package entity

type Infofech struct {
	Evtremun        string   `xml:"evtRemun"`
	Evtpgtos        string   `xml:"evtPgtos"`
	Evtcomprod      string   `xml:"evtComProd"`
	Evtcontratavnp  string   `xml:"evtContratAvNP"`
	Evtinfocomplper string   `xml:"evtInfoComplPer"`
	Indexcapur1250  []string `xml:"indExcApur1250,omitempty"`
	Transdctfweb    []string `xml:"transDCTFWeb,omitempty"`
	Naovalid        []string `xml:"naoValid,omitempty"`
}
