package tipos

type TipoAmbiente int

const (
	Producao TipoAmbiente = iota + 1
	PreProducaoDadosreais
	PreProducaoDadosFicticios
	Homologacao TipoAmbiente = iota + 3
	Validacao
	Testes
	Desenvolvimento
)
