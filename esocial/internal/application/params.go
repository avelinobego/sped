// Package application provides high-level services for the eSocial system
package application

import (
	"time"

	"github.com/avelinobego/esocial/internal/domain/ambiente_trabalho"
	"github.com/avelinobego/esocial/internal/domain/empresa"
	"github.com/avelinobego/esocial/internal/domain/trabalhador"
)

// EventParams contains common parameters for all eSocial events
type EventParams struct {
	EmpresaCNPJ string
	TpAmb       int64
	ProcEmi     int64
	VerProc     string
	TpInsc      int64
}

// WorkerPreliminaryAdmissionParams parameters for building worker preliminary admission XML
type WorkerPreliminaryAdmissionParams struct {
	EventParams
	Worker *trabalhador.Trabalhador
}

// WorkerAddressChangeParams parameters for building worker address change XML
type WorkerAddressChangeParams struct {
	EventParams
	Worker        *trabalhador.Trabalhador
	AlteracaoData time.Time
}

// WorkerContractChangeParams parameters for building worker contract change XML
type WorkerContractChangeParams struct {
	EventParams
	Worker    *trabalhador.Trabalhador
	DtEfetiva time.Time
}

// WorkerTemporaryLeaveParams parameters for building worker temporary leave XML
type WorkerTemporaryLeaveParams struct {
	EventParams
	TrabalhadorCPF   string
	DtIniAfastamento *time.Time
	DtFimAfastamento *time.Time
	TipoAfastamento  string
}

// WorkAccidentParams parameters for building work accident XML
type WorkAccidentParams struct {
	EventParams
	TrabalhadorCPF string
	DtAcidente     time.Time
}

// WorkerDismissalParams parameters for building worker dismissal XML
type WorkerDismissalParams struct {
	EventParams
	Worker             *trabalhador.Trabalhador
	DtDeslig           time.Time
	MotivoDesligamento string
}

// WorkerRemunerationParams parameters for building worker remuneration XML
type WorkerRemunerationParams struct {
	EventParams
	Worker  *trabalhador.Trabalhador
	Periodo string
}

// WorkerFGTSParams parameters for building worker FGTS XML
type WorkerFGTSParams struct {
	EventParams
	Worker  *trabalhador.Trabalhador
	Periodo string
}

// WorkerIRRFParams parameters for building worker IRRF XML
type WorkerIRRFParams struct {
	EventParams
	Worker  *trabalhador.Trabalhador
	Periodo string
}

// WorkerContributionParams parameters for building worker contribution XML
type WorkerContributionParams struct {
	EventParams
	Worker  *trabalhador.Trabalhador
	Periodo string
}

// PeriodClosureParams parameters for building period closure XML
type PeriodClosureParams struct {
	EventParams
	Periodo string
}

// EstabelecimentoTableParams parameters for building establishment table XML
type EstabelecimentoTableParams struct {
	EventParams
	Estabelecimento *empresa.Estabelecimento
}

// RubricaTableParams parameters for building rubrica table XML
type RubricaTableParams struct {
	EventParams
	Rubrica *empresa.Rubrica
}

// LotacaoTableParams parameters for building lotacao table XML
type LotacaoTableParams struct {
	EventParams
	Lotacao *empresa.Lotacao
}

// ProcessTableParams parameters for building process table XML
type ProcessTableParams struct {
	EventParams
	ProcessoID     string
	NumeroProcesso string
}

// CargoTableParams parameters for building cargo table XML
type CargoTableParams struct {
	EventParams
	CodigoCargo string
	NomeCargo   string
}

// ToxicAgentTableParams parameters for building toxic agent table XML
type ToxicAgentTableParams struct {
	EventParams
	CodigoAgente    string
	DescricaoAgente string
}

// CompleteInfoEmpresaParams parameters for building complete empresa info XML
type CompleteInfoEmpresaParams struct {
	EventParams
	Empresa *empresa.Empresa
}

// WorkLocationTableParams parameters for building work location table XML
type WorkLocationTableParams struct {
	EventParams
	Ambiente *ambiente_trabalho.AmbienteTrabalho
}

// WorkScheduleTableParams parameters for building work schedule table XML
type WorkScheduleTableParams struct {
	EventParams
	Horario *ambiente_trabalho.Horario
}

// EnvironmentDataParams parameters for building environment data XML
type EnvironmentDataParams struct {
	EventParams
	Ambiente   *ambiente_trabalho.AmbienteTrabalho
	TipoEvento string
}

// HazardousEnvironmentDataParams parameters for building hazardous environment data XML
type HazardousEnvironmentDataParams struct {
	EventParams
	Ambiente     *ambiente_trabalho.AmbienteTrabalho
	AgenteToxick string
}

// DepartmentTableParams parameters for building department table XML
type DepartmentTableParams struct {
	EventParams
	CodigoDept string
	NomeDept   string
}

// ShiftTableParams parameters for building shift table XML
type ShiftTableParams struct {
	EventParams
	CodigoTurno string
	NomeTurno   string
}
