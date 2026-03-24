package application

import (
	"strings"
	"testing"
	"time"

	"github.com/avelinobego/esocial/internal/domain/ambiente_trabalho"
	"github.com/avelinobego/esocial/internal/domain/empresa"
	"github.com/avelinobego/esocial/internal/domain/trabalhador"
)

func TestXMLBuilderService_BuildEmpresaInfoXML(t *testing.T) {
	service := NewXMLBuilderService()

	emp := &empresa.Empresa{
		CNPJ:                    "12345678000123",
		RazaoSocial:             "Empresa Teste Ltda",
		ClassificacaoTributaria: "01",
		IndCooperativa:          "N",
		IndConstrutora:          "N",
		IndDesoneracao:          "N",
	}

	params := &EventParams{
		EmpresaCNPJ: "12345678000123",
		TpAmb:       1,
		ProcEmi:     1,
		VerProc:     "1.0.0",
		TpInsc:      1,
	}

	xml, err := service.BuildEmpresaInfoXML(emp, params)
	if err != nil {
		t.Fatalf("Failed to build XML: %v", err)
	}

	// Check if XML contains expected elements
	expectedElements := []string{
		"evtInfoEmpregador",
		"ideEmpregador",
		"12345678000123",
		"classTrib",
		"01",
	}

	for _, elem := range expectedElements {
		if !strings.Contains(xml, elem) {
			t.Errorf("Expected element %s not found in XML", elem)
		}
	}

	t.Logf("Generated XML:\n%s", xml)
}

func TestXMLBuilderService_BuildTrabalhadorAdmissionXML(t *testing.T) {
	service := NewXMLBuilderService()

	trab := &trabalhador.Trabalhador{
		CPF:               "12345678901",
		Nome:              "João Silva",
		Sexo:              "M",
		DataNascimento:    time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
		PaisNascimento:    "105", // Brazil
		PaisNacionalidade: "105",
	}

	params := &EventParams{
		EmpresaCNPJ: "12345678000123",
		TpAmb:       1,
		ProcEmi:     1,
		VerProc:     "1.0.0",
		TpInsc:      1,
	}
	xml, err := service.BuildTrabalhadorAdmissionXML(trab, params)
	if err != nil {
		t.Fatalf("Failed to build XML: %v", err)
	}

	// Check if XML contains expected elements
	expectedElements := []string{
		"evtAdmissao",
		"trabalhador",
		"12345678901",
		"João Silva",
		"1990-01-01",
	}

	for _, elem := range expectedElements {
		if !strings.Contains(xml, elem) {
			t.Errorf("Expected element %s not found in XML", elem)
		}
	}

	t.Logf("Generated XML:\n%s", xml)
}

// Test empresa domain builders

func TestXMLBuilderService_BuildEstabelecimentoTableXML(t *testing.T) {
	service := NewXMLBuilderService()

	estab := &empresa.Estabelecimento{
		TipoInscricao:   "1",
		NumeroInscricao: "12345678000456",
	}

	params := &EstabelecimentoTableParams{
		EventParams: EventParams{
			EmpresaCNPJ: "12345678000123",
			TpAmb:       1,
			ProcEmi:     1,
			VerProc:     "1.0.0",
			TpInsc:      1,
		},
		Estabelecimento: estab,
	}
	xml, err := service.BuildEstabelecimentoTableXML(params)
	if err != nil {
		t.Fatalf("Failed to build XML: %v", err)
	}

	if !strings.Contains(xml, "evtTabEstab") {
		t.Errorf("Expected evtTabEstab not found in XML")
	}
}

