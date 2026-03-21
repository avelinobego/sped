package ambiente_trabalho

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/avelinobego/esocial/internal/domain/ambiente_trabalho"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}

// Horario methods
func (r *PostgresRepository) CreateHorario(h *ambiente_trabalho.Horario) error {
	query := `
		INSERT INTO horarios (
			empresa_id, codigo, descricao, tipo_jornada, duracao_jornada_semanal,
			tipo_intervalo, duracao_intervalo
		) VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, data_cadastro`

	err := r.db.QueryRow(query,
		h.EmpresaID, h.Codigo, h.Descricao, h.TipoJornada, h.DuracaoJornadaSemanal,
		h.TipoIntervalo, h.DuracaoIntervalo,
	).Scan(&h.ID, &h.DataCadastro)

	return err
}

func (r *PostgresRepository) GetHorarioByID(id string) (*ambiente_trabalho.Horario, error) {
	query := `SELECT * FROM horarios WHERE id = $1`
	h := &ambiente_trabalho.Horario{}
	err := r.db.QueryRow(query, id).Scan(
		&h.ID, &h.EmpresaID, &h.Codigo, &h.Descricao, &h.TipoJornada,
		&h.DuracaoJornadaSemanal, &h.TipoIntervalo, &h.DuracaoIntervalo,
		&h.Ativo, &h.DataCadastro,
	)
	if err != nil {
		return nil, err
	}
	return h, nil
}

func (r *PostgresRepository) UpdateHorario(id string, updates map[string]interface{}) error {
	if len(updates) == 0 {
		return nil
	}

	setParts := []string{}
	args := []interface{}{}
	argCount := 1

	for field, value := range updates {
		setParts = append(setParts, fmt.Sprintf("%s = $%d", field, argCount))
		args = append(args, value)
		argCount++
	}

	query := fmt.Sprintf("UPDATE horarios SET %s WHERE id = $%d", strings.Join(setParts, ", "), argCount)
	args = append(args, id)

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *PostgresRepository) DeleteHorario(id string) error {
	query := `DELETE FROM horarios WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *PostgresRepository) ListHorarios(filter map[string]interface{}, limit, offset int) ([]*ambiente_trabalho.Horario, int64, error) {
	whereParts := []string{}
	args := []interface{}{}
	argCount := 1

	if empresaID, ok := filter["empresa_id"]; ok && empresaID != nil {
		whereParts = append(whereParts, fmt.Sprintf("empresa_id = $%d", argCount))
		args = append(args, empresaID)
		argCount++
	}

	if search, ok := filter["search"]; ok && search != nil {
		searchStr := fmt.Sprintf("%%%s%%", search)
		whereParts = append(whereParts, fmt.Sprintf("(codigo ILIKE $%d OR descricao ILIKE $%d)", argCount, argCount))
		args = append(args, searchStr)
		argCount++
	}

	if ativo, ok := filter["ativo"]; ok && ativo != nil {
		whereParts = append(whereParts, fmt.Sprintf("ativo = $%d", argCount))
		args = append(args, ativo)
		argCount++
	}

	whereClause := ""
	if len(whereParts) > 0 {
		whereClause = "WHERE " + strings.Join(whereParts, " AND ")
	}

	// Count query
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM horarios %s", whereClause)
	var total int64
	err := r.db.QueryRow(countQuery, args...).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	// Data query
	query := fmt.Sprintf(`
		SELECT * FROM horarios %s
		ORDER BY data_cadastro DESC
		LIMIT $%d OFFSET $%d`, whereClause, argCount, argCount+1)
	args = append(args, limit, offset)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var horarios []*ambiente_trabalho.Horario
	for rows.Next() {
		h := &ambiente_trabalho.Horario{}
		err := rows.Scan(
			&h.ID, &h.EmpresaID, &h.Codigo, &h.Descricao, &h.TipoJornada,
			&h.DuracaoJornadaSemanal, &h.TipoIntervalo, &h.DuracaoIntervalo,
			&h.Ativo, &h.DataCadastro,
		)
		if err != nil {
			return nil, 0, err
		}
		horarios = append(horarios, h)
	}

	return horarios, total, nil
}

// HorarioDetalhe methods
func (r *PostgresRepository) CreateHorarioDetalhe(hd *ambiente_trabalho.HorarioDetalhe) error {
	query := `
		INSERT INTO horarios_detalhes (
			horario_id, dia_semana, horario_entrada, horario_saida, duracao_jornada
		) VALUES ($1, $2, $3, $4, $5)
		RETURNING id`

	err := r.db.QueryRow(query,
		hd.HorarioID, hd.DiaSemana, hd.HorarioEntrada, hd.HorarioSaida, hd.DuracaoJornada,
	).Scan(&hd.ID)

	return err
}

func (r *PostgresRepository) GetHorarioDetalheByID(id string) (*ambiente_trabalho.HorarioDetalhe, error) {
	query := `SELECT * FROM horarios_detalhes WHERE id = $1`
	hd := &ambiente_trabalho.HorarioDetalhe{}
	err := r.db.QueryRow(query, id).Scan(
		&hd.ID, &hd.HorarioID, &hd.DiaSemana, &hd.HorarioEntrada,
		&hd.HorarioSaida, &hd.DuracaoJornada,
	)
	if err != nil {
		return nil, err
	}
	return hd, nil
}

func (r *PostgresRepository) GetHorarioDetalhesByHorarioID(horarioID string) ([]*ambiente_trabalho.HorarioDetalhe, error) {
	query := `SELECT * FROM horarios_detalhes WHERE horario_id = $1 ORDER BY dia_semana`
	rows, err := r.db.Query(query, horarioID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var detalhes []*ambiente_trabalho.HorarioDetalhe
	for rows.Next() {
		hd := &ambiente_trabalho.HorarioDetalhe{}
		err := rows.Scan(
			&hd.ID, &hd.HorarioID, &hd.DiaSemana, &hd.HorarioEntrada,
			&hd.HorarioSaida, &hd.DuracaoJornada,
		)
		if err != nil {
			return nil, err
		}
		detalhes = append(detalhes, hd)
	}

	return detalhes, nil
}

func (r *PostgresRepository) UpdateHorarioDetalhe(id string, updates map[string]interface{}) error {
	if len(updates) == 0 {
		return nil
	}

	setParts := []string{}
	args := []interface{}{}
	argCount := 1

	for field, value := range updates {
		setParts = append(setParts, fmt.Sprintf("%s = $%d", field, argCount))
		args = append(args, value)
		argCount++
	}

	query := fmt.Sprintf("UPDATE horarios_detalhes SET %s WHERE id = $%d", strings.Join(setParts, ", "), argCount)
	args = append(args, id)

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *PostgresRepository) DeleteHorarioDetalhe(id string) error {
	query := `DELETE FROM horarios_detalhes WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *PostgresRepository) DeleteHorarioDetalhesByHorarioID(horarioID string) error {
	query := `DELETE FROM horarios_detalhes WHERE horario_id = $1`
	_, err := r.db.Exec(query, horarioID)
	return err
}

