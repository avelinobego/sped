package esocial

type Atestado struct {
	Dtatendimento string   `xml:"dtAtendimento"`
	Hratendimento string   `xml:"hrAtendimento"`
	Indinternacao string   `xml:"indInternacao"`
	Durtrat       int64    `xml:"durTrat"`
	Indafast      string   `xml:"indAfast"`
	Dsclesao      int64    `xml:"dscLesao"`
	Dsccomplesao  []string `xml:"dscCompLesao,omitempty"`
	Diagprovavel  []string `xml:"diagProvavel,omitempty"`
	Codcid        string   `xml:"codCID"`
	Observacao    []string `xml:"observacao,omitempty"`
	Emitente      Emitente `xml:"emitente"`
}
