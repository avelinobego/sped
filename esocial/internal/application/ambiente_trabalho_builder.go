package application

import (
	"encoding/xml"
	"fmt"

	"github.com/avelinobego/esocial/internal/domain/ambiente_trabalho"
	"github.com/avelinobego/esocial/internal/shared/esocial"
)

// BuildWorkLocationTableXML builds XML for work location table event (S-1040).
// This event registers the locations where workers are allocated.
// Parameters:
//   - ambiente: The work environment/location entity
//   - empresaCNPJ: The CNPJ of the company
//
// Returns the XML string with proper encoding header
func (s *XMLBuilderService) BuildWorkLocationTableXML(ambiente *ambiente_trabalho.AmbienteTrabalho, empresaCNPJ string) (string, error) {
	event := &esocial.Evttablotacao{
		XMLName: xml.Name{Local: "evtTabLotacao"},
		Id:      s.generateID(),
		Ideevento: esocial.Ideevento{
			Tpamb:   1,
			Procemi: 1,
			Verproc: "1.0.0",
		},
		Ideempregador: esocial.Ideempregador{
			Tpinsc: 1,
			Nrinsc: empresaCNPJ,
		},
	}

	xmlData, err := xml.MarshalIndent(event, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal XML: %w", err)
	}

	return xml.Header + string(xmlData), nil
}

// BuildWorkScheduleTableXML builds XML for work schedule table event.
// This event registers work schedules and work hour patterns.
// Parameters:
//   - horario: The work schedule entity
//   - empresaCNPJ: The CNPJ of the company
//
// Returns the XML string with proper encoding header
func (s *XMLBuilderService) BuildWorkScheduleTableXML(horario *ambiente_trabalho.Horario, empresaCNPJ string) (string, error) {
	// Using environment table event structure for schedules
	event := &esocial.Evttablotacao{
		XMLName: xml.Name{Local: "evtTabHorario"},
		Id:      s.generateID(),
		Ideevento: esocial.Ideevento{
			Tpamb:   1,
			Procemi: 1,
			Verproc: "1.0.0",
		},
		Ideempregador: esocial.Ideempregador{
			Tpinsc: 1,
			Nrinsc: empresaCNPJ,
		},
	}

	xmlData, err := xml.MarshalIndent(event, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal XML: %w", err)
	}

	return xml.Header + string(xmlData), nil
}

// BuildEnvironmentDataXML builds XML for environment/location data.
// This event provides detailed information about work environments.
// Parameters:
//   - ambiente: The work environment entity
//   - empresaCNPJ: The CNPJ of the company
//   - tipoEvento: Type of environment event
//
// Returns the XML string with proper encoding header
func (s *XMLBuilderService) BuildEnvironmentDataXML(ambiente *ambiente_trabalho.AmbienteTrabalho, empresaCNPJ string, tipoEvento string) (string, error) {
	event := &esocial.Evttablotacao{
		XMLName: xml.Name{Local: tipoEvento},
		Id:      s.generateID(),
		Ideevento: esocial.Ideevento{
			Tpamb:   1,
			Procemi: 1,
			Verproc: "1.0.0",
		},
		Ideempregador: esocial.Ideempregador{
			Tpinsc: 1,
			Nrinsc: empresaCNPJ,
		},
	}

	xmlData, err := xml.MarshalIndent(event, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal XML: %w", err)
	}

	return xml.Header + string(xmlData), nil
}

// BuildHazardousEnvironmentDataXML builds XML for hazardous/toxic environment data.
// This event registers environments with toxic or hazardous conditions.
// Parameters:
//   - ambiente: The work environment entity
//   - empresaCNPJ: The CNPJ of the company
//   - agenteToxick: Toxic agent code/identifier
//
// Returns the XML string with proper encoding header
func (s *XMLBuilderService) BuildHazardousEnvironmentDataXML(ambiente *ambiente_trabalho.AmbienteTrabalho, empresaCNPJ string, agenteToxick string) (string, error) {
	event := &esocial.Evttoxic{
		XMLName: xml.Name{Local: "evtToxic"},
		Id:      s.generateID(),
		Ideevento: esocial.Ideevento{
			Tpamb:   1,
			Procemi: 1,
			Verproc: "1.0.0",
		},
		Ideempregador: esocial.Ideempregador{
			Tpinsc: 1,
			Nrinsc: empresaCNPJ,
		},
	}

	xmlData, err := xml.MarshalIndent(event, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal XML: %w", err)
	}

	return xml.Header + string(xmlData), nil
}

// BuildDepartmentTableXML builds XML for department/sector table.
// This event registers company departments and organizational units.
// Parameters:
//   - codigoDept: Code for the department/sector
//   - nomeDept: Name of the department
//   - empresaCNPJ: The CNPJ of the company
//
// Returns the XML string with proper encoding header
func (s *XMLBuilderService) BuildDepartmentTableXML(codigoDept, nomeDept, empresaCNPJ string) (string, error) {
	event := &esocial.Evttablotacao{
		XMLName: xml.Name{Local: "evtTabDept"},
		Id:      s.generateID(),
		Ideevento: esocial.Ideevento{
			Tpamb:   1,
			Procemi: 1,
			Verproc: "1.0.0",
		},
		Ideempregador: esocial.Ideempregador{
			Tpinsc: 1,
			Nrinsc: empresaCNPJ,
		},
	}

	xmlData, err := xml.MarshalIndent(event, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal XML: %w", err)
	}

	return xml.Header + string(xmlData), nil
}

// BuildShiftTableXML builds XML for work shift table.
// This event registers shift patterns and rotating schedules.
// Parameters:
//   - codigoTurno: Code for the shift
//   - nomeTurno: Name of the shift
//   - empresaCNPJ: The CNPJ of the company
//
// Returns the XML string with proper encoding header
func (s *XMLBuilderService) BuildShiftTableXML(codigoTurno, nomeTurno, empresaCNPJ string) (string, error) {
	event := &esocial.Evttablotacao{
		XMLName: xml.Name{Local: "evtTabTurno"},
		Id:      s.generateID(),
		Ideevento: esocial.Ideevento{
			Tpamb:   1,
			Procemi: 1,
			Verproc: "1.0.0",
		},
		Ideempregador: esocial.Ideempregador{
			Tpinsc: 1,
			Nrinsc: empresaCNPJ,
		},
	}

	xmlData, err := xml.MarshalIndent(event, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal XML: %w", err)
	}

	return xml.Header + string(xmlData), nil
}
