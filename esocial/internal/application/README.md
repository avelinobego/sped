# eSocial XML Builder Service

Comprehensive XML builder service for eSocial (Sistema Público de Escrituração Digital - Bloco S) events.

## Overview

The XML Builder Service provides a complete implementation for building valid eSocial XML events from domain entities. It covers all three main domains with builders for the most common eSocial events.

## Architecture

The service is organized into modular builder files:

- **xml_builder.go** - Main service and utility methods
- **empresa_builder.go** - Company/employer-related events (S-1000 series)
- **trabalhador_builder.go** - Worker/employee-related events (S-2000 series)
- **ambiente_trabalho_builder.go** - Work environment-related events

## Domains & Events

### 1. Empresa Domain (Company Events)

#### S-1000: BuildEmpresaInfoXML
Registers or updates company information in eSocial.

```go
xml, err := service.BuildEmpresaInfoXML(empresa, estabelecimentos)
```

#### S-1010: BuildCompleteInfoEmpresaXML
Provides additional details about the company.

```go
xml, err := service.BuildCompleteInfoEmpresaXML(empresa, cnpj)
```

#### S-1020: BuildEstabelecimentoTableXML
Registers establishment (branch/subsidiary) data.

```go
xml, err := service.BuildEstabelecimentoTableXML(estabelecimento, cnpj)
```

#### S-1030: BuildRubricaTableXML
Registers wage types and rubrics for the company.

```go
xml, err := service.BuildRubricaTableXML(rubrica, cnpj)
```

#### S-1040: BuildLotacaoTableXML
Registers work locations and job positions.

```go
xml, err := service.BuildLotacaoTableXML(lotacao, cnpj)
```

#### S-1050: BuildProcessTableXML
Registers legal processes affecting the company/employees.

```go
xml, err := service.BuildProcessTableXML(processoID, numero, cnpj)
```

#### S-1070: BuildToxicAgentTableXML
Registers toxic/hazardous agents present in the workplace.

```go
xml, err := service.BuildToxicAgentTableXML(codigo, descricao, cnpj)
```

### 2. Trabalhador Domain (Worker Events)

#### S-2200: BuildTrabalhadorAdmissionXML
Registers a new employee admission.

```go
xml, err := service.BuildTrabalhadorAdmissionXML(trabalhador, cnpj)
```

#### S-2210: BuildWorkerPreliminaryAdmissionXML
Registers preliminary admission information.

```go
xml, err := service.BuildWorkerPreliminaryAdmissionXML(trabalhador, cnpj)
```

#### S-2230: BuildWorkerAddressChangeXML
Records changes to worker registration data (addresses, contacts, etc.).

```go
xml, err := service.BuildWorkerAddressChangeXML(trabalhador, cnpj, dataAlteracao)
```

#### S-2240: BuildWorkerContractChangeXML
Records changes to employment contract terms.

```go
xml, err := service.BuildWorkerContractChangeXML(trabalhador, cnpj, dataEfetiva)
```

#### S-2250: BuildWorkerTemporaryLeaveXML
Records temporary absences from work (sick leave, statutory leave, etc.).

```go
xml, err := service.BuildWorkerTemporaryLeaveXML(cpf, dtIni, dtFim, cnpj, tipo)
```

#### S-2260: BuildWorkAccidentXML
Records work-related accidents.

```go
xml, err := service.BuildWorkAccidentXML(cpf, dataAcidente, cnpj)
```

#### S-2280: BuildWorkerDismissalXML
Records termination of employment.

```go
xml, err := service.BuildWorkerDismissalXML(trabalhador, cnpj, dataDeslig, motivo)
```

#### S-2300: BuildPeriodClosureXML
Declares closure of a payroll period.

```go
xml, err := service.BuildPeriodClosureXML(cnpj, periodo)
```

#### S-2400: BuildWorkerRemunerationXML
Records remuneration information for a worker.

```go
xml, err := service.BuildWorkerRemunerationXML(trabalhador, cnpj, periodo)
```

#### S-2500: BuildWorkerFGTSXML
Records FGTS (severance fund) calculation bases.

```go
xml, err := service.BuildWorkerFGTSXML(trabalhador, cnpj, periodo)
```

#### S-2600: BuildWorkerIRRFXML
Records withheld income tax information.

```go
xml, err := service.BuildWorkerIRRFXML(trabalhador, cnpj, periodo)
```

#### S-2700: BuildWorkerContributionXML
Records social contribution (PIS/PASEP/COFINS) information.

```go
xml, err := service.BuildWorkerContributionXML(trabalhador, cnpj, periodo)
```

### 3. Ambiente de Trabalho Domain (Work Environment Events)

#### BuildWorkLocationTableXML
Registers work locations where workers are allocated.

