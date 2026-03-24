package application

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/avelinobego/esocial/internal/shared/esocial"
)

// BuildEstabelecimentoTableXML builds XML for establishment table event (S-1020).
// This event registers an establishment (branch/subsidiary) of the company.
//
// Returns the XML string with proper encoding header
func (s *XMLBuilderService) BuildEstabelecimentoTableXML(params *EstabelecimentoTableParams) (string, error) {
	event := &esocial.Evttabestab{
		XMLName: xml.Name{Local: "evtTabEstab"},
		Id:      s.generateID(),
		Ideevento: esocial.Ideevento{
			Tpamb:   params.TpAmb,
			Procemi: params.ProcEmi,
			Verproc: params.VerProc,
		},
		Ideempregador: esocial.Ideempregador{
			Tpinsc: params.TpInsc,
			Nrinsc: params.EmpresaCNPJ,
		},
	}

	xmlData, err := xml.MarshalIndent(event, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal XML: %w", err)
	}

	return xml.Header + string(xmlData), nil
}

// BuildRubricaTableXML builds XML for wage/rubric table event (S-1030).
// This event registers wage types and rubrics for the company.
//
// Returns the XML string with proper encoding header
func (s *XMLBuilderService) BuildRubricaTableXML(params *RubricaTableParams) (string, error) {
	event := &esocial.Evttabrubrica{
		XMLName: xml.Name{Local: "evtTabRubrica"},
		Id:      s.generateID(),
		Ideevento: esocial.Ideevento{
			Tpamb:   params.TpAmb,
			Procemi: params.ProcEmi,
			Verproc: params.VerProc,
		},
		Ideempregador: esocial.Ideempregador{
			Tpinsc: params.TpInsc,
			Nrinsc: params.EmpresaCNPJ,
		},
	}

	xmlData, err := xml.MarshalIndent(event, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal XML: %w", err)
	}

	return xml.Header + string(xmlData), nil
}

// BuildLotacaoTableXML builds XML for work location table event (S-1040).
// This event registers work locations and job positions.
//
// Returns the XML string with proper encoding header
func (s *XMLBuilderService) BuildLotacaoTableXML(params *LotacaoTableParams) (string, error) {
	event := &esocial.Evttablotacao{
		XMLName: xml.Name{Local: "evtTabLotacao"},
		Id:      s.generateID(),
		Ideevento: esocial.Ideevento{
			Tpamb:   params.TpAmb,
			Procemi: params.ProcEmi,
			Verproc: params.VerProc,
		},
		Ideempregador: esocial.Ideempregador{
			Tpinsc: params.TpInsc,
			Nrinsc: params.EmpresaCNPJ,
		},
	}

	xmlData, err := xml.MarshalIndent(event, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal XML: %w", err)
	}

	return xml.Header + string(xmlData), nil
}

// BuildProcessTableXML builds XML for legal process table event (S-1050).
// This event registers legal processes affecting the company/employees.
//
// Returns the XML string with proper encoding header
func (s *XMLBuilderService) BuildProcessTableXML(params *ProcessTableParams) (string, error) {
	event := &esocial.Evttabprocesso{
		XMLName: xml.Name{Local: "evtTabProcesso"},
		Id:      s.generateID(),
		Ideevento: esocial.Ideevento{
			Tpamb:   params.TpAmb,
			Procemi: params.ProcEmi,
			Verproc: params.VerProc,
		},
		Ideempregador: esocial.Ideempregador{
			Tpinsc: params.TpInsc,
			Nrinsc: params.EmpresaCNPJ,
		},
	}

	xmlData, err := xml.MarshalIndent(event, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal XML: %w", err)
	}

	return xml.Header + string(xmlData), nil
}

// BuildCargoTableXML builds XML for job position table event (S-1060).
// This event registers job positions and career levels.
//
// Returns the XML string with proper encoding header
func (s *XMLBuilderService) BuildCargoTableXML(params *CargoTableParams) (string, error) {
	event := &esocial.Evttabprocesso{ // Reusing similar structure
		XMLName: xml.Name{Local: "evtTabCargo"},
		Id:      s.generateID(),
		Ideevento: esocial.Ideevento{
			Tpamb:   params.TpAmb,
			Procemi: params.ProcEmi,
			Verproc: params.VerProc,
		},
		Ideempregador: esocial.Ideempregador{
			Tpinsc: params.TpInsc,
			Nrinsc: params.EmpresaCNPJ,
		},
	}

	xmlData, err := xml.MarshalIndent(event, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal XML: %w", err)
	}

	return xml.Header + string(xmlData), nil
}

// BuildToxicAgentTableXML builds XML for toxic agent table event (S-1070).
// This event registers toxic/hazardous agents present in the workplace.
//
// Returns the XML string with proper encoding header
func (s *XMLBuilderService) BuildToxicAgentTableXML(params *ToxicAgentTableParams) (string, error) {
	event := &esocial.Evttoxic{
		XMLName: xml.Name{Local: "evtToxic"},
		Id:      s.generateID(),
		Ideevento: esocial.Ideevento{
			Tpamb:   params.TpAmb,
			Procemi: params.ProcEmi,
			Verproc: params.VerProc,
		},
		Ideempregador: esocial.Ideempregador{
			Tpinsc: params.TpInsc,
			Nrinsc: params.EmpresaCNPJ,
		},
	}

	xmlData, err := xml.MarshalIndent(event, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal XML: %w", err)
	}

	return xml.Header + string(xmlData), nil
}

// BuildCompleteInfoEmpresaXML builds XML for complete company information event (S-1010).
// This event provides additional details about the company.
//
// Returns the XML string with proper encoding header
func (s *XMLBuilderService) BuildCompleteInfoEmpresaXML(params *CompleteInfoEmpresaParams) (string, error) {
	event := &esocial.Evtinfocomplper{
		XMLName: xml.Name{Local: "evtInfoComplPer"},
		Id:      s.generateID(),
		Ideevento: esocial.Ideevento{
			Tpamb:   params.TpAmb,
			Procemi: params.ProcEmi,
			Verproc: params.VerProc,
		},
		Ideempregador: esocial.Ideempregador{
			Tpinsc: params.TpInsc,
			Nrinsc: params.EmpresaCNPJ,
		},
	}

	xmlData, err := xml.MarshalIndent(event, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal XML: %w", err)
	}

	return xml.Header + string(xmlData), nil
}

// generateID creates a unique ID for eSocial events
// In production, this should be generated according to eSocial specifications
func (s *XMLBuilderService) generateID() string {
	// TODO: Implement proper ID generation following eSocial standards
	// Format should be like: ID followed by digits
	return "ID" + strconv.FormatInt(int64(1000000000+999999999), 10)
}
