package empresa

import (
	"time"
)

// CreateEmpresaRequest represents the request to create a company
type CreateEmpresaRequest struct {
	CNPJ                    string  `json:"cnpj" validate:"required,len=14"`
	RazaoSocial             string  `json:"razao_social" validate:"required,max=255"`
	NomeFantasia            *string `json:"nome_fantasia,omitempty" validate:"omitempty,max=255"`
	TipoInscricao           string  `json:"tipo_inscricao" validate:"required,oneof=1 2 3 4"`
	ClassificacaoTributaria string  `json:"classificacao_tributaria" validate:"required,len=2"`
	NaturezaJuridica        *string `json:"natureza_juridica,omitempty" validate:"omitempty,len=4"`
	IndCooperativa          string  `json:"ind_cooperativa" validate:"required,oneof=S N"`
	IndConstrutora          string  `json:"ind_construtora" validate:"required,oneof=S N"`
	IndDesoneracao          string  `json:"ind_desoneracao" validate:"required,oneof=S N"`
	UsuarioCadastro         *string `json:"usuario_cadastro,omitempty" validate:"omitempty,max=100"`
	Telefone                *string `json:"telefone,omitempty" validate:"omitempty,max=20"`
	Email                   *string `json:"email,omitempty" validate:"omitempty,email,max=255"`
	Logradouro              *string `json:"logradouro,omitempty" validate:"omitempty,max=255"`
	Numero                  *string `json:"numero,omitempty" validate:"omitempty,max=10"`
	Complemento             *string `json:"complemento,omitempty" validate:"omitempty,max=100"`
	Bairro                  *string `json:"bairro,omitempty" validate:"omitempty,max=100"`
	Cidade                  *string `json:"cidade,omitempty" validate:"omitempty,max=100"`
	UF                      *string `json:"uf,omitempty" validate:"omitempty,len=2"`
	CEP                     *string `json:"cep,omitempty" validate:"omitempty,len=8"`
}

// UpdateEmpresaRequest represents the request to update a company
type UpdateEmpresaRequest struct {
	RazaoSocial             *string    `json:"razao_social,omitempty" validate:"omitempty,max=255"`
	NomeFantasia            *string    `json:"nome_fantasia,omitempty" validate:"omitempty,max=255"`
	TipoInscricao           *string    `json:"tipo_inscricao,omitempty" validate:"omitempty,oneof=1 2 3 4"`
	ClassificacaoTributaria *string    `json:"classificacao_tributaria,omitempty" validate:"omitempty,len=2"`
	NaturezaJuridica        *string    `json:"natureza_juridica,omitempty" validate:"omitempty,len=4"`
	IndCooperativa          *string    `json:"ind_cooperativa,omitempty" validate:"omitempty,oneof=S N"`
	IndConstrutora          *string    `json:"ind_construtora,omitempty" validate:"omitempty,oneof=S N"`
	IndDesoneracao          *string    `json:"ind_desoneracao,omitempty" validate:"omitempty,oneof=S N"`
	Situacao                *string    `json:"situacao,omitempty" validate:"omitempty,oneof=ATIVA INATIVA"`
	Telefone                *string    `json:"telefone,omitempty" validate:"omitempty,max=20"`
	Email                   *string    `json:"email,omitempty" validate:"omitempty,email,max=255"`
	Logradouro              *string    `json:"logradouro,omitempty" validate:"omitempty,max=255"`
	Numero                  *string    `json:"numero,omitempty" validate:"omitempty,max=10"`
	Complemento             *string    `json:"complemento,omitempty" validate:"omitempty,max=100"`
	Bairro                  *string    `json:"bairro,omitempty" validate:"omitempty,max=100"`
	Cidade                  *string    `json:"cidade,omitempty" validate:"omitempty,max=100"`
	UF                      *string    `json:"uf,omitempty" validate:"omitempty,len=2"`
	CEP                     *string    `json:"cep,omitempty" validate:"omitempty,len=8"`
	ProtocoloESocial        *string    `json:"protocolo_esocial,omitempty" validate:"omitempty,max=50"`
	ReciboESocial           *string    `json:"recibo_esocial,omitempty" validate:"omitempty,max=50"`
	DataEnvioESocial        *time.Time `json:"data_envio_esocial,omitempty"`
	StatusESocial           *string    `json:"status_esocial,omitempty" validate:"omitempty,max=20"`
}

