package entity

// Empresas maps to the 'empresas' table.
type Empresas struct {
	Id                      string
	Cnpj                    string
	RazaoSocial             string
	NomeFantasia            *string
	TipoInscricao           uint64
	ClassificacaoTributaria uint64
	NaturezaJuridica        uint64
	// IndCooperativa          *string
	// IndConstrutora          *string
	// IndDesoneracao          *string
	// Situacao                *string
	// DataCadastro            *time.Time
	// DataAtualizacao         *time.Time
	// UsuarioCadastro         *string
	// Telefone                *string
	// Email                   *string
	// Logradouro              *string
	// Numero                  *string
	// Complemento             *string
	// Bairro                  *string
	// Cidade                  *string
	// Uf                      *string
	// Cep                     *string
	// ProtocoloEsocial        *string
	// ReciboEsocial           *string
	// DataEnvioEsocial        *time.Time
	// StatusEsocial           *string
}
