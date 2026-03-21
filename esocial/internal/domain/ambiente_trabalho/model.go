package ambiente_trabalho

import (
	"time"
)

// CreateHorarioRequest represents the request to create a schedule
type CreateHorarioRequest struct {
	EmpresaID             string  `json:"empresa_id" validate:"required"`
	Codigo                string  `json:"codigo" validate:"required,max=30"`
	Descricao             string  `json:"descricao" validate:"required,max=255"`
	TipoJornada           string  `json:"tipo_jornada" validate:"required,oneof=1 2 3 4 5 6 7 9"`
	DuracaoJornadaSemanal *int    `json:"duracao_jornada_semanal,omitempty" validate:"omitempty,min=1,max=168"`
	TipoIntervalo         *string `json:"tipo_intervalo,omitempty" validate:"omitempty,oneof=1 2"`
	DuracaoIntervalo      *int    `json:"duracao_intervalo,omitempty" validate:"omitempty,min=1,max=1440"`
}

// UpdateHorarioRequest represents the request to update a schedule
type UpdateHorarioRequest struct {
	Codigo                *string `json:"codigo,omitempty" validate:"omitempty,max=30"`
	Descricao             *string `json:"descricao,omitempty" validate:"omitempty,max=255"`
	TipoJornada           *string `json:"tipo_jornada,omitempty" validate:"omitempty,oneof=1 2 3 4 5 6 7 9"`
	DuracaoJornadaSemanal *int    `json:"duracao_jornada_semanal,omitempty" validate:"omitempty,min=1,max=168"`
	TipoIntervalo         *string `json:"tipo_intervalo,omitempty" validate:"omitempty,oneof=1 2"`
	DuracaoIntervalo      *int    `json:"duracao_intervalo,omitempty" validate:"omitempty,min=1,max=1440"`
	Ativo                 *bool   `json:"ativo,omitempty"`
}

// HorarioResponse represents the schedule response
type HorarioResponse struct {
	ID                    string                   `json:"id"`
	EmpresaID             string                   `json:"empresa_id"`
	Codigo                string                   `json:"codigo"`
	Descricao             string                   `json:"descricao"`
	TipoJornada           string                   `json:"tipo_jornada"`
	DuracaoJornadaSemanal *int                     `json:"duracao_jornada_semanal,omitempty"`
	TipoIntervalo         *string                  `json:"tipo_intervalo,omitempty"`
	DuracaoIntervalo      *int                     `json:"duracao_intervalo,omitempty"`
	Ativo                 bool                     `json:"ativo"`
	DataCadastro          time.Time                `json:"data_cadastro"`
	Detalhes              []HorarioDetalheResponse `json:"detalhes,omitempty"`
}

// CreateHorarioDetalheRequest represents the request to create schedule details
type CreateHorarioDetalheRequest struct {
	HorarioID      string `json:"horario_id" validate:"required"`
	DiaSemana      int    `json:"dia_semana" validate:"required,min=1,max=7"`
	HorarioEntrada string `json:"horario_entrada" validate:"required"`
	HorarioSaida   string `json:"horario_saida" validate:"required"`
	DuracaoJornada int    `json:"duracao_jornada" validate:"required,min=1,max=1440"`
}

// UpdateHorarioDetalheRequest represents the request to update schedule details
type UpdateHorarioDetalheRequest struct {
	DiaSemana      *int    `json:"dia_semana,omitempty" validate:"omitempty,min=1,max=7"`
	HorarioEntrada *string `json:"horario_entrada,omitempty"`
	HorarioSaida   *string `json:"horario_saida,omitempty"`
	DuracaoJornada *int    `json:"duracao_jornada,omitempty" validate:"omitempty,min=1,max=1440"`
}

// HorarioDetalheResponse represents the schedule detail response
type HorarioDetalheResponse struct {
	ID             string `json:"id"`
	HorarioID      string `json:"horario_id"`
	DiaSemana      int    `json:"dia_semana"`
	HorarioEntrada string `json:"horario_entrada"`
	HorarioSaida   string `json:"horario_saida"`
	DuracaoJornada int    `json:"duracao_jornada"`
}

// CreateAmbienteTrabalhoRequest represents the request to create a work environment
type CreateAmbienteTrabalhoRequest struct {
	EmpresaID       string  `json:"empresa_id" validate:"required"`
	Codigo          string  `json:"codigo" validate:"required,max=30"`
	Descricao       string  `json:"descricao" validate:"required,max=255"`
	LocalAmbiente   string  `json:"local_ambiente" validate:"required,oneof=1 2"`
	TipoInscricao   *string `json:"tipo_inscricao,omitempty" validate:"omitempty,oneof=1 2 3 4"`
	NumeroInscricao *string `json:"numero_inscricao,omitempty" validate:"omitempty,max=15"`
}

// UpdateAmbienteTrabalhoRequest represents the request to update a work environment
type UpdateAmbienteTrabalhoRequest struct {
	Codigo          *string `json:"codigo,omitempty" validate:"omitempty,max=30"`
	Descricao       *string `json:"descricao,omitempty" validate:"omitempty,max=255"`
	LocalAmbiente   *string `json:"local_ambiente,omitempty" validate:"omitempty,oneof=1 2"`
	TipoInscricao   *string `json:"tipo_inscricao,omitempty" validate:"omitempty,oneof=1 2 3 4"`
	NumeroInscricao *string `json:"numero_inscricao,omitempty" validate:"omitempty,max=15"`
	Ativo           *bool   `json:"ativo,omitempty"`
}

// AmbienteTrabalhoResponse represents the work environment response
type AmbienteTrabalhoResponse struct {
	ID              string    `json:"id"`
	EmpresaID       string    `json:"empresa_id"`
	Codigo          string    `json:"codigo"`
	Descricao       string    `json:"descricao"`
	LocalAmbiente   string    `json:"local_ambiente"`
	TipoInscricao   *string   `json:"tipo_inscricao,omitempty"`
	NumeroInscricao *string   `json:"numero_inscricao,omitempty"`
	Ativo           bool      `json:"ativo"`
	DataCadastro    time.Time `json:"data_cadastro"`
}

// ListRequest represents common list/filter parameters
type ListRequest struct {
	Page      int     `json:"page,omitempty" validate:"omitempty,min=1"`
	Limit     int     `json:"limit,omitempty" validate:"omitempty,min=1,max=100"`
	EmpresaID *string `json:"empresa_id,omitempty"`
	HorarioID *string `json:"horario_id,omitempty"`
	Search    *string `json:"search,omitempty"`
	Ativo     *bool   `json:"ativo,omitempty"`
}

// ListResponse represents a paginated list response
type ListResponse[T any] struct {
	Data       []T   `json:"data"`
	Total      int64 `json:"total"`
	Page       int   `json:"page"`
	Limit      int   `json:"limit"`
	TotalPages int   `json:"total_pages"`
}