// EmpresaResponse represents the company response
type EmpresaResponse struct {
	ID                      string     `json:"id"`
	CNPJ                    string     `json:"cnpj"`
	RazaoSocial             string     `json:"razao_social"`
	NomeFantasia            *string    `json:"nome_fantasia,omitempty"`
	TipoInscricao           string     `json:"tipo_inscricao"`
	ClassificacaoTributaria string     `json:"classificacao_tributaria"`
	NaturezaJuridica        *string    `json:"natureza_juridica,omitempty"`
	IndCooperativa          string     `json:"ind_cooperativa"`
	IndConstrutora          string     `json:"ind_construtora"`
	IndDesoneracao          string     `json:"ind_desoneracao"`
	Situacao                string     `json:"situacao"`
	DataCadastro            time.Time  `json:"data_cadastro"`
	DataAtualizacao         time.Time  `json:"data_atualizacao"`
	UsuarioCadastro         *string    `json:"usuario_cadastro,omitempty"`
	Telefone                *string    `json:"telefone,omitempty"`
	Email                   *string    `json:"email,omitempty"`
	Logradouro              *string    `json:"logradouro,omitempty"`
	Numero                  *string    `json:"numero,omitempty"`
	Complemento             *string    `json:"complemento,omitempty"`
	Bairro                  *string    `json:"bairro,omitempty"`
	Cidade                  *string    `json:"cidade,omitempty"`
	UF                      *string    `json:"uf,omitempty"`
	CEP                     *string    `json:"cep,omitempty"`
	ProtocoloESocial        *string    `json:"protocolo_esocial,omitempty"`
	ReciboESocial           *string    `json:"recibo_esocial,omitempty"`
	DataEnvioESocial        *time.Time `json:"data_envio_esocial,omitempty"`
	StatusESocial           *string    `json:"status_esocial,omitempty"`
}

// CreateEstabelecimentoRequest represents the request to create an establishment
type CreateEstabelecimentoRequest struct {
	EmpresaID            string     `json:"empresa_id" validate:"required"`
	TipoInscricao        string     `json:"tipo_inscricao" validate:"required,oneof=1 2 3 4"`
	NumeroInscricao      string     `json:"numero_inscricao" validate:"required,max=15"`
	CNAEPrincipal        *string    `json:"cnae_principal,omitempty" validate:"omitempty,len=7"`
	CNAEPreparatorio     *string    `json:"cnae_preparatorio,omitempty" validate:"omitempty,len=7"`
	AlvaraFuncionamento  *string    `json:"alvara_funcionamento,omitempty" validate:"omitempty,max=50"`
	DataInicioAtividades *time.Time `json:"data_inicio_atividades,omitempty"`
	IndObra              string     `json:"ind_obra" validate:"required,oneof=S N"`
	Logradouro           *string    `json:"logradouro,omitempty" validate:"omitempty,max=255"`
	Numero               *string    `json:"numero,omitempty" validate:"omitempty,max=10"`
	Complemento          *string    `json:"complemento,omitempty" validate:"omitempty,max=100"`
	Bairro               *string    `json:"bairro,omitempty" validate:"omitempty,max=100"`
	Cidade               *string    `json:"cidade,omitempty" validate:"omitempty,max=100"`
	UF                   *string    `json:"uf,omitempty" validate:"omitempty,len=2"`
	CEP                  *string    `json:"cep,omitempty" validate:"omitempty,len=8"`
}

// UpdateEstabelecimentoRequest represents the request to update an establishment
type UpdateEstabelecimentoRequest struct {
	TipoInscricao        *string    `json:"tipo_inscricao,omitempty" validate:"omitempty,oneof=1 2 3 4"`
	NumeroInscricao      *string    `json:"numero_inscricao,omitempty" validate:"omitempty,max=15"`
	CNAEPrincipal        *string    `json:"cnae_principal,omitempty" validate:"omitempty,len=7"`
	CNAEPreparatorio     *string    `json:"cnae_preparatorio,omitempty" validate:"omitempty,len=7"`
	AlvaraFuncionamento  *string    `json:"alvara_funcionamento,omitempty" validate:"omitempty,max=50"`
	DataInicioAtividades *time.Time `json:"data_inicio_atividades,omitempty"`
	IndObra              *string    `json:"ind_obra,omitempty" validate:"omitempty,oneof=S N"`
	Logradouro           *string    `json:"logradouro,omitempty" validate:"omitempty,max=255"`
	Numero               *string    `json:"numero,omitempty" validate:"omitempty,max=10"`
	Complemento          *string    `json:"complemento,omitempty" validate:"omitempty,max=100"`
	Bairro               *string    `json:"bairro,omitempty" validate:"omitempty,max=100"`
	Cidade               *string    `json:"cidade,omitempty" validate:"omitempty,max=100"`
	UF                   *string    `json:"uf,omitempty" validate:"omitempty,len=2"`
	CEP                  *string    `json:"cep,omitempty" validate:"omitempty,len=8"`
	Situacao             *string    `json:"situacao,omitempty" validate:"omitempty,oneof=ATIVO INATIVO"`
}

