package tipos

type IdeEvento struct {
	TipoAmbiente TipoAmbiente    `xml:"tpAmb"`
	ProcEmi      ProcessoEmissor `xml:"procEmi"`
	Versao       string          `xml:"verProc"`
}
