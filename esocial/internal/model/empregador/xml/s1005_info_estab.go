package xml

import "sped/esocial/pkg/types"

// ============================================================================
// EvtInfoEstab - S-1005: Dados do Estabelecimento
// ============================================================================

// EvtInfoEstab representa o evento S-1005 (Dados do Estabelecimento)
// Deve ser transmitido após S-1000 para cada filial/estabelecimento da empresa
type EvtInfoEstab struct {
	// Id é um atributo XML que identifica o evento
	Id string `xml:"id,attr,omitempty"`

	// IdeEmpregador contém a identificação do empregador (referência a S-1000)
	IdeEmpregador IdeEmpregador `xml:"ideEmpregador"`

	// InfoCadastro contém os dados cadastrais do estabelecimento
	InfoCadastro InfoCadastroEstab `xml:"infoCadastro"`

	// ContatoPJ contém informações de contato da pessoa jurídica (opcional)
	ContatoPJ *Contato `xml:"contatoPJ,omitempty"`
}

// ============================================================================
// InfoCadastroEstab - Informações Cadastrais do Estabelecimento
// ============================================================================

// InfoCadastroEstab contém os dados cadastrais completos do estabelecimento
type InfoCadastroEstab struct {
	// IdeEstab contém a identificação do estabelecimento
	IdeEstab IdeEstab `xml:"ideEstab"`

	// DadosEstab contém os dados específicos do estabelecimento
	DadosEstab DadosEstab `xml:"dadosEstab"`

	// Alteracao contém informações sobre alterações (opcional)
	Alteracao *Alteracao `xml:"alteracao,omitempty"`
}

// ============================================================================
// IdeEstab - Identificação do Estabelecimento
// ============================================================================

// IdeEstab contém a identificação do estabelecimento
type IdeEstab struct {
	// TpInscEstab é o tipo de inscrição do estabelecimento (1=CNPJ, 2=CEI)
	TpInscEstab int `xml:"tpInscEstab"`

	// NrInscEstab é o número de inscrição do estabelecimento sem formatação
	NrInscEstab string `xml:"nrInscEstab"`
}

// ============================================================================
// DadosEstab - Dados do Estabelecimento
// ============================================================================

// DadosEstab contém os dados específicos do estabelecimento
type DadosEstab struct {
	// NomeEstab é o nome do estabelecimento (1-180 chars)
	NomeEstab string `xml:"nomeEstab"`

	// DsCategEmpregador é a categoria do empregador (Tabela 07: 01-10)
	// Deve ser compatível com S-1000
	DsCategEmpregador string `xml:"dsCategEmpregador"`

	// AtivEcon é a atividade econômica (CNAE - 7 dígitos, ex: 6201500)
	AtivEcon string `xml:"ativEcon"`

	// IndDesFolha indica se o estabelecimento é desonerado (opcional, 1=Sim, 2=Não)
	IndDesFolha types.SimNao `xml:"indDesFolha,omitempty"`

	// IndConstrutora indica se é construtor (opcional, 1=Sim, 2=Não)
	IndConstrutora types.SimNao `xml:"indConstrutora,omitempty"`

	// IndIntermediaria indica se é intermediária (opcional, 1=Sim, 2=Não)
	IndIntermediaria types.SimNao `xml:"indIntermediaria,omitempty"`

	// EnderEstab contém o endereço do estabelecimento
	EnderEstab Endereco `xml:"enderEstab"`

	// InfoSoftHouse contém informações sobre a software house (opcional)
	InfoSoftHouse *InfoSoftHouse `xml:"infoSoftHouse,omitempty"`

	// InfoComplementar contém informações complementares (opcional)
	InfoComplementar *InfoComplementarEstab `xml:"infoComplementar,omitempty"`
}

// ============================================================================
// Alteracao - Informações sobre Alterações
// ============================================================================

// Alteracao contém informações sobre alterações no cadastro do estabelecimento
type Alteracao struct {
	// DtAlteracao é a data da alteração (YYYY-MM-DD)
	DtAlteracao types.DataApenas `xml:"dtAlteracao"`

	// DsAlteracao é a descrição da alteração
	DsAlteracao string `xml:"dsAlteracao,omitempty"`
}

// ============================================================================
// InfoComplementarEstab - Informações Complementares do Estabelecimento
// ============================================================================

// InfoComplementarEstab contém informações complementares do estabelecimento
type InfoComplementarEstab struct {
	// IndAprendiz indica se o estabelecimento contrata aprendiz (opcional, 1=Sim, 2=Não)
	IndAprendiz types.SimNao `xml:"indAprendiz,omitempty"`

	// outras informações complementares podem ser adicionadas aqui
}
