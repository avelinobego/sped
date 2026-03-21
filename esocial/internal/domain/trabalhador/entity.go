package trabalhador

import (
	"time"
)

// Trabalhador represents a worker entity
type Trabalhador struct {
	ID                    string     `json:"id" db:"id"`
	EmpresaID             string     `json:"empresa_id" db:"empresa_id"`
	CPF                   string     `json:"cpf" db:"cpf"`
	NIS                   *string    `json:"nis,omitempty" db:"nis"`
	Nome                  string     `json:"nome" db:"nome"`
	NomeSocial            *string    `json:"nome_social,omitempty" db:"nome_social"`
	DataNascimento        time.Time  `json:"data_nascimento" db:"data_nascimento"`
	Sexo                  string     `json:"sexo" db:"sexo"`
	RacaCor               *string    `json:"raca_cor,omitempty" db:"raca_cor"`
	EstadoCivil           *string    `json:"estado_civil,omitempty" db:"estado_civil"`
	GrauInstrucao         *string    `json:"grau_instrucao,omitempty" db:"grau_instrucao"`
	PaisNascimento        string     `json:"pais_nascimento" db:"pais_nascimento"`
	PaisNacionalidade     string     `json:"pais_nacionalidade" db:"pais_nacionalidade"`
	NumeroCTPS            *string    `json:"numero_ctps,omitempty" db:"numero_ctps"`
	SerieCTPS             *string    `json:"serie_ctps,omitempty" db:"serie_ctps"`
	UFCTPS                *string    `json:"uf_ctps,omitempty" db:"uf_ctps"`
	DataEmissaoCTPS       *time.Time `json:"data_emissao_ctps,omitempty" db:"data_emissao_ctps"`
	NumeroRG              *string    `json:"numero_rg,omitempty" db:"numero_rg"`
	OrgaoEmissorRG        *string    `json:"orgao_emissor_rg,omitempty" db:"orgao_emissor_rg"`
	UFRG                  *string    `json:"uf_rg,omitempty" db:"uf_rg"`
	DataEmissaoRG         *time.Time `json:"data_emissao_rg,omitempty" db:"data_emissao_rg"`
	NumeroCNH             *string    `json:"numero_cnh,omitempty" db:"numero_cnh"`
	CategoriaCNH          *string    `json:"categoria_cnh,omitempty" db:"categoria_cnh"`
	ValidadeCNH           *time.Time `json:"validade_cnh,omitempty" db:"validade_cnh"`
	NumeroRNE             *string    `json:"numero_rne,omitempty" db:"numero_rne"`
	Telefone              *string    `json:"telefone,omitempty" db:"telefone"`
	Celular               *string    `json:"celular,omitempty" db:"celular"`
	Email                 *string    `json:"email,omitempty" db:"email"`
	CEP                   *string    `json:"cep,omitempty" db:"cep"`
	Logradouro            *string    `json:"logradouro,omitempty" db:"logradouro"`
	Numero                *string    `json:"numero,omitempty" db:"numero"`
	Complemento           *string    `json:"complemento,omitempty" db:"complemento"`
	Bairro                *string    `json:"bairro,omitempty" db:"bairro"`
	Cidade                *string    `json:"cidade,omitempty" db:"cidade"`
	UF                    *string    `json:"uf,omitempty" db:"uf"`
	Banco                 *string    `json:"banco,omitempty" db:"banco"`
	Agencia               *string    `json:"agencia,omitempty" db:"agencia"`
	Conta                 *string    `json:"conta,omitempty" db:"conta"`
	TipoConta             *string    `json:"tipo_conta,omitempty" db:"tipo_conta"`
	PossuiDeficiencia     bool       `json:"possui_deficiencia" db:"possui_deficiencia"`
	TipoDeficiencia       *string    `json:"tipo_deficiencia,omitempty" db:"tipo_deficiencia"`
	ObservacaoDeficiencia *string    `json:"observacao_deficiencia,omitempty" db:"observacao_deficiencia"`
	Situacao              string     `json:"situacao" db:"situacao"`
	DataCadastro          time.Time  `json:"data_cadastro" db:"data_cadastro"`
	DataAtualizacao       time.Time  `json:"data_atualizacao" db:"data_atualizacao"`
	FotoURL               *string    `json:"foto_url,omitempty" db:"foto_url"`
}
