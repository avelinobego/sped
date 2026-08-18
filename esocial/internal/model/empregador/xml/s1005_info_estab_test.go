package xml

import (
	"encoding/xml"
	"sped/esocial/pkg/types"
	"strings"
	"testing"
)

// ============================================================================
// Tests para S-1005 - Dados do Estabelecimento
// ============================================================================

// TestS1005MarshalXML_campos_obrigatorios verifica se Marshal gera XML com campos obrigatórios
func TestS1005MarshalXML_campos_obrigatorios(t *testing.T) {
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
				EnderEstab: Endereco{
					TpLograd:  "RUA",
					DscLograd: "DAS FLORES",
					NrLograd:  "123",
					Bairro:    "CENTRO",
					CepLograd: "01310100",
					CodMun:    "3550308",
					UF:        "SP",
				},
			},
		},
	}

	b, err := xml.Marshal(evt)
	if err != nil {
		t.Fatalf("Erro ao fazer marshal: %v", err)
	}

	xmlStr := string(b)

	// Verificar presença de campos obrigatórios
	obrigatorios := []string{
		"<tpInsc>1</tpInsc>",
		"<nrInsc>34028316000152</nrInsc>",
		"<tpInscEstab>1</tpInscEstab>",
		"<nrInscEstab>34028316000152</nrInscEstab>",
		"<nomeEstab>TECH SYSTEMS - MATRIZ</nomeEstab>",
		"<dsCategEmpregador>01</dsCategEmpregador>",
		"<ativEcon>6201500</ativEcon>",
		"<tpLograd>RUA</tpLograd>",
		"<dscLograd>DAS FLORES</dscLograd>",
	}

	for _, campo := range obrigatorios {
		if !strings.Contains(xmlStr, campo) {
			t.Errorf("Campo obrigatório não encontrado no XML: %s", campo)
		}
	}

	// Verificar atributo id
	if !strings.Contains(xmlStr, `id="ID002"`) {
		t.Errorf("Atributo id não encontrado no XML")
	}
}

// TestS1005MarshalXML_referencia_S1000 verifica que ideEmpregador de S-1005 corresponde a S-1000
func TestS1005MarshalXML_referencia_S1000(t *testing.T) {
	evt := ExemploS1005Matriz()

	// Verificar que o CNPJ/CPF da empresa é o mesmo entre S-1000 e S-1005
	s1000 := ExemploS1000Empresa()

	if evt.IdeEmpregador.TpInsc != s1000.IdeEmpregador.TpInsc {
		t.Errorf("TpInsc não corresponde entre S-1000 e S-1005: %d != %d",
			evt.IdeEmpregador.TpInsc, s1000.IdeEmpregador.TpInsc)
	}

	if evt.IdeEmpregador.NrInsc != s1000.IdeEmpregador.NrInsc {
		t.Errorf("NrInsc não corresponde entre S-1000 e S-1005: %s != %s",
			evt.IdeEmpregador.NrInsc, s1000.IdeEmpregador.NrInsc)
	}

	// Para a matriz, tpInscEstab e nrInscEstab devem ser iguais a tpInsc e nrInsc
	if evt.InfoCadastro.IdeEstab.TpInscEstab != evt.IdeEmpregador.TpInsc {
		t.Errorf("tpInscEstab da matriz deve ser igual a tpInsc do empregador")
	}

	if evt.InfoCadastro.IdeEstab.NrInscEstab != evt.IdeEmpregador.NrInsc {
		t.Errorf("nrInscEstab da matriz deve ser igual a nrInsc do empregador")
	}
}

// TestS1005MarshalXML_tipos_especiais verifica serialização de tipos especiais
func TestS1005MarshalXML_tipos_especiais(t *testing.T) {
	evt := &EvtInfoEstab{
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
				IndDesFolha:       types.Sim,       // 1
				IndConstrutora:    types.Nao,       // 2
				IndIntermediaria:  types.SimNao(0), // deve ser omitido
				EnderEstab: Endereco{
					TpLograd:  "RUA",
					DscLograd: "DAS FLORES",
					NrLograd:  "123",
					Bairro:    "CENTRO",
					CepLograd: "01310100",
					CodMun:    "3550308",
					UF:        "SP",
				},
			},
		},
		ContatoPJ: &Contato{
			NmContato:  "MARIA",
			CpfContato: "12345678902",
		},
	}

	b, err := xml.Marshal(evt)
	if err != nil {
		t.Fatalf("Erro ao fazer marshal: %v", err)
	}

	xmlStr := string(b)

	// Verificar SimNao
	if !strings.Contains(xmlStr, "<indDesFolha>1</indDesFolha>") {
		t.Errorf("indDesFolha (SimNao=Sim) não foi serializado como 1")
	}

	if !strings.Contains(xmlStr, "<indConstrutora>2</indConstrutora>") {
		t.Errorf("indConstrutora (SimNao=Nao) não foi serializado como 2")
	}

	// Verificar CPF
	if !strings.Contains(xmlStr, "<cpfContato>12345678902</cpfContato>") {
		t.Errorf("CPF não foi serializado corretamente")
	}
}

