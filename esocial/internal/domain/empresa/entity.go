package empresa

import (
	"time"
)

// Empresa represents a company entity
type Empresa struct {
	ID                      string     `json:"id" db:"id"`
	CNPJ                    string     `json:"cnpj" db:"cnpj"`
	RazaoSocial             string     `json:"razao_social" db:"razao_social"`
	NomeFantasia            *string    `json:"nome_fantasia,omitempty" db:"nome_fantasia"`
	TipoInscricao           string     `json:"tipo_inscricao" db:"tipo_inscricao"`
	ClassificacaoTributaria string     `json:"classificacao_tributaria" db:"classificacao_tributaria"`
	NaturezaJuridica        *string    `json:"natureza_juridica,omitempty" db:"natureza_juridica"`
	IndCooperativa          string     `json:"ind_cooperativa" db:"ind_cooperativa"`
	IndConstrutora          string     `json:"ind_construtora" db:"ind_construtora"`
	IndDesoneracao          string     `json:"ind_desoneracao" db:"ind_desoneracao"`
	Situacao                string     `json:"situacao" db:"situacao"`
	DataCadastro            time.Time  `json:"data_cadastro" db:"data_cadastro"`
	DataAtualizacao         time.Time  `json:"data_atualizacao" db:"data_atualizacao"`
	UsuarioCadastro         *string    `json:"usuario_cadastro,omitempty" db:"usuario_cadastro"`
	Telefone                *string    `json:"telefone,omitempty" db:"telefone"`
	Email                   *string    `json:"email,omitempty" db:"email"`
	Logradouro              *string    `json:"logradouro,omitempty" db:"logradouro"`
	Numero                  *string    `json:"numero,omitempty" db:"numero"`
	Complemento             *string    `json:"complemento,omitempty" db:"complemento"`
	Bairro                  *string    `json:"bairro,omitempty" db:"bairro"`
	Cidade                  *string    `json:"cidade,omitempty" db:"cidade"`
	UF                      *string    `json:"uf,omitempty" db:"uf"`
	CEP                     *string    `json:"cep,omitempty" db:"cep"`
	ProtocoloESocial        *string    `json:"protocolo_esocial,omitempty" db:"protocolo_esocial"`
	ReciboESocial           *string    `json:"recibo_esocial,omitempty" db:"recibo_esocial"`
	DataEnvioESocial        *time.Time `json:"data_envio_esocial,omitempty" db:"data_envio_esocial"`
	StatusESocial           *string    `json:"status_esocial,omitempty" db:"status_esocial"`
}

// Estabelecimento represents a company establishment/branch
type Estabelecimento struct {
	ID                   string     `json:"id" db:"id"`
	EmpresaID            string     `json:"empresa_id" db:"empresa_id"`
	TipoInscricao        string     `json:"tipo_inscricao" db:"tipo_inscricao"`
	NumeroInscricao      string     `json:"numero_inscricao" db:"numero_inscricao"`
	CNAEPrincipal        *string    `json:"cnae_principal,omitempty" db:"cnae_principal"`
	CNAEPreparatorio     *string    `json:"cnae_preparatorio,omitempty" db:"cnae_preparatorio"`
	AlvaraFuncionamento  *string    `json:"alvara_funcionamento,omitempty" db:"alvara_funcionamento"`
	DataInicioAtividades *time.Time `json:"data_inicio_atividades,omitempty" db:"data_inicio_atividades"`
	IndObra              string     `json:"ind_obra" db:"ind_obra"`
	Logradouro           *string    `json:"logradouro,omitempty" db:"logradouro"`
	Numero               *string    `json:"numero,omitempty" db:"numero"`
	Complemento          *string    `json:"complemento,omitempty" db:"complemento"`
	Bairro               *string    `json:"bairro,omitempty" db:"bairro"`
	Cidade               *string    `json:"cidade,omitempty" db:"cidade"`
	UF                   *string    `json:"uf,omitempty" db:"uf"`
	CEP                  *string    `json:"cep,omitempty" db:"cep"`
	Situacao             string     `json:"situacao" db:"situacao"`
	DataCadastro         time.Time  `json:"data_cadastro" db:"data_cadastro"`
}

// Rubrica represents a payroll rubric/item
type Rubrica struct {
	ID                 string     `json:"id" db:"id"`
	EmpresaID          string     `json:"empresa_id" db:"empresa_id"`
	Codigo             string     `json:"codigo" db:"codigo"`
	Descricao          string     `json:"descricao" db:"descricao"`
	NaturezaRubrica    string     `json:"natureza_rubrica" db:"natureza_rubrica"`
	TipoRubrica        string     `json:"tipo_rubrica" db:"tipo_rubrica"`
	CodINCCP           *string    `json:"cod_inccp,omitempty" db:"cod_inccp"`
	CodINCIRRF         *string    `json:"cod_incirrf,omitempty" db:"cod_incirrf"`
	CodINCFGTS         *string    `json:"cod_incfgts,omitempty" db:"cod_incfgts"`
	CodINCSind         *string    `json:"cod_incsind,omitempty" db:"cod_incsind"`
	Ativa              bool       `json:"ativa" db:"ativa"`
	DataCadastro       time.Time  `json:"data_cadastro" db:"data_cadastro"`
	DataInicioValidade time.Time  `json:"data_inicio_validade" db:"data_inicio_validade"`
	DataFimValidade    *time.Time `json:"data_fim_validade,omitempty" db:"data_fim_validade"`
}

// Lotacao represents a work location/lotation
type Lotacao struct {
	ID                 string     `json:"id" db:"id"`
	EmpresaID          string     `json:"empresa_id" db:"empresa_id"`
	Codigo             string     `json:"codigo" db:"codigo"`
	Descricao          string     `json:"descricao" db:"descricao"`
	TipoLotacao        string     `json:"tipo_lotacao" db:"tipo_lotacao"`
	FPAS               *string    `json:"fpas,omitempty" db:"fpas"`
	CodTerceiros       *string    `json:"cod_terceiros,omitempty" db:"cod_terceiros"`
	Ativa              bool       `json:"ativa" db:"ativa"`
	DataCadastro       time.Time  `json:"data_cadastro" db:"data_cadastro"`
	DataInicioValidade time.Time  `json:"data_inicio_validade" db:"data_inicio_validade"`
	DataFimValidade    *time.Time `json:"data_fim_validade,omitempty" db:"data_fim_validade"`
}

// Cargo represents a job position
type Cargo struct {
	ID                 string     `json:"id" db:"id"`
	EmpresaID          string     `json:"empresa_id" db:"empresa_id"`
	Codigo             string     `json:"codigo" db:"codigo"`
	Descricao          string     `json:"descricao" db:"descricao"`
	CBO                string     `json:"cbo" db:"cbo"`
	Ativo              bool       `json:"ativo" db:"ativo"`
	DataCadastro       time.Time  `json:"data_cadastro" db:"data_cadastro"`
	DataInicioValidade time.Time  `json:"data_inicio_validade" db:"data_inicio_validade"`
	DataFimValidade    *time.Time `json:"data_fim_validade,omitempty" db:"data_fim_validade"`
}
