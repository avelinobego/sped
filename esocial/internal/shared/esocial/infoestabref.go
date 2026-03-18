package esocial

type Infoestabref struct {
	Aliqrat      int64    `xml:"aliqRat"`
	Fap          []string `xml:"fap,omitempty"`
	Aliqratajust []string `xml:"aliqRatAjust,omitempty"`
}