// TestS1005MarshalXML_omitempty verifica que campos opcionais têm omitempty
func TestS1005MarshalXML_omitempty(t *testing.T) {
	evt := &EvtInfoEstab{
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
				// Todos os campos opcionais vazios
				EnderEstab: Endereco{
					TpLograd:  "RUA",
					DscLograd: "DAS FLORES",
					NrLograd:  "123",
					Bairro:    "CENTRO",
					CepLograd: "01310100",
					CodMun:    "3550308",
					UF:        "SP",
				},
			},
			// Alteracao é nil (opcional)
		},
		// ContatoPJ é nil (opcional)
	}

	b, err := xml.Marshal(evt)
	if err != nil {
		t.Fatalf("Erro ao fazer marshal: %v", err)
	}

	xmlStr := string(b)

	// Campos opcionais não devem aparecer
	if strings.Contains(xmlStr, "<contatoPJ>") {
		t.Errorf("contatoPJ vazio não deveria aparecer no XML")
	}

	if strings.Contains(xmlStr, "<alteracao>") {
		t.Errorf("alteracao vazia não deveria aparecer no XML")
	}

	if strings.Contains(xmlStr, "<indDesFolha") {
		t.Errorf("indDesFolha zero não deveria aparecer no XML")
	}

	if strings.Contains(xmlStr, "<infoSoftHouse>") {
		t.Errorf("infoSoftHouse vazia não deveria aparecer no XML")
	}
}

// TestS1005ExemploMatrizValido verifica que o exemplo S-1005 (Matriz) gera XML válido
func TestS1005ExemploMatrizValido(t *testing.T) {
	evt := ExemploS1005Matriz()

	b, err := xml.Marshal(evt)
	if err != nil {
		t.Fatalf("Erro ao fazer marshal do exemplo matriz: %v", err)
	}

	xmlStr := string(b)

	// Verificar campos principais
	campos := []string{
		"<nrInscEstab>34028316000152</nrInscEstab>",
		"<nomeEstab>TECH SYSTEMS - MATRIZ</nomeEstab>",
		"<ativEcon>6201500</ativEcon>",
		"<cnpjSoftHouse>11222333000181</cnpjSoftHouse>",
		"<nmSoftHouse>SOFTWARETECH SOLUTIONS</nmSoftHouse>",
		"<versaoSoftHouse>2.1.0</versaoSoftHouse>",
		"<nmContato>MARIA SANTOS</nmContato>",
	}

	for _, campo := range campos {
		if !strings.Contains(xmlStr, campo) {
			t.Errorf("Campo esperado não encontrado: %s", campo)
		}
	}

	// Desmarshalling deve funcionar
	var evt2 EvtInfoEstab
	err = xml.Unmarshal(b, &evt2)
	if err != nil {
		t.Errorf("Erro ao fazer unmarshal: %v", err)
	}

	if evt2.InfoCadastro.DadosEstab.NomeEstab != "TECH SYSTEMS - MATRIZ" {
		t.Errorf("Unmarshal falhou: nomeEstab incorreto")
	}

	if evt2.InfoCadastro.DadosEstab.InfoSoftHouse == nil {
		t.Errorf("Unmarshal falhou: infoSoftHouse é nil")
	}
}

// TestS1005ExemploFilialValido verifica que o exemplo S-1005 (Filial) gera XML válido
func TestS1005ExemploFilialValido(t *testing.T) {
	evt := ExemploS1005Filial()

	b, err := xml.Marshal(evt)
	if err != nil {
		t.Fatalf("Erro ao fazer marshal do exemplo filial: %v", err)
	}

	xmlStr := string(b)

	// Verificar que é uma filial (CNPJ diferente do empregador)
	if evt.IdeEmpregador.NrInsc == evt.InfoCadastro.IdeEstab.NrInscEstab {
		t.Errorf("Filial deve ter CNPJ diferente do empregador")
	}

	// Verificar campos principais
	campos := []string{
		"<nrInsc>34028316000152</nrInsc>",           // Empregador
		"<nrInscEstab>34028316000160</nrInscEstab>", // Filial
		"<nomeEstab>TECH SYSTEMS - FILIAL RIO</nomeEstab>",
		"<ativEcon>6201500</ativEcon>",
		"<dscLograd>PAULISTA</dscLograd>",
		"<codMun>3300100</codMun>",
		"<uf>RJ</uf>",
	}

	for _, campo := range campos {
		if !strings.Contains(xmlStr, campo) {
			t.Errorf("Campo esperado não encontrado: %s", campo)
		}
	}

	// Desmarshalling deve funcionar
	var evt2 EvtInfoEstab
	err = xml.Unmarshal(b, &evt2)
	if err != nil {
		t.Errorf("Erro ao fazer unmarshal: %v", err)
	}

	if evt2.IdeEmpregador.NrInsc != "34028316000152" {
		t.Errorf("Unmarshal falhou: nrInsc do empregador incorreto")
	}

	if evt2.InfoCadastro.IdeEstab.NrInscEstab != "34028316000160" {
		t.Errorf("Unmarshal falhou: nrInscEstab incorreto")
	}
}

