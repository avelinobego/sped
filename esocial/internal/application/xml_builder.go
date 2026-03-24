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
//	xml, err := service.BuildEmpresaInfoXML(empresa, nil)
type XMLBuilderService struct{}

// NewXMLBuilderService creates a new XML builder service
func NewXMLBuilderService() *XMLBuilderService {
	return &XMLBuilderService{}
}

// BuildEmpresaInfoXML builds XML for empresa info event (S-1000).
// This event is used to register or update company information in eSocial.
// Parameters:
//   - emp: The company entity from the domain
//   - estabs: Optional list of establishments (currently not implemented in basic version)
//
// Returns the XML string with proper encoding header
func (s *XMLBuilderService) BuildEmpresaInfoXML(emp *empresa.Empresa, estabs []*empresa.Estabelecimento) (string, error) {
	// Create the event
	event := &esocial.Evtinfoempregador{
		Id: "ID123456789", // This should be generated properly
		Ideevento: esocial.Ideevento{
			Tpamb:   1, // Production environment
			Procemi: 1, // Application
			Verproc: "1.0.0",
		},
		Ideempregador: esocial.Ideempregador{
			Tpinsc: 1, // CNPJ
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
// Parameters:
//   - trab: The worker entity from the domain
//   - empresaCNPJ: The CNPJ of the employing company
//
// Returns the XML string with proper encoding header
func (s *XMLBuilderService) BuildTrabalhadorAdmissionXML(trab *trabalhador.Trabalhador, empresaCNPJ string) (string, error) {
	// Create the admission event
	event := &esocial.Evtadmissao{
		Id: "ID123456789", // Should be generated properly
		Ideevento: esocial.Ideevento{
			Tpamb:   1, // Production
			Procemi: 1, // Application
			Verproc: "1.0.0",
		},
		Ideempregador: esocial.Ideempregador{
			Tpinsc: 1, // CNPJ
			Nrinsc: empresaCNPJ,
		},
		Trabalhador: esocial.Trabalhador{
			Cpftrab: trab.CPF,
			Nmtrab:  trab.Nome,
			Sexo:    trab.Sexo,
			Racacor: func() int64 {
				if trab.RacaCor != nil {
					// Convert string to int64 based on eSocial codes
					switch *trab.RacaCor {
					case "1":
						return 1 // Branca
					case "2":
						return 2 // Preta
					case "3":
						return 3 // Parda
					case "4":
						return 4 // Amarela
					case "5":
						return 5 // Indígena
					case "6":
						return 6 // Não informado
					default:
						return 6
					}
				}
				return 6 // Default
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
