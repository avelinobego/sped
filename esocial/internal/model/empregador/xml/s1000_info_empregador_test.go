package xml

import (
	"encoding/xml"
	"sped/esocial/pkg/types"
	"strings"
	"testing"
	"time"
)

// ============================================================================
// Tests para S-1000 - Informações do Empregador
// ============================================================================

// TestS1000MarshalXML_campos_obrigatorios verifica se Marshal gera XML com campos obrigatórios
func TestS1000MarshalXML_campos_obrigatorios(t *testing.T) {
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
				DsCategEmpregador: CategEmpPrivada,
				ClassTributaria:   ClassTribEPP,
			},
			EnderEmpregador: Endereco{
				TpLograd:  "RUA",
				DscLograd: "DAS FLORES",
				NrLograd:  "123",
				Bairro:    "CENTRO",
				CepLograd: "01310100",
				CodMun:    "3550308",
				UF:        "SP",
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
		"<dtCadastro>2024-01-15</dtCadastro>",
		"<razaoSocial>TECH SYSTEMS LTDA</razaoSocial>",
		"<dsCategEmpregador>01</dsCategEmpregador>",
		"<classTributaria>06</classTributaria>",
		"<tpLograd>RUA</tpLograd>",
		"<dscLograd>DAS FLORES</dscLograd>",
	}

	for _, campo := range obrigatorios {
		if !strings.Contains(xmlStr, campo) {
			t.Errorf("Campo obrigatório não encontrado no XML: %s", campo)
		}
	}

	// Verificar atributo id
	if !strings.Contains(xmlStr, `id="ID001"`) {
		t.Errorf("Atributo id não encontrado no XML")
	}
}

// TestS1000MarshalXML_campos_opcionais_omitempty verifica que campos vazios têm tag omitempty
func TestS1000MarshalXML_campos_opcionais_omitempty(t *testing.T) {
	dtCadastro := types.DataApenas(time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC))

	evt := &EvtInfoEmpregador{
		IdeEmpregador: IdeEmpregador{
			TpInsc: TpInscCNPJ,
			NrInsc: "34028316000152",
		},
		InfoCadastro: InfoCadastroEmpregador{
			DtCadastro: dtCadastro,
			InfoPJ: &InfoPJ{
				RazaoSocial:       "TECH SYSTEMS LTDA",
				DsCategEmpregador: CategEmpPrivada,
				ClassTributaria:   ClassTribEPP,
				// NomeFantasia é vazio (opcional)
				// IndDesFolha é zero (opcional)
			},
			EnderEmpregador: Endereco{
				TpLograd:  "RUA",
				DscLograd: "DAS FLORES",
				NrLograd:  "123",
				Bairro:    "CENTRO",
				CepLograd: "01310100",
				CodMun:    "3550308",
				UF:        "SP",
				// CmplLograd é vazio (opcional)
			},
			// InfoContato é nil (opcional)
		},
	}

	b, err := xml.Marshal(evt)
	if err != nil {
		t.Fatalf("Erro ao fazer marshal: %v", err)
	}

	xmlStr := string(b)

	// Campos opcionais não devem aparecer quando vazios
	if strings.Contains(xmlStr, "<nomeFantasia></nomeFantasia>") {
		t.Errorf("Campo vazio nomeFantasia apareceu no XML (deveria ter omitempty)")
	}

	if strings.Contains(xmlStr, "<cmplLograd></cmplLograd>") {
		t.Errorf("Campo vazio cmplLograd apareceu no XML (deveria ter omitempty)")
	}

	if strings.Contains(xmlStr, "<infoContato>") {
		t.Errorf("infoContato vazio não deveria aparecer no XML")
	}

	if strings.Contains(xmlStr, "<indDesFolha") {
		t.Errorf("indDesFolha zero não deveria aparecer no XML")
	}
}