// AmbienteTrabalho methods
func (r *PostgresRepository) CreateAmbienteTrabalho(at *ambiente_trabalho.AmbienteTrabalho) error {
	query := `
		INSERT INTO ambientes_trabalho (
			empresa_id, codigo, descricao, local_ambiente, tipo_inscricao, numero_inscricao
		) VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, data_cadastro`

	err := r.db.QueryRow(query,
		at.EmpresaID, at.Codigo, at.Descricao, at.LocalAmbiente, at.TipoInscricao, at.NumeroInscricao,
	).Scan(&at.ID, &at.DataCadastro)

	return err
}

func (r *PostgresRepository) GetAmbienteTrabalhoByID(id string) (*ambiente_trabalho.AmbienteTrabalho, error) {
	query := `SELECT * FROM ambientes_trabalho WHERE id = $1`
	at := &ambiente_trabalho.AmbienteTrabalho{}
	err := r.db.QueryRow(query, id).Scan(
		&at.ID, &at.EmpresaID, &at.Codigo, &at.Descricao, &at.LocalAmbiente,
		&at.TipoInscricao, &at.NumeroInscricao, &at.Ativo, &at.DataCadastro,
	)
	if err != nil {
		return nil, err
	}
	return at, nil
}

func (r *PostgresRepository) UpdateAmbienteTrabalho(id string, updates map[string]interface{}) error {
	if len(updates) == 0 {
		return nil
	}

	setParts := []string{}
	args := []interface{}{}
	argCount := 1

	for field, value := range updates {
		setParts = append(setParts, fmt.Sprintf("%s = $%d", field, argCount))
		args = append(args, value)
		argCount++
	}

	query := fmt.Sprintf("UPDATE ambientes_trabalho SET %s WHERE id = $%d", strings.Join(setParts, ", "), argCount)
	args = append(args, id)

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *PostgresRepository) DeleteAmbienteTrabalho(id string) error {
	query := `DELETE FROM ambientes_trabalho WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *PostgresRepository) ListAmbientesTrabalho(filter map[string]interface{}, limit, offset int) ([]*ambiente_trabalho.AmbienteTrabalho, int64, error) {
	whereParts := []string{}
	args := []interface{}{}
	argCount := 1

	if empresaID, ok := filter["empresa_id"]; ok && empresaID != nil {
		whereParts = append(whereParts, fmt.Sprintf("empresa_id = $%d", argCount))
		args = append(args, empresaID)
		argCount++
	}

	if search, ok := filter["search"]; ok && search != nil {
		searchStr := fmt.Sprintf("%%%s%%", search)
		whereParts = append(whereParts, fmt.Sprintf("(codigo ILIKE $%d OR descricao ILIKE $%d)", argCount, argCount))
		args = append(args, searchStr)
		argCount++
	}

	if ativo, ok := filter["ativo"]; ok && ativo != nil {
		whereParts = append(whereParts, fmt.Sprintf("ativo = $%d", argCount))
		args = append(args, ativo)
		argCount++
	}

	whereClause := ""
	if len(whereParts) > 0 {
		whereClause = "WHERE " + strings.Join(whereParts, " AND ")
	}

	// Count query
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM ambientes_trabalho %s", whereClause)
	var total int64
	err := r.db.QueryRow(countQuery, args...).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	// Data query
	query := fmt.Sprintf(`
		SELECT * FROM ambientes_trabalho %s
		ORDER BY data_cadastro DESC
		LIMIT $%d OFFSET $%d`, whereClause, argCount, argCount+1)
	args = append(args, limit, offset)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var ambientes []*ambiente_trabalho.AmbienteTrabalho
	for rows.Next() {
		at := &ambiente_trabalho.AmbienteTrabalho{}
		err := rows.Scan(
			&at.ID, &at.EmpresaID, &at.Codigo, &at.Descricao, &at.LocalAmbiente,
			&at.TipoInscricao, &at.NumeroInscricao, &at.Ativo, &at.DataCadastro,
		)
		if err != nil {
			return nil, 0, err
		}
		ambientes = append(ambientes, at)
	}

	return ambientes, total, nil
}
