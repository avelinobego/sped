package esocial

type Dadosisencao struct {
	Ideminlei    string   `xml:"ideMinLei"`
	Nrcertif     string   `xml:"nrCertif"`
	Dtemiscertif string   `xml:"dtEmisCertif"`
	Dtvenccertif string   `xml:"dtVencCertif"`
	Nrprotrenov  []string `xml:"nrProtRenov,omitempty"`
	Dtprotrenov  []string `xml:"dtProtRenov,omitempty"`
	Dtdou        []string `xml:"dtDou,omitempty"`
	Pagdou       []string `xml:"pagDou,omitempty"`
}