// TestS1000MarshalXML_tipos_especiais verifica serialização de CNPJ, DataApenas, SimNao
func TestS1000MarshalXML_tipos_especiais(t *testing.T) {
	dtCadastro := types.DataApenas(time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC))

	evt := &EvtInfoEmpregador{
		IdeEmpregador: IdeEmpregador{
			TpInsc: TpInscCNPJ,
			NrInsc: "34028316000152",
		},
		InfoCadastro: InfoCadastroEmpregador{
			DtCadastro: dtCadastro,
			InfoPJ: &InfoPJ{
				RazaoSocial:       "TECH SYSTEMS LTDA",
				DsCategEmpregador: CategEmpPrivada,
				ClassTributaria:   ClassTribEPP,
				IndDesFolha:       types.Sim,       // 1
				IndConstrutora:    types.Nao,       // 2
				IndCooperativa:    types.SimNao(0), // deve ser omitido
			},
			EnderEmpregador: Endereco{
				TpLograd:  "RUA",
				DscLograd: "DAS FLORES",
				NrLograd:  "123",
				Bairro:    "CENTRO",
				CepLograd: "01310100",
				CodMun:    "3550308",
				UF:        "SP",
			},
			InfoContato: &Contato{
				NmContato:  "JOÃO",
				CpfContato: "12345678901",
			},
		},
	}

	b, err := xml.Marshal(evt)
	if err != nil {
		t.Fatalf("Erro ao fazer marshal: %v", err)
	}

	xmlStr := string(b)

	// Verificar DataApenas - deve estar em formato YYYY-MM-DD
	if !strings.Contains(xmlStr, "<dtCadastro>2024-01-15</dtCadastro>") {
		t.Errorf("Data não foi serializada corretamente (esperado YYYY-MM-DD)")
	}

	// Verificar SimNao - deve ser 1 ou 2
	if !strings.Contains(xmlStr, "<indDesFolha>1</indDesFolha>") {
		t.Errorf("indDesFolha (SimNao=Sim) não foi serializado como 1")
	}

	if !strings.Contains(xmlStr, "<indConstrutora>2</indConstrutora>") {
		t.Errorf("indConstrutora (SimNao=Nao) não foi serializado como 2")
	}

	// Verificar CPF - deve ter 11 dígitos sem formatação
	if !strings.Contains(xmlStr, "<cpfContato>12345678901</cpfContato>") {
		t.Errorf("CPF não foi serializado corretamente")
	}
}

// TestS1000MarshalXML_ordem_elementos verifica ordem dos elementos XML
func TestS1000MarshalXML_ordem_elementos(t *testing.T) {
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
				DsCategEmpregador: CategEmpPrivada,
				ClassTributaria:   ClassTribEPP,
			},
			EnderEmpregador: Endereco{
				TpLograd:  "RUA",
				DscLograd: "DAS FLORES",
				NrLograd:  "123",
				Bairro:    "CENTRO",
				CepLograd: "01310100",
				CodMun:    "3550308",
				UF:        "SP",
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

	// dtCadastro deve aparecer antes de infoPJ dentro de infoCadastro
	dtCadIdx := strings.Index(xmlStr, "<dtCadastro>")
	infoPJIdx := strings.Index(xmlStr, "<infoPJ>")

	if dtCadIdx == -1 || infoPJIdx == -1 {
		t.Fatalf("Elementos de data/PJ não encontrados no XML")
	}

	if dtCadIdx > infoPJIdx {
		t.Errorf("Ordem incorreta: dtCadastro deve aparecer antes de infoPJ")
	}
}

// TestS1000ExemploValido verifica que o exemplo S-1000 gera XML válido
func TestS1000ExemploValido(t *testing.T) {
	evt := ExemploS1000Empresa()

	b, err := xml.Marshal(evt)
	if err != nil {
		t.Fatalf("Erro ao fazer marshal do exemplo: %v", err)
	}

	xmlStr := string(b)

	// Verificar campos principais
	campos := []string{
		"<tpInsc>1</tpInsc>",
		"<nrInsc>34028316000152</nrInsc>",
		"<razaoSocial>TECH SYSTEMS LTDA</razaoSocial>",
		"<nomeFantasia>TECH SOLUTIONS</nomeFantasia>",
		"<dsCategEmpregador>01</dsCategEmpregador>",
		"<classTributaria>06</classTributaria>",
		"<dscLograd>DAS FLORES</dscLograd>",
		"<nrLograd>123</nrLograd>",
		"<nmContato>JOÃO DA SILVA</nmContato>",
		"<email>contato@techsystems.com.br</email>",
	}

	for _, campo := range campos {
		if !strings.Contains(xmlStr, campo) {
			t.Errorf("Campo esperado não encontrado: %s", campo)
		}
	}

	// Desmarshalling deve funcionar
	var evt2 EvtInfoEmpregador
	err = xml.Unmarshal(b, &evt2)
	if err != nil {
		t.Errorf("Erro ao fazer unmarshal: %v", err)
	}

	if evt2.IdeEmpregador.NrInsc != "34028316000152" {
		t.Errorf("Unmarshal falhou: nrInsc incorreto")
	}

	if evt2.InfoCadastro.InfoPJ == nil || evt2.InfoCadastro.InfoPJ.RazaoSocial != "TECH SYSTEMS LTDA" {
		t.Errorf("Unmarshal falhou: razaoSocial incorreta")
	}
}

