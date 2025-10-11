package retorno

type Processamento struct {
	Codigo       int          `xml:"cdResposta"`
	DescResposta string       `xml:"descResposta"`
	VersaoApp    string       `xml:"versaoAppRecepcao"`
	Ocorrencias  []Ocorrencia `xml:"ocorrencias,omitempty"`
}
