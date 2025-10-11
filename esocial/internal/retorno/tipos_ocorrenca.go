package retorno

type TipoAdvertencia int

const (
	Erro TipoAdvertencia = iota + 1
	Advertencia
	HistoricoValidacao
)
