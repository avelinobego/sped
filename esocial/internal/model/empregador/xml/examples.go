package xml

import (
	"sped/esocial/pkg/types"
	"time"
)

// ============================================================================
// Exemplo S-1000: Informações do Empregador
// ============================================================================

// ExemploS1000Empresa cria um exemplo válido de evento S-1000 para uma empresa privada
//
// Dados da empresa:
// - Nome: TECH SYSTEMS LTDA
// - CNPJ: 34028316000152
// - Categoria: 01 (Empresa privada)
// - Classificação: 06 (EPP)
// - Endereço: Rua das Flores, 123, Centro, São Paulo, SP
// - CEP: 01310100, Município: 3550308 (São Paulo)
// - Contato: João da Silva, CPF, telefone e e-mail
func ExemploS1000Empresa() *EvtInfoEmpregador {
	dtCadastro := types.DataApenas(time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC))

	evt := &EvtInfoEmpregador{
		Id: "ID001",
		IdeEmpregador: IdeEmpregador{
			TpInsc: TpInscCNPJ,
			NrInsc: "34028316000152",
		},
		InfoCadastro: InfoCadastroEmpregador{
			DtCadastro: dtCadastro,
			InfoPJ: &InfoPJ{
				RazaoSocial:       "TECH SYSTEMS LTDA",
				NomeFantasia:      "TECH SOLUTIONS",
				DsCategEmpregador: CategEmpPrivada,
				ClassTributaria:   ClassTribEPP,
				IndDesFolha:       types.Nao,
				IndConstrutora:    types.Nao,
				IndCooperativa:    types.Nao,
			},
			EnderEmpregador: Endereco{
				TpLograd:   "RUA",
				DscLograd:  "DAS FLORES",
				NrLograd:   "123",
				CmplLograd: "APTO 45",
				Bairro:     "CENTRO",
				CepLograd:  "01310100",
				CodMun:     "3550308",
				UF:         "SP",
			},
			InfoContato: &Contato{
				NmContato:  "JOÃO DA SILVA",
				CpfContato: "12345678901",
				Telefone:   "1133334444",
				Email:      "contato@techsystems.com.br",
			},
		},
	}

	return evt
}

// ============================================================================
// Exemplo S-1005: Dados do Estabelecimento - Matriz
// ============================================================================

// ExemploS1005Matriz cria um exemplo válido de evento S-1005 para a matriz
//
// Dados:
// - Empregador: CNPJ 34028316000152 (referência a S-1000)
// - Estabelecimento: CNPJ 34028316000152 (mesmo que matriz)
// - Nome: TECH SYSTEMS - MATRIZ
// - CNAE: 6201500 (Desenvolvimento de software customizado)
// - Endereço: Rua das Flores, 123, Centro, São Paulo, SP
// - Software House: SOFTWARETECH SOLUTIONS, versão 2.1.0
func ExemploS1005Matriz() *EvtInfoEstab {
	evt := &EvtInfoEstab{
		Id: "ID002",
		IdeEmpregador: IdeEmpregador{
			TpInsc: TpInscCNPJ,
			NrInsc: "34028316000152",
		},
		InfoCadastro: InfoCadastroEstab{
			IdeEstab: IdeEstab{
				TpInscEstab: TpInscEstabCNPJ,
				NrInscEstab: "34028316000152",
			},
			DadosEstab: DadosEstab{
				NomeEstab:         "TECH SYSTEMS - MATRIZ",
				DsCategEmpregador: CategEmpPrivada,
				AtivEcon:          "6201500",
				IndDesFolha:       types.Nao,
				IndConstrutora:    types.Nao,
				EnderEstab: Endereco{
					TpLograd:   "RUA",
					DscLograd:  "DAS FLORES",
					NrLograd:   "123",
					CmplLograd: "APTO 45",
					Bairro:     "CENTRO",
					CepLograd:  "01310100",
					CodMun:     "3550308",
					UF:         "SP",
				},
				InfoSoftHouse: &InfoSoftHouse{
					CnpjSoftHouse:   "11222333000181",
					NmSoftHouse:     "SOFTWARETECH SOLUTIONS",
					VersaoSoftHouse: "2.1.0",
				},
			},
		},
		ContatoPJ: &Contato{
			NmContato:  "MARIA SANTOS",
			CpfContato: "12345678902",
			Telefone:   "1133334444",
			Email:      "maria@techsystems.com.br",
		},
	}

	return evt
}

// ============================================================================
// Exemplo S-1005: Dados do Estabelecimento - Filial
// ============================================================================

// ExemploS1005Filial cria um exemplo válido de evento S-1005 para uma filial
//
// Dados:
// - Empregador: CNPJ 34028316000152 (referência a S-1000)
// - Estabelecimento: CNPJ 34028316000160 (CNPJ da filial)
// - Nome: TECH SYSTEMS - FILIAL RIO
// - CNAE: 6201500 (Desenvolvimento de software customizado)
// - Endereço: Avenida Paulista, 1000, Bela Vista, Rio de Janeiro, RJ
// - CEP: 01311100, Município: 3300100 (Rio de Janeiro)
func ExemploS1005Filial() *EvtInfoEstab {
	evt := &EvtInfoEstab{
		Id: "ID003",
		IdeEmpregador: IdeEmpregador{
			TpInsc: TpInscCNPJ,
			NrInsc: "34028316000152",
		},
		InfoCadastro: InfoCadastroEstab{
			IdeEstab: IdeEstab{
				TpInscEstab: TpInscEstabCNPJ,
				NrInscEstab: "34028316000160",
			},
			DadosEstab: DadosEstab{
				NomeEstab:         "TECH SYSTEMS - FILIAL RIO",
				DsCategEmpregador: CategEmpPrivada,
				AtivEcon:          "6201500",
				IndDesFolha:       types.Nao,
				IndConstrutora:    types.Nao,
				EnderEstab: Endereco{
					TpLograd:  "AVENIDA",
					DscLograd: "PAULISTA",
					NrLograd:  "1000",
					Bairro:    "BELA VISTA",
					CepLograd: "01311100",
					CodMun:    "3300100",
					UF:        "RJ",
				},
			},
		},
	}

	return evt
}