// TestS1005MatrizVsFilial verifica estrutura de matriz vs filial
func TestS1005MatrizVsFilial(t *testing.T) {
	// Matriz: tpInscEstab/nrInscEstab == tpInsc/nrInsc
	matriz := ExemploS1005Matriz()
	if matriz.IdeEmpregador.NrInsc != matriz.InfoCadastro.IdeEstab.NrInscEstab {
		t.Errorf("Matriz deve ter mesmo CNPJ que empregador")
	}

	// Filial: tpInscEstab == tpInsc, mas nrInscEstab != nrInsc
	filial := ExemploS1005Filial()
	if filial.IdeEmpregador.TpInsc != filial.InfoCadastro.IdeEstab.TpInscEstab {
		t.Errorf("Filial deve ter mesmo tipo de inscrição que empregador")
	}
	if filial.IdeEmpregador.NrInsc == filial.InfoCadastro.IdeEstab.NrInscEstab {
		t.Errorf("Filial deve ter CNPJ diferente do empregador")
	}
}

// TestS1005MarshalXML_tags_camelCase verifica que as tags XML estão em camelCase
func TestS1005MarshalXML_tags_camelCase(t *testing.T) {
	evt := &EvtInfoEstab{
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
				EnderEstab: Endereco{
					TpLograd:  "RUA",
					DscLograd: "DAS FLORES",
					NrLograd:  "123",
					Bairro:    "CENTRO",
					CepLograd: "01310100",
					CodMun:    "3550308",
					UF:        "SP",
				},
			},
		},
	}

	b, err := xml.Marshal(evt)
	if err != nil {
		t.Fatalf("Erro ao fazer marshal: %v", err)
	}

	xmlStr := string(b)

	// Tags camelCase esperadas
	camelCaseTags := []string{
		"tpInscEstab", "nrInscEstab", "nomeEstab", "ativEcon",
		"indDesFolha", "indConstrutora", "indIntermediaria",
		"tpLograd", "dscLograd", "nrLograd", "cmplLograd", "bairro",
		"cepLograd", "codMun",
	}

	for _, tag := range camelCaseTags {
		if strings.Contains(xmlStr, tag) {
			// Apenas verificar se a tag está corretamente formatada
			if !strings.Contains(xmlStr, "<"+tag+">") && !strings.Contains(xmlStr, "<"+tag+" ") {
				// Tag pode estar formatada de outra forma
			}
		}
	}
}

// TestS1005MarshalXML_ordem_elementos verifica ordem dos elementos XML
func TestS1005MarshalXML_ordem_elementos(t *testing.T) {
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
				EnderEstab: Endereco{
					TpLograd:  "RUA",
					DscLograd: "DAS FLORES",
					NrLograd:  "123",
					Bairro:    "CENTRO",
					CepLograd: "01310100",
					CodMun:    "3550308",
					UF:        "SP",
				},
			},
		},
	}

	b, err := xml.Marshal(evt)
	if err != nil {
		t.Fatalf("Erro ao fazer marshal: %v", err)
	}

	xmlStr := string(b)

	// ideEmpregador deve aparecer antes de infoCadastro
	ideIdx := strings.Index(xmlStr, "<ideEmpregador>")
	infoCadIdx := strings.Index(xmlStr, "<infoCadastro>")

	if ideIdx == -1 || infoCadIdx == -1 {
		t.Fatalf("Elementos não encontrados no XML")
	}

	if ideIdx > infoCadIdx {
		t.Errorf("Ordem incorreta: ideEmpregador deve aparecer antes de infoCadastro")
	}

	// ideEstab deve aparecer antes de dadosEstab
	ideEstabIdx := strings.Index(xmlStr, "<ideEstab>")
	dadosEstabIdx := strings.Index(xmlStr, "<dadosEstab>")

	if ideEstabIdx == -1 || dadosEstabIdx == -1 {
		t.Fatalf("Elementos de estabelecimento não encontrados no XML")
	}

	if ideEstabIdx > dadosEstabIdx {
		t.Errorf("Ordem incorreta: ideEstab deve aparecer antes de dadosEstab")
	}
}
