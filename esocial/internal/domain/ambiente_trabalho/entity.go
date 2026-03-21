package ambiente_trabalho

import (
	"time"
)

// Horario represents a work schedule entity
type Horario struct {
	ID                    string    `json:"id" db:"id"`
	EmpresaID             string    `json:"empresa_id" db:"empresa_id"`
	Codigo                string    `json:"codigo" db:"codigo"`
	Descricao             string    `json:"descricao" db:"descricao"`
	TipoJornada           string    `json:"tipo_jornada" db:"tipo_jornada"`
	DuracaoJornadaSemanal *int      `json:"duracao_jornada_semanal,omitempty" db:"duracao_jornada_semanal"`
	TipoIntervalo         *string   `json:"tipo_intervalo,omitempty" db:"tipo_intervalo"`
	DuracaoIntervalo      *int      `json:"duracao_intervalo,omitempty" db:"duracao_intervalo"`
	Ativo                 bool      `json:"ativo" db:"ativo"`
	DataCadastro          time.Time `json:"data_cadastro" db:"data_cadastro"`
}

// HorarioDetalhe represents schedule details entity
type HorarioDetalhe struct {
	ID             string `json:"id" db:"id"`
	HorarioID      string `json:"horario_id" db:"horario_id"`
	DiaSemana      int    `json:"dia_semana" db:"dia_semana"`
	HorarioEntrada string `json:"horario_entrada" db:"horario_entrada"`
	HorarioSaida   string `json:"horario_saida" db:"horario_saida"`
	DuracaoJornada int    `json:"duracao_jornada" db:"duracao_jornada"`
}

// AmbienteTrabalho represents work environment entity
type AmbienteTrabalho struct {
	ID              string    `json:"id" db:"id"`
	EmpresaID       string    `json:"empresa_id" db:"empresa_id"`
	Codigo          string    `json:"codigo" db:"codigo"`
	Descricao       string    `json:"descricao" db:"descricao"`
	LocalAmbiente   string    `json:"local_ambiente" db:"local_ambiente"`
	TipoInscricao   *string   `json:"tipo_inscricao,omitempty" db:"tipo_inscricao"`
	NumeroInscricao *string   `json:"numero_inscricao,omitempty" db:"numero_inscricao"`
	Ativo           bool      `json:"ativo" db:"ativo"`
	DataCadastro    time.Time `json:"data_cadastro" db:"data_cadastro"`
}