func TestXMLBuilderService_BuildRubricaTableXML(t *testing.T) {
	service := NewXMLBuilderService()

	rubrica := &empresa.Rubrica{
		Codigo: "001",
	}

	params := &RubricaTableParams{
		EventParams: EventParams{
			EmpresaCNPJ: "12345678000123",
			TpAmb:       1,
			ProcEmi:     1,
			VerProc:     "1.0.0",
			TpInsc:      1,
		},
		Rubrica: rubrica,
	}
	xml, err := service.BuildRubricaTableXML(params)
	if err != nil {
		t.Fatalf("Failed to build XML: %v", err)
	}

	if !strings.Contains(xml, "evtTabRubrica") {
		t.Errorf("Expected evtTabRubrica not found in XML")
	}

	t.Logf("Generated XML:\n%s", xml)
}

func TestXMLBuilderService_BuildLotacaoTableXML(t *testing.T) {
	service := NewXMLBuilderService()

	lotacao := &empresa.Lotacao{
		Codigo: "0001",
	}

	params := &LotacaoTableParams{
		EventParams: EventParams{
			EmpresaCNPJ: "12345678000123",
			TpAmb:       1,
			ProcEmi:     1,
			VerProc:     "1.0.0",
			TpInsc:      1,
		},
		Lotacao: lotacao,
	}
	xml, err := service.BuildLotacaoTableXML(params)
	if err != nil {
		t.Fatalf("Failed to build XML: %v", err)
	}

	if !strings.Contains(xml, "evtTabLotacao") {
		t.Errorf("Expected evtTabLotacao not found in XML")
	}
}

// Test trabalhador domain builders

func TestXMLBuilderService_BuildWorkerPreliminaryAdmissionXML(t *testing.T) {
	service := NewXMLBuilderService()

	trab := &trabalhador.Trabalhador{
		CPF:  "12345678901",
		Nome: "Maria Silva",
	}

	params := &WorkerPreliminaryAdmissionParams{
		EventParams: EventParams{
			EmpresaCNPJ: "12345678000123",
			TpAmb:       1,
			ProcEmi:     1,
			VerProc:     "1.0.0",
			TpInsc:      1,
		},
		Worker: trab,
	}
	xml, err := service.BuildWorkerPreliminaryAdmissionXML(params)
	if err != nil {
		t.Fatalf("Failed to build XML: %v", err)
	}

	if !strings.Contains(xml, "evtAdmPrelim") {
		t.Errorf("Expected evtAdmPrelim not found in XML")
	}
}

func TestXMLBuilderService_BuildWorkerRemunerationXML(t *testing.T) {
	service := NewXMLBuilderService()

	trab := &trabalhador.Trabalhador{
		CPF:  "12345678901",
		Nome: "Carlos Santos",
	}

	params := &WorkerRemunerationParams{
		EventParams: EventParams{
			EmpresaCNPJ: "12345678000123",
			TpAmb:       1,
			ProcEmi:     1,
			VerProc:     "1.0.0",
			TpInsc:      1,
		},
		Worker:  trab,
		Periodo: "202501",
	}
	xml, err := service.BuildWorkerRemunerationXML(params)
	if err != nil {
		t.Fatalf("Failed to build XML: %v", err)
	}

	if !strings.Contains(xml, "evtRemun") {
		t.Errorf("Expected evtRemun not found in XML")
	}
}

func TestXMLBuilderService_BuildWorkerDismissalXML(t *testing.T) {
	service := NewXMLBuilderService()

	trab := &trabalhador.Trabalhador{
		CPF:  "12345678901",
		Nome: "Pedro Costa",
	}

	params := &WorkerDismissalParams{
		EventParams: EventParams{
			EmpresaCNPJ: "12345678000123",
			TpAmb:       1,
			ProcEmi:     1,
			VerProc:     "1.0.0",
			TpInsc:      1,
		},
		Worker:             trab,
		DtDeslig:           time.Now(),
		MotivoDesligamento: "pedido_demissao",
	}
	xml, err := service.BuildWorkerDismissalXML(params)
	if err != nil {
		t.Fatalf("Failed to build XML: %v", err)
	}

	if !strings.Contains(xml, "evtDeslig") {
		t.Errorf("Expected evtDeslig not found in XML")
	}
}

