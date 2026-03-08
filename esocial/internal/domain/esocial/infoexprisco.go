package esocial

type Infoexprisco struct {
	Dtinicondicao string    `xml:"dtIniCondicao"`
	Dtfimcondicao []string  `xml:"dtFimCondicao,omitempty"`
	Infoamb       []Infoamb `xml:"infoAmb"`
	Infoativ      Infoativ  `xml:"infoAtiv"`
	Agnoc         []Agnoc   `xml:"agNoc"`
	Respreg       []Respreg `xml:"respReg"`
	Obs           []Obs     `xml:"obs,omitempty"`
}
