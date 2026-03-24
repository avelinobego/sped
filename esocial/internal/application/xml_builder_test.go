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

	xml, err := service.BuildEmpresaInfoXML(emp, nil)
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

	xml, err := service.BuildTrabalhadorAdmissionXML(trab, "12345678000123")
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

	xml, err := service.BuildEstabelecimentoTableXML(estab, "12345678000123")
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

	xml, err := service.BuildRubricaTableXML(rubrica, "12345678000123")
	if err != nil {
		t.Fatalf("Failed to build XML: %v", err)
	}

	if !strings.Contains(xml, "evtTabRubrica") {
		t.Errorf("Expected evtTabRubrica not found in XML")
	}
}

func TestXMLBuilderService_BuildLotacaoTableXML(t *testing.T) {
	service := NewXMLBuilderService()

	lotacao := &empresa.Lotacao{
		Codigo: "0001",
	}

	xml, err := service.BuildLotacaoTableXML(lotacao, "12345678000123")
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

	xml, err := service.BuildWorkerPreliminaryAdmissionXML(trab, "12345678000123")
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

	xml, err := service.BuildWorkerRemunerationXML(trab, "12345678000123", "202501")
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

	xml, err := service.BuildWorkerDismissalXML(trab, "12345678000123", time.Now(), "pedido_demissao")
	if err != nil {
		t.Fatalf("Failed to build XML: %v", err)
	}

	if !strings.Contains(xml, "evtDeslig") {
		t.Errorf("Expected evtDeslig not found in XML")
	}
}

func TestXMLBuilderService_BuildPeriodClosureXML(t *testing.T) {
	service := NewXMLBuilderService()

	xml, err := service.BuildPeriodClosureXML("12345678000123", "202501")
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

	xml, err := service.BuildWorkLocationTableXML(ambiente, "12345678000123")
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

	xml, err := service.BuildWorkScheduleTableXML(horario, "12345678000123")
	if err != nil {
		t.Fatalf("Failed to build XML: %v", err)
	}

	if !strings.Contains(xml, "evtTabHorario") {
		t.Errorf("Expected evtTabHorario not found in XML")
	}
}

func TestXMLBuilderService_BuildCompleteInfoEmpresaXML(t *testing.T) {
	service := NewXMLBuilderService()

	emp := &empresa.Empresa{
		CNPJ:                    "12345678000123",
		RazaoSocial:             "Empresa Completa Ltda",
		ClassificacaoTributaria: "02",
	}

	xml, err := service.BuildCompleteInfoEmpresaXML(emp, "12345678000123")
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
