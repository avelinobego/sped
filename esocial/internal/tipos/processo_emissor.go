package tipos

// Valores válidos:
// 1 - Aplicativo do empregador
// 2 - Aplicativo governamental - Simplificado PessoaFísica
// 3 - Aplicativo governamental - Web Geral
// 4 - Aplicativo governamental - Simplificado PessoaJurídica
// 8 - Aplicativo governamental para envio de eventospelo Judiciário
// 9 - Aplicativo governamental - Integração com a JuntaComercial
// 22 - Aplicativo governamental para dispositivosmóveis - Empregador Domésti
type ProcessoEmissor int

const (
	AplicacaoEmpregador           ProcessoEmissor = iota + 1
	AplicacaoSimplificadoPF                       = 2
	AplicacaoWebGeral                             = 3
	AplicacaoSimplificadoPJ                       = 4
	AplicacaoJudiciario                           = 8
	AplicacaoJuntaComercial                       = 9
	AplicacaoEmpregadorDomiciliar                 = 22
)
