package xml

import "sped/esocial/pkg/types"

// ============================================================================
// EvtInfoEmpregador - S-1000: Informações do Empregador
// ============================================================================

// EvtInfoEmpregador representa o evento S-1000 (Informações do Empregador)
// Estrutura raiz para registro de dados cadastrais da empresa/empregador (matriz)
type EvtInfoEmpregador struct {
	// Id é um atributo XML que identifica o evento
	Id string `xml:"id,attr,omitempty"`

	// IdeEmpregador contém a identificação do empregador
	IdeEmpregador IdeEmpregador `xml:"ideEmpregador"`

	// InfoCadastro contém os dados cadastrais do empregador
	InfoCadastro InfoCadastroEmpregador `xml:"infoCadastro"`
}

// ============================================================================
// IdeEmpregador - Identificação do Empregador
// ============================================================================

// IdeEmpregador contém a identificação do empregador (tipo e número de inscrição)
type IdeEmpregador struct {
	// TpInsc é o tipo de inscrição (1=CNPJ, 2=CPF, 3=CAEPF, 4=CNO)
	TpInsc int `xml:"tpInsc"`

	// NrInsc é o número de inscrição sem formatação
	NrInsc string `xml:"nrInsc"`
}

// ============================================================================
// InfoCadastroEmpregador - Informações Cadastrais do Empregador
// ============================================================================

// InfoCadastroEmpregador contém os dados cadastrais completos do empregador
type InfoCadastroEmpregador struct {
	// DtCadastro é a data do cadastro (YYYY-MM-DD)
	DtCadastro types.DataApenas `xml:"dtCadastro"`

	// InfoPJ contém dados da pessoa jurídica (obrigatório para TpInsc=1/3/4)
	InfoPJ *InfoPJ `xml:"infoPJ,omitempty"`

	// InfoPF contém dados da pessoa física (obrigatório para TpInsc=2)
	InfoPF *InfoPF `xml:"infoPF,omitempty"`

	// EnderEmpregador contém o endereço do empregador
	EnderEmpregador Endereco `xml:"enderEmpregador"`

	// InfoContato contém informações de contato (opcional)
	InfoContato *Contato `xml:"infoContato,omitempty"`

	// InfoComplementar contém informações complementares (opcional)
	InfoComplementar *InfoComplementar `xml:"infoComplementar,omitempty"`
}

// ============================================================================
// InfoPJ - Dados da Pessoa Jurídica
// ============================================================================

// InfoPJ contém dados específicos de uma empresa (pessoa jurídica)
type InfoPJ struct {
	// RazaoSocial é a razão social da empresa (4-180 chars)
	RazaoSocial string `xml:"razaoSocial"`

	// NomeFantasia é o nome fantasia da empresa (opcional, 0-180 chars)
	NomeFantasia string `xml:"nomeFantasia,omitempty"`

	// DsCategEmpregador é a categoria do empregador (Tabela 07: 01-10)
	DsCategEmpregador string `xml:"dsCategEmpregador"`

	// ClassTributaria é a classificação tributária (Tabela 52: 01-11)
	ClassTributaria string `xml:"classTributaria"`

	// IndConstrutora indica se a empresa é construtora de imóveis (opcional, 1=Sim, 2=Não)
	IndConstrutora types.SimNao `xml:"indConstrutora,omitempty"`

	// IndDesFolha indica se a empresa é desonerada (opcional, 1=Sim, 2=Não)
	IndDesFolha types.SimNao `xml:"indDesFolha,omitempty"`

	// IndCooperativa indica se a empresa é cooperativa (opcional, 1=Sim, 2=Não)
	IndCooperativa types.SimNao `xml:"indCooperativa,omitempty"`
}

// ============================================================================
// InfoPF - Dados da Pessoa Física
// ============================================================================

// InfoPF contém dados específicos de uma pessoa física (autônomo, produtor rural, etc.)
type InfoPF struct {
	// NmPF é o nome da pessoa física (1-180 chars)
	NmPF string `xml:"nmPF"`

	// DtNascto é a data de nascimento (YYYY-MM-DD)
	DtNascto types.DataApenas `xml:"dtNascto"`

	// DsCategEmpregador é a categoria do empregador (Tabela 07: 01-10)
	DsCategEmpregador string `xml:"dsCategEmpregador"`

	// IndConstrutora indica se é construtor autônomo (opcional, 1=Sim, 2=Não)
	IndConstrutora types.SimNao `xml:"indConstrutora,omitempty"`

	// IndDesFolha indica se é desonerado (opcional, 1=Sim, 2=Não)
	IndDesFolha types.SimNao `xml:"indDesFolha,omitempty"`
}

// ============================================================================
// InfoComplementar - Informações Complementares
// ============================================================================

// InfoComplementar contém informações complementares do empregador (opcional)
type InfoComplementar struct {
	// IndAprendiz indica se a empresa contrata aprendiz (opcional, 1=Sim, 2=Não)
	IndAprendiz types.SimNao `xml:"indAprendiz,omitempty"`

	// NrRegPJ é o número de registro no CRJ (opcional)
	NrRegPJ string `xml:"nrRegPJ,omitempty"`

	// outras informações complementares podem ser adicionadas aqui
}
