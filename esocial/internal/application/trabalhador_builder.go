package application

import (
	"encoding/xml"
	"fmt"

	"github.com/avelinobego/esocial/internal/shared/esocial"
)

// BuildWorkerPreliminaryAdmissionXML builds XML for preliminary worker admission event (S-2210).
// This event registers initial/preliminary information about a new worker admission.
//
// Returns the XML string with proper encoding header
func (s *XMLBuilderService) BuildWorkerPreliminaryAdmissionXML(params *WorkerPreliminaryAdmissionParams) (string, error) {
	event := &esocial.Evtadmprelim{
		XMLName: xml.Name{Local: "evtAdmPrelim"},
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

// BuildWorkerAddressChangeXML builds XML for worker registration change event (S-2230).
// This event records changes to worker registration data (addresses, contact info, etc.).
//
// Returns the XML string with proper encoding header
func (s *XMLBuilderService) BuildWorkerAddressChangeXML(params *WorkerAddressChangeParams) (string, error) {
	event := &esocial.Evtaltcadastral{
		XMLName: xml.Name{Local: "evtAltCadastral"},
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

// BuildWorkerContractChangeXML builds XML for worker contract change event (S-2240).
// This event records changes to employment contract terms and conditions.
//
// Returns the XML string with proper encoding header
func (s *XMLBuilderService) BuildWorkerContractChangeXML(params *WorkerContractChangeParams) (string, error) {
	event := &esocial.Evtaltcontratual{
		XMLName: xml.Name{Local: "evtAltContratual"},
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

// BuildWorkerTemporaryLeaveXML builds XML for worker temporary leave event (S-2250).
// This event records temporary absences from work (sick leave, statutory leave, etc.).
//
// Returns the XML string with proper encoding header
func (s *XMLBuilderService) BuildWorkerTemporaryLeaveXML(params *WorkerTemporaryLeaveParams) (string, error) {
	event := &esocial.Evtafasttemp{
		XMLName: xml.Name{Local: "evtAfastTemp"},
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

// BuildWorkAccidentXML builds XML for work accident event (S-2260).
// This event records work-related accidents involving workers.
//
// Returns the XML string with proper encoding header
func (s *XMLBuilderService) BuildWorkAccidentXML(params *WorkAccidentParams) (string, error) {
	event := &esocial.Evtcat{
		XMLName: xml.Name{Local: "evtCAT"},
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

// BuildWorkerDismissalXML builds XML for worker dismissal event (S-2280).
// This event records the termination of an employment relationship.
//
// Returns the XML string with proper encoding header
func (s *XMLBuilderService) BuildWorkerDismissalXML(params *WorkerDismissalParams) (string, error) {
	event := &esocial.Evtdeslig{
		XMLName: xml.Name{Local: "evtDeslig"},
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

// BuildWorkerRemunerationXML builds XML for worker remuneration event (S-2400).
// This event records remuneration information for a worker in a period.
//
// Returns the XML string with proper encoding header
func (s *XMLBuilderService) BuildWorkerRemunerationXML(params *WorkerRemunerationParams) (string, error) {
	trabID := esocial.Idetrabalhador{
		Cpftrab: params.Worker.CPF,
	}

	event := &esocial.Evtremun{
		XMLName: xml.Name{Local: "evtRemun"},
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
		Idetrabalhador: trabID,
	}

	xmlData, err := xml.MarshalIndent(event, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal XML: %w", err)
	}

	return xml.Header + string(xmlData), nil
}

// BuildWorkerFGTSXML builds XML for worker FGTS bases event (S-2500).
// This event records the FGTS (severance fund) calculation bases for workers.
//
// Returns the XML string with proper encoding header
func (s *XMLBuilderService) BuildWorkerFGTSXML(params *WorkerFGTSParams) (string, error) {
	event := &esocial.Evtbasesfgts{
		XMLName: xml.Name{Local: "evtBasesFGTS"},
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

// BuildWorkerIRRFXML builds XML for worker IRRF event (S-2600).
// This event records withheld income tax information.
//
// Returns the XML string with proper encoding header
func (s *XMLBuilderService) BuildWorkerIRRFXML(params *WorkerIRRFParams) (string, error) {
	event := &esocial.Evtirrf{
		XMLName: xml.Name{Local: "evtIRRF"},
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

// BuildWorkerContributionXML builds XML for worker contribution event (S-2700).
// This event records social contribution (PIS/PASEP/COFINS) information.
//
// Returns the XML string with proper encoding header
func (s *XMLBuilderService) BuildWorkerContributionXML(params *WorkerContributionParams) (string, error) {
	event := &esocial.Evtcs{
		XMLName: xml.Name{Local: "evtCS"},
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

// BuildPeriodClosureXML builds XML for period closure event (S-2300).
// This event declares the closure of a payroll period.
//
// Returns the XML string with proper encoding header
func (s *XMLBuilderService) BuildPeriodClosureXML(params *PeriodClosureParams) (string, error) {
	event := &esocial.Evtfechaevper{
		XMLName: xml.Name{Local: "evtFechaEvPer"},
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
