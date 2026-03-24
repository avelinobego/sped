package application

import (
	"encoding/xml"
	"fmt"

	"github.com/avelinobego/esocial/internal/shared/esocial"
)

// BuildWorkLocationTableXML builds XML for work location table event (S-1040).
// This event registers the locations where workers are allocated.
//
// Returns the XML string with proper encoding header
func (s *XMLBuilderService) BuildWorkLocationTableXML(params *WorkLocationTableParams) (string, error) {
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

// BuildWorkScheduleTableXML builds XML for work schedule table event.
// This event registers work schedules and work hour patterns.
//
// Returns the XML string with proper encoding header
func (s *XMLBuilderService) BuildWorkScheduleTableXML(params *WorkScheduleTableParams) (string, error) {
	// Using environment table event structure for schedules
	event := &esocial.Evttablotacao{
		XMLName: xml.Name{Local: "evtTabHorario"},
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

// BuildEnvironmentDataXML builds XML for environment/location data.
// This event provides detailed information about work environments.
//
// Returns the XML string with proper encoding header
func (s *XMLBuilderService) BuildEnvironmentDataXML(params *EnvironmentDataParams) (string, error) {
	event := &esocial.Evttablotacao{
		XMLName: xml.Name{Local: params.TipoEvento},
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

// BuildHazardousEnvironmentDataXML builds XML for hazardous/toxic environment data.
// This event registers environments with toxic or hazardous conditions.
//
// Returns the XML string with proper encoding header
func (s *XMLBuilderService) BuildHazardousEnvironmentDataXML(params *HazardousEnvironmentDataParams) (string, error) {
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

// BuildDepartmentTableXML builds XML for department/sector table.
// This event registers company departments and organizational units.
//
// Returns the XML string with proper encoding header
func (s *XMLBuilderService) BuildDepartmentTableXML(params *DepartmentTableParams) (string, error) {
	event := &esocial.Evttablotacao{
		XMLName: xml.Name{Local: "evtTabDept"},
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

// BuildShiftTableXML builds XML for work shift table.
// This event registers shift patterns and rotating schedules.
//
// Returns the XML string with proper encoding header
func (s *XMLBuilderService) BuildShiftTableXML(params *ShiftTableParams) (string, error) {
	event := &esocial.Evttablotacao{
		XMLName: xml.Name{Local: "evtTabTurno"},
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