// EstabelecimentoResponse represents the establishment response
type EstabelecimentoResponse struct {
	ID                   string     `json:"id"`
	EmpresaID            string     `json:"empresa_id"`
	TipoInscricao        string     `json:"tipo_inscricao"`
	NumeroInscricao      string     `json:"numero_inscricao"`
	CNAEPrincipal        *string    `json:"cnae_principal,omitempty"`
	CNAEPreparatorio     *string    `json:"cnae_preparatorio,omitempty"`
	AlvaraFuncionamento  *string    `json:"alvara_funcionamento,omitempty"`
	DataInicioAtividades *time.Time `json:"data_inicio_atividades,omitempty"`
	IndObra              string     `json:"ind_obra"`
	Logradouro           *string    `json:"logradouro,omitempty"`
	Numero               *string    `json:"numero,omitempty"`
	Complemento          *string    `json:"complemento,omitempty"`
	Bairro               *string    `json:"bairro,omitempty"`
	Cidade               *string    `json:"cidade,omitempty"`
	UF                   *string    `json:"uf,omitempty"`
	CEP                  *string    `json:"cep,omitempty"`
	Situacao             string     `json:"situacao"`
	DataCadastro         time.Time  `json:"data_cadastro"`
}

// CreateRubricaRequest represents the request to create a rubric
type CreateRubricaRequest struct {
	EmpresaID          string     `json:"empresa_id" validate:"required"`
	Codigo             string     `json:"codigo" validate:"required,max=30"`
	Descricao          string     `json:"descricao" validate:"required,max=255"`
	NaturezaRubrica    string     `json:"natureza_rubrica" validate:"required,len=4"`
	TipoRubrica        string     `json:"tipo_rubrica" validate:"required,oneof=D C"`
	CodINCCP           *string    `json:"cod_inccp,omitempty" validate:"omitempty,len=2"`
	CodINCIRRF         *string    `json:"cod_incirrf,omitempty" validate:"omitempty,len=2"`
	CodINCFGTS         *string    `json:"cod_incfgts,omitempty" validate:"omitempty,len=2"`
	CodINCSind         *string    `json:"cod_incsind,omitempty" validate:"omitempty,len=2"`
	DataInicioValidade time.Time  `json:"data_inicio_validade" validate:"required"`
	DataFimValidade    *time.Time `json:"data_fim_validade,omitempty"`
}

// UpdateRubricaRequest represents the request to update a rubric
type UpdateRubricaRequest struct {
	Codigo             *string    `json:"codigo,omitempty" validate:"omitempty,max=30"`
	Descricao          *string    `json:"descricao,omitempty" validate:"omitempty,max=255"`
	NaturezaRubrica    *string    `json:"natureza_rubrica,omitempty" validate:"omitempty,len=4"`
	TipoRubrica        *string    `json:"tipo_rubrica,omitempty" validate:"omitempty,oneof=D C"`
	CodINCCP           *string    `json:"cod_inccp,omitempty" validate:"omitempty,len=2"`
	CodINCIRRF         *string    `json:"cod_incirrf,omitempty" validate:"omitempty,len=2"`
	CodINCFGTS         *string    `json:"cod_incfgts,omitempty" validate:"omitempty,len=2"`
	CodINCSind         *string    `json:"cod_incsind,omitempty" validate:"omitempty,len=2"`
	Ativa              *bool      `json:"ativa,omitempty"`
	DataInicioValidade *time.Time `json:"data_inicio_validade,omitempty"`
	DataFimValidade    *time.Time `json:"data_fim_validade,omitempty"`
}

// RubricaResponse represents the rubric response
type RubricaResponse struct {
	ID                 string     `json:"id"`
	EmpresaID          string     `json:"empresa_id"`
	Codigo             string     `json:"codigo"`
	Descricao          string     `json:"descricao"`
	NaturezaRubrica    string     `json:"natureza_rubrica"`
	TipoRubrica        string     `json:"tipo_rubrica"`
	CodINCCP           *string    `json:"cod_inccp,omitempty"`
	CodINCIRRF         *string    `json:"cod_incirrf,omitempty"`
	CodINCFGTS         *string    `json:"cod_incfgts,omitempty"`
	CodINCSind         *string    `json:"cod_incsind,omitempty"`
	Ativa              bool       `json:"ativa"`
	DataCadastro       time.Time  `json:"data_cadastro"`
	DataInicioValidade time.Time  `json:"data_inicio_validade"`
	DataFimValidade    *time.Time `json:"data_fim_validade,omitempty"`
}