func TestXMLBuilderService_BuildPeriodClosureXML(t *testing.T) {
	service := NewXMLBuilderService()

	params := &PeriodClosureParams{
		EventParams: EventParams{
			EmpresaCNPJ: "12345678000123",
			TpAmb:       1,
			ProcEmi:     1,
			VerProc:     "1.0.0",
			TpInsc:      1,
		},
		Periodo: "202501",
	}
	xml, err := service.BuildPeriodClosureXML(params)
	if err != nil {
		t.Fatalf("Failed to build XML: %v", err)
	}

	if !strings.Contains(xml, "evtFechaEvPer") {
		t.Errorf("Expected evtFechaEvPer not found in XML")
	}
}

// Test ambiente_trabalho domain builders

func TestXMLBuilderService_BuildWorkLocationTableXML(t *testing.T) {
	service := NewXMLBuilderService()

	ambiente := &ambiente_trabalho.AmbienteTrabalho{
		Codigo:    "LOC001",
		Descricao: "Escritório Principal",
	}

	params := &WorkLocationTableParams{
		EventParams: EventParams{
			EmpresaCNPJ: "12345678000123",
			TpAmb:       1,
			ProcEmi:     1,
			VerProc:     "1.0.0",
			TpInsc:      1,
		},
		Ambiente: ambiente,
	}
	xml, err := service.BuildWorkLocationTableXML(params)
	if err != nil {
		t.Fatalf("Failed to build XML: %v", err)
	}

	if !strings.Contains(xml, "evtTabLotacao") {
		t.Errorf("Expected evtTabLotacao not found in XML")
	}
}

func TestXMLBuilderService_BuildWorkScheduleTableXML(t *testing.T) {
	service := NewXMLBuilderService()

	horario := &ambiente_trabalho.Horario{
		Codigo:      "HORARIO001",
		Descricao:   "Jornada Padrão",
		TipoJornada: "normal",
	}

	params := &WorkScheduleTableParams{
		EventParams: EventParams{
			EmpresaCNPJ: "12345678000123",
			TpAmb:       1,
			ProcEmi:     1,
			VerProc:     "1.0.0",
			TpInsc:      1,
		},
		Horario: horario,
	}
	xml, err := service.BuildWorkScheduleTableXML(params)
	if err != nil {
		t.Fatalf("Failed to build XML: %v", err)
	}

	if !strings.Contains(xml, "evtTabLotacao") {
		t.Errorf("Expected evtTabLotacao not found in XML")
	}
}

func TestXMLBuilderService_BuildCompleteInfoEmpresaXML(t *testing.T) {
	service := NewXMLBuilderService()

	emp := &empresa.Empresa{
		CNPJ:                    "12345678000123",
		RazaoSocial:             "Empresa Completa Ltda",
		ClassificacaoTributaria: "02",
	}

	params := &CompleteInfoEmpresaParams{
		EventParams: EventParams{
			EmpresaCNPJ: "12345678000123",
			TpAmb:       1,
			ProcEmi:     1,
			VerProc:     "1.0.0",
			TpInsc:      1,
		},
		Empresa: emp,
	}
	xml, err := service.BuildCompleteInfoEmpresaXML(params)
	if err != nil {
		t.Fatalf("Failed to build XML: %v", err)
	}

	if !strings.Contains(xml, "evtInfoComplPer") {
		t.Errorf("Expected evtInfoComplPer not found in XML")
	}
}

func TestXMLBuilderService_BuildGenericEventXML(t *testing.T) {
	service := NewXMLBuilderService()

	// Test with a simple struct that has xml tags
	type TestEvent struct {
		Name  string `xml:"name"`
		Value string `xml:"value"`
	}

	event := &TestEvent{
		Name:  "Test",
		Value: "123",
	}

	xml, err := service.BuildGenericEventXML(event)
	if err != nil {
		t.Fatalf("Failed to build XML: %v", err)
	}

	if !strings.Contains(xml, "Test") || !strings.Contains(xml, "123") {
		t.Errorf("Expected test data not found in XML")
	}
}