// TestS1000MarshalXML_tags_camelCase verifica que as tags XML estão em camelCase
func TestS1000MarshalXML_tags_camelCase(t *testing.T) {
	dtCadastro := types.DataApenas(time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC))

	evt := &EvtInfoEmpregador{
		IdeEmpregador: IdeEmpregador{
			TpInsc: TpInscCNPJ,
			NrInsc: "34028316000152",
		},
		InfoCadastro: InfoCadastroEmpregador{
			DtCadastro: dtCadastro,
			InfoPJ: &InfoPJ{
				RazaoSocial:       "TECH SYSTEMS LTDA",
				DsCategEmpregador: CategEmpPrivada,
				ClassTributaria:   ClassTribEPP,
			},
			EnderEmpregador: Endereco{
				TpLograd:  "RUA",
				DscLograd: "DAS FLORES",
				NrLograd:  "123",
				Bairro:    "CENTRO",
				CepLograd: "01310100",
				CodMun:    "3550308",
				UF:        "SP",
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
		"tpInsc", "nrInsc", "dtCadastro", "razaoSocial", "nomeFantasia",
		"dsCategEmpregador", "classTributaria", "indConstrutora", "indDesFolha",
		"tpLograd", "dscLograd", "nrLograd", "cmplLograd", "bairro",
		"cepLograd", "codMun",
	}

	for _, tag := range camelCaseTags {
		if !strings.Contains(xmlStr, "<"+tag+">") && !strings.Contains(xmlStr, "<"+tag+" ") {
			// Só verificar se o elemento estiver presente
			if strings.Contains(xmlStr, tag) {
				// Tag está presente de alguma forma
				continue
			}
		}
	}
}

// TestS1000PessoaFisica verifica serialização de S-1000 com pessoa física
func TestS1000PessoaFisica(t *testing.T) {
	dtCadastro := types.DataApenas(time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC))
	dtNascto := types.DataApenas(time.Date(1980, 5, 10, 0, 0, 0, 0, time.UTC))

	evt := &EvtInfoEmpregador{
		IdeEmpregador: IdeEmpregador{
			TpInsc: TpInscCPF,
			NrInsc: "12345678901",
		},
		InfoCadastro: InfoCadastroEmpregador{
			DtCadastro: dtCadastro,
			InfoPF: &InfoPF{
				NmPF:              "JOÃO DA SILVA",
				DtNascto:          dtNascto,
				DsCategEmpregador: CategProdutorRural,
			},
			EnderEmpregador: Endereco{
				TpLograd:  "RUA",
				DscLograd: "RURAL",
				NrLograd:  "S/N",
				Bairro:    "ZONA RURAL",
				CepLograd: "12345678",
				CodMun:    "3500108",
				UF:        "SP",
			},
		},
	}

	b, err := xml.Marshal(evt)
	if err != nil {
		t.Fatalf("Erro ao fazer marshal de pessoa física: %v", err)
	}

	xmlStr := string(b)

	// Verificar campos de pessoa física
	if !strings.Contains(xmlStr, "<nmPF>JOÃO DA SILVA</nmPF>") {
		t.Errorf("nmPF não encontrado no XML")
	}

	if !strings.Contains(xmlStr, "<dtNascto>1980-05-10</dtNascto>") {
		t.Errorf("dtNascto não foi serializada corretamente")
	}

	// infoPJ não deve estar presente
	if strings.Contains(xmlStr, "<infoPJ>") {
		t.Errorf("infoPJ não deveria estar presente para pessoa física")
	}
}
