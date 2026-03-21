package trabalhador

import (
	"time"
)

// CreateTrabalhadorRequest represents the request to create a worker
type CreateTrabalhadorRequest struct {
	EmpresaID             string     `json:"empresa_id" validate:"required"`
	CPF                   string     `json:"cpf" validate:"required,len=11"`
	NIS                   *string    `json:"nis,omitempty" validate:"omitempty,len=11"`
	Nome                  string     `json:"nome" validate:"required,max=255"`
	NomeSocial            *string    `json:"nome_social,omitempty" validate:"omitempty,max=255"`
	DataNascimento        time.Time  `json:"data_nascimento" validate:"required"`
	Sexo                  string     `json:"sexo" validate:"required,oneof=M F"`
	RacaCor               *string    `json:"raca_cor,omitempty" validate:"omitempty,oneof=1 2 3 4 5 6"`
	EstadoCivil           *string    `json:"estado_civil,omitempty" validate:"omitempty,oneof=1 2 3 4 5"`
	GrauInstrucao         *string    `json:"grau_instrucao,omitempty" validate:"omitempty,oneof=01 02 03 04 05 06 07 08 09 10 11 12"`
	PaisNascimento        string     `json:"pais_nascimento" validate:"required,len=3"`
	PaisNacionalidade     string     `json:"pais_nacionalidade" validate:"required,len=3"`
	NumeroCTPS            *string    `json:"numero_ctps,omitempty" validate:"omitempty,max=11"`
	SerieCTPS             *string    `json:"serie_ctps,omitempty" validate:"omitempty,max=5"`
	UFCTPS                *string    `json:"uf_ctps,omitempty" validate:"omitempty,len=2"`
	DataEmissaoCTPS       *time.Time `json:"data_emissao_ctps,omitempty"`
	NumeroRG              *string    `json:"numero_rg,omitempty" validate:"omitempty,max=20"`
	OrgaoEmissorRG        *string    `json:"orgao_emissor_rg,omitempty" validate:"omitempty,max=10"`
	UFRG                  *string    `json:"uf_rg,omitempty" validate:"omitempty,len=2"`
	DataEmissaoRG         *time.Time `json:"data_emissao_rg,omitempty"`
	NumeroCNH             *string    `json:"numero_cnh,omitempty" validate:"omitempty,max=11"`
	CategoriaCNH          *string    `json:"categoria_cnh,omitempty" validate:"omitempty,oneof=A B C D E AB AC AD AE"`
	ValidadeCNH           *time.Time `json:"validade_cnh,omitempty"`
	NumeroRNE             *string    `json:"numero_rne,omitempty" validate:"omitempty,max=20"`
	Telefone              *string    `json:"telefone,omitempty" validate:"omitempty,max=20"`
	Celular               *string    `json:"celular,omitempty" validate:"omitempty,max=20"`
	Email                 *string    `json:"email,omitempty" validate:"omitempty,email,max=255"`
	CEP                   *string    `json:"cep,omitempty" validate:"omitempty,len=8"`
	Logradouro            *string    `json:"logradouro,omitempty" validate:"omitempty,max=255"`
	Numero                *string    `json:"numero,omitempty" validate:"omitempty,max=10"`
	Complemento           *string    `json:"complemento,omitempty" validate:"omitempty,max=100"`
	Bairro                *string    `json:"bairro,omitempty" validate:"omitempty,max=100"`
	Cidade                *string    `json:"cidade,omitempty" validate:"omitempty,max=100"`
	UF                    *string    `json:"uf,omitempty" validate:"omitempty,len=2"`
	Banco                 *string    `json:"banco,omitempty" validate:"omitempty,len=3"`
	Agencia               *string    `json:"agencia,omitempty" validate:"omitempty,max=5"`
	Conta                 *string    `json:"conta,omitempty" validate:"omitempty,max=20"`
	TipoConta             *string    `json:"tipo_conta,omitempty" validate:"omitempty,oneof=1 2"`
	PossuiDeficiencia     bool       `json:"possui_deficiencia"`
	TipoDeficiencia       *string    `json:"tipo_deficiencia,omitempty" validate:"omitempty,oneof=1 2 3 4 5 6"`
	ObservacaoDeficiencia *string    `json:"observacao_deficiencia,omitempty" validate:"omitempty,max=255"`
	FotoURL               *string    `json:"foto_url,omitempty" validate:"omitempty,max=500"`
}

