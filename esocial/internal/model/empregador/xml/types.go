package xml

import "sped/esocial/pkg/types"

// ============================================================================
// Tipos Auxiliares Compartilhados
// ============================================================================

// Endereco representa o endereço de um empregador ou estabelecimento
// Reutilizável entre S-1000 e S-1005
type Endereco struct {
	TpLograd   string `xml:"tpLograd"`             // Tipo de logradouro (RUA, AV, PÇ, etc.)
	DscLograd  string `xml:"dscLograd"`            // Nome do logradouro (1-100 chars)
	NrLograd   string `xml:"nrLograd"`             // Número (1-10 chars)
	CmplLograd string `xml:"cmplLograd,omitempty"` // Complemento (opcional, 0-30 chars)
	Bairro     string `xml:"bairro"`               // Nome do bairro (1-90 chars)
	CepLograd  string `xml:"cepLograd"`            // CEP (8 dígitos sem formatação)
	CodMun     string `xml:"codMun"`               // Código IBGE do município (7 dígitos)
	UF         string `xml:"uf"`                   // Sigla UF (2 chars: SP, RJ, etc.)
}

// Contato representa informações de contato
// Reutilizável entre S-1000 e S-1005
type Contato struct {
	NmContato  string    `xml:"nmContato,omitempty"`  // Nome (opcional, 1-70 chars)
	CpfContato types.CPF `xml:"cpfContato,omitempty"` // CPF (opcional, 11 dígitos)
	Telefone   string    `xml:"telefone,omitempty"`   // Telefone (opcional, 6-15 chars)
	Email      string    `xml:"email,omitempty"`      // E-mail (opcional, 0-120 chars)
}

// InfoSoftHouse representa informações sobre a software house que mantém o sistema
// Utilizado em S-1005 (opcional)
type InfoSoftHouse struct {
	CnpjSoftHouse   types.CNPJ `xml:"cnpjSoftHouse"`   // CNPJ da software house (14 dígitos)
	NmSoftHouse     string     `xml:"nmSoftHouse"`     // Nome (1-80 chars)
	VersaoSoftHouse string     `xml:"versaoSoftHouse"` // Versão (1-20 chars)
}

// ============================================================================
// Constantes - Tipo de Inscrição
// ============================================================================

const (
	TpInscCNPJ  = 1 // CNPJ
	TpInscCPF   = 2 // CPF
	TpInscCAEPF = 3 // CAEPF
	TpInscCNO   = 4 // CNO
)

// ============================================================================
// Constantes - Tipo de Inscrição do Estabelecimento
// ============================================================================

const (
	TpInscEstabCNPJ = 1 // CNPJ
	TpInscEstabCEI  = 2 // CEI
)

// ============================================================================
// Constantes - Categoria do Empregador (Tabela 07 eSocial)
// ============================================================================

const (
	CategEmpPrivada            = "01" // Empresa privada
	CategAdmFederal            = "02" // Administração Pública Federal
	CategAdmEstadual           = "03" // Administração Pública Estadual
	CategAdmMunicipal          = "04" // Administração Pública Municipal
	CategOrgaoPodJudiciario    = "05" // Órgão do Poder Judiciário
	CategOrgaoPodLegislativo   = "06" // Órgão do Poder Legislativo
	CategSindicatTrabalhadores = "07" // Sindicato de Trabalhadores
	CategSindicatEmpregadores  = "08" // Sindicato de Empregadores
	CategEntidadeSemFinsLucro  = "09" // Entidade sem fins lucrativos
	CategProdutorRural         = "10" // Produtor rural
)

// ============================================================================
// Constantes - Classificação Tributária (Tabela 52 eSocial)
// ============================================================================

const (
	ClassTribMEI               = "01" // Microempreendedor Individual
	ClassTribEconomyMista      = "02" // Economia Mista
	ClassTribSocieNomeColetivo = "03" // Sociedade Nome Coletivo
	ClassTribSocieComandita    = "04" // Sociedade Comandita
	ClassTribSocieAnonima      = "05" // Sociedade Anônima
	ClassTribEPP               = "06" // Empresa de Pequeno Porte
	ClassTribPessoaFisica      = "07" // Pessoa Física
	ClassTribLucroReal         = "08" // Lucro Real
	ClassTribLucroPresumido    = "09" // Lucro Presumido
	ClassTribLucroArbitrado    = "10" // Lucro Arbitrado
	ClassTribSimplesNacional   = "11" // Simples Nacional
)