```go
xml, err := service.BuildWorkLocationTableXML(ambiente, cnpj)
```

#### BuildWorkScheduleTableXML
Registers work schedules and hour patterns.

```go
xml, err := service.BuildWorkScheduleTableXML(horario, cnpj)
```

#### BuildEnvironmentDataXML
Provides detailed work environment information.

```go
xml, err := service.BuildEnvironmentDataXML(ambiente, cnpj, tipoEvento)
```

#### BuildHazardousEnvironmentDataXML
Registers hazardous/toxic work environments.

```go
xml, err := service.BuildHazardousEnvironmentDataXML(ambiente, cnpj, agenteToxick)
```

#### BuildDepartmentTableXML
Registers company departments and organizational units.

```go
xml, err := service.BuildDepartmentTableXML(codigo, nome, cnpj)
```

#### BuildShiftTableXML
Registers shift patterns and rotating schedules.

```go
xml, err := service.BuildShiftTableXML(codigo, nome, cnpj)
```

## Usage Examples

### Basic Usage

```go
package main

import (
	"fmt"
	"github.com/avelinobego/esocial/internal/application"
	"github.com/avelinobego/esocial/internal/domain/empresa"
	"github.com/avelinobego/esocial/internal/domain/trabalhador"
	"time"
)

func main() {
	// Create service
	service := application.NewXMLBuilderService()

	// Example 1: Create company info XML
	emp := &empresa.Empresa{
		CNPJ:                    "12345678000123",
		RazaoSocial:             "Acme Corp",
		ClassificacaoTributaria: "01",
		IndCooperativa:          "N",
		IndConstrutora:          "N",
		IndDesoneracao:          "N",
	}

	xmlEmpresa, err := service.BuildEmpresaInfoXML(emp, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(xmlEmpresa)

	// Example 2: Create worker admission XML
	trab := &trabalhador.Trabalhador{
		CPF:               "12345678901",
		Nome:              "João Silva",
		Sexo:              "M",
		DataNascimento:    time.Date(1990, 1, 15, 0, 0, 0, 0, time.UTC),
		PaisNascimento:    "105",
		PaisNacionalidade: "105",
	}

	xmlTrabalhador, err := service.BuildTrabalhadorAdmissionXML(trab, "12345678000123")
	if err != nil {
		panic(err)
	}
	fmt.Println(xmlTrabalhador)
}
```

## Generated XML Format

All builder methods return properly formatted XML with:
- XML declaration header
- Proper eSocial event structure
- Correct field mappings from domain entities to eSocial tags
- Indented output for readability

Example output:
```xml
<?xml version="1.0" encoding="UTF-8"?>
<evtInfoEmpregador Id="ID123456789">
  <ideEvento>
    <tpAmb>1</tpAmb>
    <procEmi>1</procEmi>
    <verProc>1.0.0</verProc>
  </ideEvento>
  <ideEmpregador>
    <tpInsc>1</tpInsc>
    <nrInsc>12345678000123</nrInsc>
  </ideEmpregador>
  ...
</evtInfoEmpregador>
```

## Error Handling

All builder methods return an error as the second return value. Always check for errors:

```go
xml, err := service.BuildEmpresaInfoXML(emp, nil)
if err != nil {
    log.Fatalf("Failed to build XML: %v", err)
}
```

## Parameters Guide

### Standard Parameters

- **cnpj** (string): CNPJ format "00000000000000" (14 digits)
- **cpf** (string): CPF format "00000000000" (11 digits)
- **periodo** (string): Format "AAAAMMM" (e.g., "202501" for January 2025)

### Type Codes

- **Tpinsc**: 1 for CNPJ, 2 for CPF, 3 for PIS, 4 for CNO
- **Tpamb**: 1 for Production, 2 for Testing
- **Procemi**: 1 for Application (sistema)

## Testing

Comprehensive test suite included:

```bash
go test -v ./internal/application
```

Tests cover:
- All builder methods
- XML structure validation
- Required field presence
- Error handling

## Notes

- All events use production environment (tpamb=1) by default
- ID generation is placeholder; implement proper ID generation per eSocial spec
- Establishment handling is basic; extend as needed
- Additional fields can be added to events by modifying constructors
- XML marshaling respects all struct xml tags

## References

- [Official eSocial Documentation](https://www.gov.br/esocial/pt-br/documentacao-tecnica/)
- eSocial Layout v13.0 (Circular nº 5/2025, DGP)
- Events S-1000 through S-2700

## Version Compatibility

- Go 1.16+
- eSocial Layout v13.0+

## Future Enhancements

- [ ] ID generation following eSocial specifications
- [ ] Validation of business rules
- [ ] Support for additional events (S-5000 series)
- [ ] Configuration for environment (test/production)
- [ ] Batch event processing
- [ ] Event signing and encryption