// UpdateTrabalhadorRequest represents the request to update a worker
type UpdateTrabalhadorRequest struct {
	NIS                   *string    `json:"nis,omitempty" validate:"omitempty,len=11"`
	Nome                  *string    `json:"nome,omitempty" validate:"omitempty,max=255"`
	NomeSocial            *string    `json:"nome_social,omitempty" validate:"omitempty,max=255"`
	DataNascimento        *time.Time `json:"data_nascimento,omitempty"`
	Sexo                  *string    `json:"sexo,omitempty" validate:"omitempty,oneof=M F"`
	RacaCor               *string    `json:"raca_cor,omitempty" validate:"omitempty,oneof=1 2 3 4 5 6"`
	EstadoCivil           *string    `json:"estado_civil,omitempty" validate:"omitempty,oneof=1 2 3 4 5"`
	GrauInstrucao         *string    `json:"grau_instrucao,omitempty" validate:"omitempty,oneof=01 02 03 04 05 06 07 08 09 10 11 12"`
	PaisNascimento        *string    `json:"pais_nascimento,omitempty" validate:"omitempty,len=3"`
	PaisNacionalidade     *string    `json:"pais_nacionalidade,omitempty" validate:"omitempty,len=3"`
	NumeroCTPS            *string    `json:"numero_ctps,omitempty" validate:"omitempty,max=11"`
	SerieCTPS             *string    `json:"serie_ctps,omitempty" validate:"omitempty,max=5"`
	UFCTPS                *string    `json:"uf_ctps,omitempty" validate:"omitempty,len=2"`
	DataEmissaoCTPS       *time.Time `json:"data_emissao_ctps,omitempty"`
	NumeroRG              *string    `json:"numero_rg,omitempty" validate:"omitempty,max=20"`
	OrgaoEmissorRG        *string    `json:"orgao_emissor_rg,omitempty" validate:"omitempty,max=10"`
	UFRG                  *string    `json:"uf_rg,omitempty" validate:"omitempty,len=2"`
	DataEmissaoRG         *time.Time `json:"data_emissao_rg,omitempty"`
	NumeroCNH             *string    `json:"numero_cnh,omitempty" validate:"omitempty,max=11"`
	CategoriaCNH          *string    `json:"categoria_cnh,omitempty" validate:"omitempty,oneof=A B C D E AB AC AD AE"`
	ValidadeCNH           *time.Time `json:"validade_cnh,omitempty"`
	NumeroRNE             *string    `json:"numero_rne,omitempty" validate:"omitempty,max=20"`
	Telefone              *string    `json:"telefone,omitempty" validate:"omitempty,max=20"`
	Celular               *string    `json:"celular,omitempty" validate:"omitempty,max=20"`
	Email                 *string    `json:"email,omitempty" validate:"omitempty,email,max=255"`
	CEP                   *string    `json:"cep,omitempty" validate:"omitempty,len=8"`
	Logradouro            *string    `json:"logradouro,omitempty" validate:"omitempty,max=255"`
	Numero                *string    `json:"numero,omitempty" validate:"omitempty,max=10"`
	Complemento           *string    `json:"complemento,omitempty" validate:"omitempty,max=100"`
	Bairro                *string    `json:"bairro,omitempty" validate:"omitempty,max=100"`
	Cidade                *string    `json:"cidade,omitempty" validate:"omitempty,max=100"`
	UF                    *string    `json:"uf,omitempty" validate:"omitempty,len=2"`
	Banco                 *string    `json:"banco,omitempty" validate:"omitempty,len=3"`
	Agencia               *string    `json:"agencia,omitempty" validate:"omitempty,max=5"`
	Conta                 *string    `json:"conta,omitempty" validate:"omitempty,max=20"`
	TipoConta             *string    `json:"tipo_conta,omitempty" validate:"omitempty,oneof=1 2"`
	PossuiDeficiencia     *bool      `json:"possui_deficiencia,omitempty"`
	TipoDeficiencia       *string    `json:"tipo_deficiencia,omitempty" validate:"omitempty,oneof=1 2 3 4 5 6"`
	ObservacaoDeficiencia *string    `json:"observacao_deficiencia,omitempty" validate:"omitempty,max=255"`
	Situacao              *string    `json:"situacao,omitempty" validate:"omitempty,oneof=ATIVO INATIVO"`
	FotoURL               *string    `json:"foto_url,omitempty" validate:"omitempty,max=500"`
}

