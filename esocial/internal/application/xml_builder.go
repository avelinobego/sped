// Package application provides high-level services for the eSocial system
package application

import (
	"encoding/xml"
	"fmt"

	"github.com/avelinobego/esocial/internal/domain/empresa"
	"github.com/avelinobego/esocial/internal/domain/trabalhador"
	"github.com/avelinobego/esocial/internal/shared/esocial"
)

// XMLBuilderService handles building XML from domain entities using eSocial structs.
// It provides methods to convert business domain objects into valid eSocial XML events
// that can be sent to the Brazilian government systems.
//
// Usage:
//
//	service := NewXMLBuilderService()
//	params := &EventParams{EmpresaCNPJ: "12345678000123", TpAmb: 1, ProcEmi: 1, VerProc: "1.0.0", TpInsc: 1}
//	xml, err := service.BuildEmpresaInfoXML(empresa, params)
type XMLBuilderService struct{}

// NewXMLBuilderService creates a new XML builder service
func NewXMLBuilderService() *XMLBuilderService {
	return &XMLBuilderService{}
}

// BuildEmpresaInfoXML builds XML for empresa info event (S-1000).
// This event is used to register or update company information in eSocial.
//
// Returns the XML string with proper encoding header
func (s *XMLBuilderService) BuildEmpresaInfoXML(emp *empresa.Empresa, params *EventParams) (string, error) {
	// Create the event
	event := &esocial.Evtinfoempregador{
		Id: "ID123456789", // This should be generated properly
		Ideevento: esocial.Ideevento{
			Tpamb:   params.TpAmb,
			Procemi: params.ProcEmi,
			Verproc: params.VerProc,
		},
		Ideempregador: esocial.Ideempregador{
			Tpinsc: params.TpInsc,
			Nrinsc: emp.CNPJ,
		},
		Infoempregador: esocial.Infoempregador{
			Inclusao: &esocial.Inclusao{
				Infocadastro: &esocial.Infocadastro{
					Classtrib: emp.ClassificacaoTributaria,
					Indcoop: func() []string {
						if emp.IndCooperativa == "S" {
							return []string{"S"}
						}
						return []string{"N"}
					}(),
					Indconstr: func() []string {
						if emp.IndConstrutora == "S" {
							return []string{"S"}
						}
						return []string{"N"}
					}(),
					Inddesfolha: func() int64 {
						if emp.IndDesoneracao == "S" {
							return 1
						}
						return 0
					}(),
					Indopccp:         []string{"2"}, // Not subject to CPP
					Indoptregeletron: 0,             // Not opting for electronic receipt
					// Add other fields as needed
				},
			},
		},
	}

	// For establishments, we might need separate events or handle differently
	// For now, skip establishments in this basic implementation

	// Marshal to XML
	xmlData, err := xml.MarshalIndent(event, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal XML: %w", err)
	}

	return xml.Header + string(xmlData), nil
}

// BuildTrabalhadorAdmissionXML builds XML for trabalhador admission event (S-2200).
// This event is used to register a new employee admission in eSocial.
//
// Returns the XML string with proper encoding header
func (s *XMLBuilderService) BuildTrabalhadorAdmissionXML(trab *trabalhador.Trabalhador, params *EventParams) (string, error) {
	// Create the admission event
	event := &esocial.Evtadmissao{
		Id: "ID123456789", // Should be generated properly
		Ideevento: esocial.Ideevento{
			Tpamb:   params.TpAmb,
			Procemi: params.ProcEmi,
			Verproc: params.VerProc,
		},
		Ideempregador: esocial.Ideempregador{
			Tpinsc: params.TpInsc,
			Nrinsc: params.EmpresaCNPJ,
		},
		Trabalhador: esocial.Trabalhador{
			Cpftrab: trab.CPF,
			Nmtrab:  trab.Nome,
			Sexo:    trab.Sexo,
			Racacor: func() int64 {
				if trab.RacaCor != nil {
					// Convert string to RacaCor enumeration
					return ParseRacaCor(*trab.RacaCor).ToInt64()
				}
				return RacaCorNãoInformado.ToInt64() // Default
			}(),
			Grauinstr: func() string {
				if trab.GrauInstrucao != nil {
					return *trab.GrauInstrucao
				}
				return "00" // Default
			}(),
			Nascimento: esocial.Nascimento{
				Dtnascto:   trab.DataNascimento.Format("2006-01-02"),
				Paisnascto: trab.PaisNascimento,
				Paisnac:    trab.PaisNacionalidade,
			},
			// Add more fields as needed: endereco, contato, etc.
		},
		Vinculo: esocial.Vinculo{
			TpRegPrev: 1, // Default regime
		},
		// Add InfoContrato, InfoRegimeTrab, etc. as needed
	}

	// Marshal to XML
	xmlData, err := xml.MarshalIndent(event, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal XML: %w", err)
	}

	return xml.Header + string(xmlData), nil
}

// BuildGenericEventXML builds XML for any eSocial event.
// This is a utility method for events that don't have specific builder methods yet.
// Parameters:
//   - event: Any struct that can be marshaled to XML (should have proper xml tags)
//
// Returns the XML string with proper encoding header
func (s *XMLBuilderService) BuildGenericEventXML(event interface{}) (string, error) {
	xmlData, err := xml.MarshalIndent(event, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal XML: %w", err)
	}

	return xml.Header + string(xmlData), nil
}
