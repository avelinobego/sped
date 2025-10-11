package retorno

type Ocorrencia struct {
	Tipo        TipoAdvertencia `xml:"tipo"`
	Codigo      string          `xml:"codigo"`
	Descricao   string          `xml:"descricao"`
	Localizacao string          `xml:"localizacao,omitempty"`
}