// TrabalhadorResponse represents the worker response
type TrabalhadorResponse struct {
	ID                    string     `json:"id"`
	EmpresaID             string     `json:"empresa_id"`
	CPF                   string     `json:"cpf"`
	NIS                   *string    `json:"nis,omitempty"`
	Nome                  string     `json:"nome"`
	NomeSocial            *string    `json:"nome_social,omitempty"`
	DataNascimento        time.Time  `json:"data_nascimento"`
	Sexo                  string     `json:"sexo"`
	RacaCor               *string    `json:"raca_cor,omitempty"`
	EstadoCivil           *string    `json:"estado_civil,omitempty"`
	GrauInstrucao         *string    `json:"grau_instrucao,omitempty"`
	PaisNascimento        string     `json:"pais_nascimento"`
	PaisNacionalidade     string     `json:"pais_nacionalidade"`
	NumeroCTPS            *string    `json:"numero_ctps,omitempty"`
	SerieCTPS             *string    `json:"serie_ctps,omitempty"`
	UFCTPS                *string    `json:"uf_ctps,omitempty"`
	DataEmissaoCTPS       *time.Time `json:"data_emissao_ctps,omitempty"`
	NumeroRG              *string    `json:"numero_rg,omitempty"`
	OrgaoEmissorRG        *string    `json:"orgao_emissor_rg,omitempty"`
	UFRG                  *string    `json:"uf_rg,omitempty"`
	DataEmissaoRG         *time.Time `json:"data_emissao_rg,omitempty"`
	NumeroCNH             *string    `json:"numero_cnh,omitempty"`
	CategoriaCNH          *string    `json:"categoria_cnh,omitempty"`
	ValidadeCNH           *time.Time `json:"validade_cnh,omitempty"`
	NumeroRNE             *string    `json:"numero_rne,omitempty"`
	Telefone              *string    `json:"telefone,omitempty"`
	Celular               *string    `json:"celular,omitempty"`
	Email                 *string    `json:"email,omitempty"`
	CEP                   *string    `json:"cep,omitempty"`
	Logradouro            *string    `json:"logradouro,omitempty"`
	Numero                *string    `json:"numero,omitempty"`
	Complemento           *string    `json:"complemento,omitempty"`
	Bairro                *string    `json:"bairro,omitempty"`
	Cidade                *string    `json:"cidade,omitempty"`
	UF                    *string    `json:"uf,omitempty"`
	Banco                 *string    `json:"banco,omitempty"`
	Agencia               *string    `json:"agencia,omitempty"`
	Conta                 *string    `json:"conta,omitempty"`
	TipoConta             *string    `json:"tipo_conta,omitempty"`
	PossuiDeficiencia     bool       `json:"possui_deficiencia"`
	TipoDeficiencia       *string    `json:"tipo_deficiencia,omitempty"`
	ObservacaoDeficiencia *string    `json:"observacao_deficiencia,omitempty"`
	Situacao              string     `json:"situacao"`
	DataCadastro          time.Time  `json:"data_cadastro"`
	DataAtualizacao       time.Time  `json:"data_atualizacao"`
	FotoURL               *string    `json:"foto_url,omitempty"`
}

// ListRequest represents common list/filter parameters
type ListRequest struct {
	Page              int     `json:"page,omitempty" validate:"omitempty,min=1"`
	Limit             int     `json:"limit,omitempty" validate:"omitempty,min=1,max=100"`
	EmpresaID         *string `json:"empresa_id,omitempty"`
	Search            *string `json:"search,omitempty"`
	Situacao          *string `json:"situacao,omitempty"`
	Sexo              *string `json:"sexo,omitempty"`
	PossuiDeficiencia *bool   `json:"possui_deficiencia,omitempty"`
}

// ListResponse represents a paginated list response
type ListResponse[T any] struct {
	Data       []T   `json:"data"`
	Total      int64 `json:"total"`
	Page       int   `json:"page"`
	Limit      int   `json:"limit"`
	TotalPages int   `json:"total_pages"`
}
