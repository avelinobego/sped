package esocial

type Infosusp struct {
	Codsusp     int64  `xml:"codSusp"`
	Indsusp     string `xml:"indSusp"`
	Dtdecisao   string `xml:"dtDecisao"`
	Inddeposito string `xml:"indDeposito"`
}
