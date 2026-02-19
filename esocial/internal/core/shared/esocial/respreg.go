package esocial

type Respreg struct {
	Cpfresp string   `xml:"cpfResp"`
	Ideoc   []string `xml:"ideOC,omitempty"`
	Dscoc   []string `xml:"dscOC,omitempty"`
	Nroc    []string `xml:"nrOC,omitempty"`
	Ufoc    []string `xml:"ufOC,omitempty"`
}
