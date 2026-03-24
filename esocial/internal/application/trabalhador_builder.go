package application

import (
	"encoding/xml"
	"fmt"
	"time"

	"github.com/avelinobego/esocial/internal/domain/trabalhador"
	"github.com/avelinobego/esocial/internal/shared/esocial"
)

// BuildWorkerPreliminaryAdmissionXML builds XML for preliminary worker admission event (S-2210).
// This event registers initial/preliminary information about a new worker admission.
// Parameters:
//   - trab: The worker entity
//   - empresaCNPJ: The CNPJ of the employing company
//
// Returns the XML string with proper encoding header
func (s *XMLBuilderService) BuildWorkerPreliminaryAdmissionXML(trab *trabalhador.Trabalhador, empresaCNPJ string) (string, error) {
	event := &esocial.Evtadmprelim{
		XMLName: xml.Name{Local: "evtAdmPrelim"},
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

// BuildWorkerAddressChangeXML builds XML for worker registration change event (S-2230).
// This event records changes to worker registration data (addresses, contact info, etc.).
// Parameters:
//   - trab: The updated worker entity
//   - empresaCNPJ: The CNPJ of the company
//   - alteracaoData: Date of the change
//
// Returns the XML string with proper encoding header
func (s *XMLBuilderService) BuildWorkerAddressChangeXML(trab *trabalhador.Trabalhador, empresaCNPJ string, alteracaoData time.Time) (string, error) {
	event := &esocial.Evtaltcadastral{
		XMLName: xml.Name{Local: "evtAltCadastral"},
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

// BuildWorkerContractChangeXML builds XML for worker contract change event (S-2240).
// This event records changes to employment contract terms and conditions.
// Parameters:
//   - trab: The worker entity
//   - empresaCNPJ: The CNPJ of the company
//   - dtEfetiva: Effective date of the change
//
// Returns the XML string with proper encoding header
func (s *XMLBuilderService) BuildWorkerContractChangeXML(trab *trabalhador.Trabalhador, empresaCNPJ string, dtEfetiva time.Time) (string, error) {
	event := &esocial.Evtaltcontratual{
		XMLName: xml.Name{Local: "evtAltContratual"},
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

// BuildWorkerTemporaryLeaveXML builds XML for worker temporary leave event (S-2250).
// This event records temporary absences from work (sick leave, statutory leave, etc.).
// Parameters:
//   - trabalhadorCPF: The worker's CPF
//   - dtIniAfastamento: Start date of the leave
//   - dtFimAfastamento: End date of the leave (optional)
//   - empresaCNPJ: The CNPJ of the company
//   - tipoAfastamento: Type of leave (1-Doença, 2-Acidente, etc.)
//
// Returns the XML string with proper encoding header
func (s *XMLBuilderService) BuildWorkerTemporaryLeaveXML(trabalhadorCPF string, dtIniAfastamento, dtFimAfastamento *time.Time, empresaCNPJ string, tipoAfastamento string) (string, error) {
	event := &esocial.Evtafasttemp{
		XMLName: xml.Name{Local: "evtAfastTemp"},
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

// BuildWorkAccidentXML builds XML for work accident event (S-2260).
// This event records work-related accidents involving workers.
// Parameters:
//   - trabalhadorCPF: The worker's CPF
//   - dtAcidente: Date of the accident
//   - empresaCNPJ: The CNPJ of the company
//
// Returns the XML string with proper encoding header
func (s *XMLBuilderService) BuildWorkAccidentXML(trabalhadorCPF string, dtAcidente time.Time, empresaCNPJ string) (string, error) {
	event := &esocial.Evtcat{
		XMLName: xml.Name{Local: "evtCAT"},
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

// BuildWorkerDismissalXML builds XML for worker dismissal event (S-2280).
// This event records the termination of an employment relationship.
// Parameters:
//   - trab: The worker entity
//   - empresaCNPJ: The CNPJ of the company
//   - dtDeslig: Dismissal date
//   - motivoDesligamento: Reason for dismissal
//
// Returns the XML string with proper encoding header
func (s *XMLBuilderService) BuildWorkerDismissalXML(trab *trabalhador.Trabalhador, empresaCNPJ string, dtDeslig time.Time, motivoDesligamento string) (string, error) {
	event := &esocial.Evtdeslig{
		XMLName: xml.Name{Local: "evtDeslig"},
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

// BuildWorkerRemunerationXML builds XML for worker remuneration event (S-2400).
// This event records remuneration information for a worker in a period.
// Parameters:
//   - trab: The worker entity
//   - empresaCNPJ: The CNPJ of the company
//   - periodo: Period for the remuneration (AAAAMM format)
//
// Returns the XML string with proper encoding header
func (s *XMLBuilderService) BuildWorkerRemunerationXML(trab *trabalhador.Trabalhador, empresaCNPJ string, periodo string) (string, error) {
	trabID := esocial.Idetrabalhador{
		Cpftrab: trab.CPF,
	}

	event := &esocial.Evtremun{
		XMLName: xml.Name{Local: "evtRemun"},
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
// Parameters:
//   - trab: The worker entity
//   - empresaCNPJ: The CNPJ of the company
//   - periodo: Period for the FGTS (AAAAMM format)
//
// Returns the XML string with proper encoding header
func (s *XMLBuilderService) BuildWorkerFGTSXML(trab *trabalhador.Trabalhador, empresaCNPJ string, periodo string) (string, error) {
	event := &esocial.Evtbasesfgts{
		XMLName: xml.Name{Local: "evtBasesFGTS"},
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

// BuildWorkerIRRFXML builds XML for worker IRRF event (S-2600).
// This event records withheld income tax information.
// Parameters:
//   - trab: The worker entity
//   - empresaCNPJ: The CNPJ of the company
//   - periodo: Period for the IRRF (AAAAMM format)
//
// Returns the XML string with proper encoding header
func (s *XMLBuilderService) BuildWorkerIRRFXML(trab *trabalhador.Trabalhador, empresaCNPJ string, periodo string) (string, error) {
	event := &esocial.Evtirrf{
		XMLName: xml.Name{Local: "evtIRRF"},
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

// BuildWorkerContributionXML builds XML for worker contribution event (S-2700).
// This event records social contribution (PIS/PASEP/COFINS) information.
// Parameters:
//   - trab: The worker entity
//   - empresaCNPJ: The CNPJ of the company
//   - periodo: Period for the contribution (AAAAMM format)
//
// Returns the XML string with proper encoding header
func (s *XMLBuilderService) BuildWorkerContributionXML(trab *trabalhador.Trabalhador, empresaCNPJ string, periodo string) (string, error) {
	event := &esocial.Evtcs{
		XMLName: xml.Name{Local: "evtCS"},
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

// BuildPeriodClosureXML builds XML for period closure event (S-2300).
// This event declares the closure of a payroll period.
// Parameters:
//   - empresaCNPJ: The CNPJ of the company
//   - periodo: Period to be closed (AAAAMM format)
//
// Returns the XML string with proper encoding header
func (s *XMLBuilderService) BuildPeriodClosureXML(empresaCNPJ string, periodo string) (string, error) {
	event := &esocial.Evtfechaevper{
		XMLName: xml.Name{Local: "evtFechaEvPer"},
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
