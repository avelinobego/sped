package empregador

import (
	"time"
)

// Empresas maps to the 'empresas' table.
type Empresas struct {
	Id                      string     `json:"id" db:"id"`
	Cnpj                    string     `json:"cnpj" db:"cnpj"`
	RazaoSocial             string     `json:"razao_social" db:"razao_social"`
	NomeFantasia            *string    `json:"nome_fantasia" db:"nome_fantasia"`
	TipoInscricao           string     `json:"tipo_inscricao" db:"tipo_inscricao"`
	ClassificacaoTributaria string     `json:"classificacao_tributaria" db:"classificacao_tributaria"`
	NaturezaJuridica        *string    `json:"natureza_juridica" db:"natureza_juridica"`
	IndCooperativa          *string    `json:"ind_cooperativa" db:"ind_cooperativa"`
	IndConstrutora          *string    `json:"ind_construtora" db:"ind_construtora"`
	IndDesoneracao          *string    `json:"ind_desoneracao" db:"ind_desoneracao"`
	Situacao                *string    `json:"situacao" db:"situacao"`
	DataCadastro            *time.Time `json:"data_cadastro" db:"data_cadastro"`
	DataAtualizacao         *time.Time `json:"data_atualizacao" db:"data_atualizacao"`
	UsuarioCadastro         *string    `json:"usuario_cadastro" db:"usuario_cadastro"`
	Telefone                *string    `json:"telefone" db:"telefone"`
	Email                   *string    `json:"email" db:"email"`
	Logradouro              *string    `json:"logradouro" db:"logradouro"`
	Numero                  *string    `json:"numero" db:"numero"`
	Complemento             *string    `json:"complemento" db:"complemento"`
	Bairro                  *string    `json:"bairro" db:"bairro"`
	Cidade                  *string    `json:"cidade" db:"cidade"`
	Uf                      *string    `json:"uf" db:"uf"`
	Cep                     *string    `json:"cep" db:"cep"`
	ProtocoloEsocial        *string    `json:"protocolo_esocial" db:"protocolo_esocial"`
	ReciboEsocial           *string    `json:"recibo_esocial" db:"recibo_esocial"`
	DataEnvioEsocial        *time.Time `json:"data_envio_esocial" db:"data_envio_esocial"`
	StatusEsocial           *string    `json:"status_esocial" db:"status_esocial"`
}