// CreateLotacaoRequest represents the request to create a lotation
type CreateLotacaoRequest struct {
	EmpresaID          string     `json:"empresa_id" validate:"required"`
	Codigo             string     `json:"codigo" validate:"required,max=30"`
	Descricao          string     `json:"descricao" validate:"required,max=255"`
	TipoLotacao        string     `json:"tipo_lotacao" validate:"required,len=2"`
	FPAS               *string    `json:"fpas,omitempty" validate:"omitempty,len=3"`
	CodTerceiros       *string    `json:"cod_terceiros,omitempty" validate:"omitempty,len=4"`
	DataInicioValidade time.Time  `json:"data_inicio_validade" validate:"required"`
	DataFimValidade    *time.Time `json:"data_fim_validade,omitempty"`
}

// UpdateLotacaoRequest represents the request to update a lotation
type UpdateLotacaoRequest struct {
	Codigo             *string    `json:"codigo,omitempty" validate:"omitempty,max=30"`
	Descricao          *string    `json:"descricao,omitempty" validate:"omitempty,max=255"`
	TipoLotacao        *string    `json:"tipo_lotacao,omitempty" validate:"omitempty,len=2"`
	FPAS               *string    `json:"fpas,omitempty" validate:"omitempty,len=3"`
	CodTerceiros       *string    `json:"cod_terceiros,omitempty" validate:"omitempty,len=4"`
	Ativa              *bool      `json:"ativa,omitempty"`
	DataInicioValidade *time.Time `json:"data_inicio_validade,omitempty"`
	DataFimValidade    *time.Time `json:"data_fim_validade,omitempty"`
}

// LotacaoResponse represents the lotation response
type LotacaoResponse struct {
	ID                 string     `json:"id"`
	EmpresaID          string     `json:"empresa_id"`
	Codigo             string     `json:"codigo"`
	Descricao          string     `json:"descricao"`
	TipoLotacao        string     `json:"tipo_lotacao"`
	FPAS               *string    `json:"fpas,omitempty"`
	CodTerceiros       *string    `json:"cod_terceiros,omitempty"`
	Ativa              bool       `json:"ativa"`
	DataCadastro       time.Time  `json:"data_cadastro"`
	DataInicioValidade time.Time  `json:"data_inicio_validade"`
	DataFimValidade    *time.Time `json:"data_fim_validade,omitempty"`
}

// CreateCargoRequest represents the request to create a position
type CreateCargoRequest struct {
	EmpresaID          string     `json:"empresa_id" validate:"required"`
	Codigo             string     `json:"codigo" validate:"required,max=30"`
	Descricao          string     `json:"descricao" validate:"required,max=255"`
	CBO                string     `json:"cbo" validate:"required,len=6"`
	DataInicioValidade time.Time  `json:"data_inicio_validade" validate:"required"`
	DataFimValidade    *time.Time `json:"data_fim_validade,omitempty"`
}

// UpdateCargoRequest represents the request to update a position
type UpdateCargoRequest struct {
	Codigo             *string    `json:"codigo,omitempty" validate:"omitempty,max=30"`
	Descricao          *string    `json:"descricao,omitempty" validate:"omitempty,max=255"`
	CBO                *string    `json:"cbo,omitempty" validate:"omitempty,len=6"`
	Ativo              *bool      `json:"ativo,omitempty"`
	DataInicioValidade *time.Time `json:"data_inicio_validade,omitempty"`
	DataFimValidade    *time.Time `json:"data_fim_validade,omitempty"`
}

// CargoResponse represents the position response
type CargoResponse struct {
	ID                 string     `json:"id"`
	EmpresaID          string     `json:"empresa_id"`
	Codigo             string     `json:"codigo"`
	Descricao          string     `json:"descricao"`
	CBO                string     `json:"cbo"`
	Ativo              bool       `json:"ativo"`
	DataCadastro       time.Time  `json:"data_cadastro"`
	DataInicioValidade time.Time  `json:"data_inicio_validade"`
	DataFimValidade    *time.Time `json:"data_fim_validade,omitempty"`
}

// ListRequest represents common list/filter parameters
type ListRequest struct {
	Page      int     `json:"page,omitempty" validate:"omitempty,min=1"`
	Limit     int     `json:"limit,omitempty" validate:"omitempty,min=1,max=100"`
	EmpresaID *string `json:"empresa_id,omitempty"`
	Search    *string `json:"search,omitempty"`
	Ativo     *bool   `json:"ativo,omitempty"`
}

// ListResponse represents a paginated list response
type ListResponse[T any] struct {
	Data       []T   `json:"data"`
	Total      int64 `json:"total"`
	Page       int   `json:"page"`
	Limit      int   `json:"limit"`
	TotalPages int   `json:"total_pages"`
}
