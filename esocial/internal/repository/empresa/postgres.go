package empresa

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/avelinobego/esocial/internal/domain/empresa"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}

// Empresa methods
func (r *PostgresRepository) CreateEmpresa(e *empresa.Empresa) error {
	query := `
		INSERT INTO empresas (
			cnpj, razao_social, nome_fantasia, tipo_inscricao, classificacao_tributaria,
			natureza_juridica, ind_cooperativa, ind_construtora, ind_desoneracao,
			usuario_cadastro, telefone, email, logradouro, numero, complemento,
			bairro, cidade, uf, cep
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19)
		RETURNING id, data_cadastro, data_atualizacao`

	err := r.db.QueryRow(query,
		e.CNPJ, e.RazaoSocial, e.NomeFantasia, e.TipoInscricao, e.ClassificacaoTributaria,
		e.NaturezaJuridica, e.IndCooperativa, e.IndConstrutora, e.IndDesoneracao,
		e.UsuarioCadastro, e.Telefone, e.Email, e.Logradouro, e.Numero, e.Complemento,
		e.Bairro, e.Cidade, e.UF, e.CEP,
	).Scan(&e.ID, &e.DataCadastro, &e.DataAtualizacao)

	return err
}

func (r *PostgresRepository) GetEmpresaByID(id string) (*empresa.Empresa, error) {
	query := `SELECT * FROM empresas WHERE id = $1`
	e := &empresa.Empresa{}
	err := r.db.QueryRow(query, id).Scan(
		&e.ID, &e.CNPJ, &e.RazaoSocial, &e.NomeFantasia, &e.TipoInscricao, &e.ClassificacaoTributaria,
		&e.NaturezaJuridica, &e.IndCooperativa, &e.IndConstrutora, &e.IndDesoneracao, &e.Situacao,
		&e.DataCadastro, &e.DataAtualizacao, &e.UsuarioCadastro, &e.Telefone, &e.Email,
		&e.Logradouro, &e.Numero, &e.Complemento, &e.Bairro, &e.Cidade, &e.UF, &e.CEP,
		&e.ProtocoloESocial, &e.ReciboESocial, &e.DataEnvioESocial, &e.StatusESocial,
	)
	if err != nil {
		return nil, err
	}
	return e, nil
}

func (r *PostgresRepository) GetEmpresaByCNPJ(cnpj string) (*empresa.Empresa, error) {
	query := `SELECT * FROM empresas WHERE cnpj = $1`
	e := &empresa.Empresa{}
	err := r.db.QueryRow(query, cnpj).Scan(
		&e.ID, &e.CNPJ, &e.RazaoSocial, &e.NomeFantasia, &e.TipoInscricao, &e.ClassificacaoTributaria,
		&e.NaturezaJuridica, &e.IndCooperativa, &e.IndConstrutora, &e.IndDesoneracao, &e.Situacao,
		&e.DataCadastro, &e.DataAtualizacao, &e.UsuarioCadastro, &e.Telefone, &e.Email,
		&e.Logradouro, &e.Numero, &e.Complemento, &e.Bairro, &e.Cidade, &e.UF, &e.CEP,
		&e.ProtocoloESocial, &e.ReciboESocial, &e.DataEnvioESocial, &e.StatusESocial,
	)
	if err != nil {
		return nil, err
	}
	return e, nil
}

func (r *PostgresRepository) UpdateEmpresa(id string, updates map[string]interface{}) error {
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

	setParts = append(setParts, "data_atualizacao = CURRENT_TIMESTAMP")
	query := fmt.Sprintf("UPDATE empresas SET %s WHERE id = $%d", strings.Join(setParts, ", "), argCount)
	args = append(args, id)

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *PostgresRepository) DeleteEmpresa(id string) error {
	query := `DELETE FROM empresas WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *PostgresRepository) ListEmpresas(filter map[string]interface{}, limit, offset int) ([]*empresa.Empresa, int64, error) {
	whereParts := []string{}
	args := []interface{}{}
	argCount := 1

	if empresaID, ok := filter["empresa_id"]; ok && empresaID != nil {
		whereParts = append(whereParts, fmt.Sprintf("id = $%d", argCount))
		args = append(args, empresaID)
		argCount++
	}

	if search, ok := filter["search"]; ok && search != nil {
		searchStr := fmt.Sprintf("%%%s%%", search)
		whereParts = append(whereParts, fmt.Sprintf("(razao_social ILIKE $%d OR nome_fantasia ILIKE $%d OR cnpj ILIKE $%d)", argCount, argCount, argCount))
		args = append(args, searchStr)
		argCount++
	}

	if situacao, ok := filter["situacao"]; ok && situacao != nil {
		whereParts = append(whereParts, fmt.Sprintf("situacao = $%d", argCount))
		args = append(args, situacao)
		argCount++
	}

	whereClause := ""
	if len(whereParts) > 0 {
		whereClause = "WHERE " + strings.Join(whereParts, " AND ")
	}

	// Count query
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM empresas %s", whereClause)
	var total int64
	err := r.db.QueryRow(countQuery, args...).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	// Data query
	query := fmt.Sprintf(`
		SELECT * FROM empresas %s
		ORDER BY data_cadastro DESC
		LIMIT $%d OFFSET $%d`, whereClause, argCount, argCount+1)
	args = append(args, limit, offset)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var empresas []*empresa.Empresa
	for rows.Next() {
		e := &empresa.Empresa{}
		err := rows.Scan(
			&e.ID, &e.CNPJ, &e.RazaoSocial, &e.NomeFantasia, &e.TipoInscricao, &e.ClassificacaoTributaria,
			&e.NaturezaJuridica, &e.IndCooperativa, &e.IndConstrutora, &e.IndDesoneracao, &e.Situacao,
			&e.DataCadastro, &e.DataAtualizacao, &e.UsuarioCadastro, &e.Telefone, &e.Email,
			&e.Logradouro, &e.Numero, &e.Complemento, &e.Bairro, &e.Cidade, &e.UF, &e.CEP,
			&e.ProtocoloESocial, &e.ReciboESocial, &e.DataEnvioESocial, &e.StatusESocial,
		)
		if err != nil {
			return nil, 0, err
		}
		empresas = append(empresas, e)
	}

	return empresas, total, nil
}

// Estabelecimento methods
func (r *PostgresRepository) CreateEstabelecimento(e *empresa.Estabelecimento) error {
	query := `
		INSERT INTO estabelecimentos (
			empresa_id, tipo_inscricao, numero_inscricao, cnae_principal, cnae_preparatorio,
			alvara_funcionamento, data_inicio_atividades, ind_obra, logradouro, numero,
			complemento, bairro, cidade, uf, cep
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)
		RETURNING id, data_cadastro`

	err := r.db.QueryRow(query,
		e.EmpresaID, e.TipoInscricao, e.NumeroInscricao, e.CNAEPrincipal, e.CNAEPreparatorio,
		e.AlvaraFuncionamento, e.DataInicioAtividades, e.IndObra, e.Logradouro, e.Numero,
		e.Complemento, e.Bairro, e.Cidade, e.UF, e.CEP,
	).Scan(&e.ID, &e.DataCadastro)

	return err
}

func (r *PostgresRepository) GetEstabelecimentoByID(id string) (*empresa.Estabelecimento, error) {
	query := `SELECT * FROM estabelecimentos WHERE id = $1`
	e := &empresa.Estabelecimento{}
	err := r.db.QueryRow(query, id).Scan(
		&e.ID, &e.EmpresaID, &e.TipoInscricao, &e.NumeroInscricao, &e.CNAEPrincipal,
		&e.CNAEPreparatorio, &e.AlvaraFuncionamento, &e.DataInicioAtividades, &e.IndObra,
		&e.Logradouro, &e.Numero, &e.Complemento, &e.Bairro, &e.Cidade, &e.UF, &e.CEP,
		&e.Situacao, &e.DataCadastro,
	)
	if err != nil {
		return nil, err
	}
	return e, nil
}

func (r *PostgresRepository) UpdateEstabelecimento(id string, updates map[string]interface{}) error {
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

	query := fmt.Sprintf("UPDATE estabelecimentos SET %s WHERE id = $%d", strings.Join(setParts, ", "), argCount)
	args = append(args, id)

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *PostgresRepository) DeleteEstabelecimento(id string) error {
	query := `DELETE FROM estabelecimentos WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *PostgresRepository) ListEstabelecimentos(filter map[string]interface{}, limit, offset int) ([]*empresa.Estabelecimento, int64, error) {
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
		whereParts = append(whereParts, fmt.Sprintf("(numero_inscricao ILIKE $%d)", argCount))
		args = append(args, searchStr)
		argCount++
	}

	if situacao, ok := filter["situacao"]; ok && situacao != nil {
		whereParts = append(whereParts, fmt.Sprintf("situacao = $%d", argCount))
		args = append(args, situacao)
		argCount++
	}

	whereClause := ""
	if len(whereParts) > 0 {
		whereClause = "WHERE " + strings.Join(whereParts, ", ")
	}

	// Count query
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM estabelecimentos %s", whereClause)
	var total int64
	err := r.db.QueryRow(countQuery, args...).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	// Data query
	query := fmt.Sprintf(`
		SELECT * FROM estabelecimentos %s
		ORDER BY data_cadastro DESC
		LIMIT $%d OFFSET $%d`, whereClause, argCount, argCount+1)
	args = append(args, limit, offset)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var estabelecimentos []*empresa.Estabelecimento
	for rows.Next() {
		e := &empresa.Estabelecimento{}
		err := rows.Scan(
			&e.ID, &e.EmpresaID, &e.TipoInscricao, &e.NumeroInscricao, &e.CNAEPrincipal,
			&e.CNAEPreparatorio, &e.AlvaraFuncionamento, &e.DataInicioAtividades, &e.IndObra,
			&e.Logradouro, &e.Numero, &e.Complemento, &e.Bairro, &e.Cidade, &e.UF, &e.CEP,
			&e.Situacao, &e.DataCadastro,
		)
		if err != nil {
			return nil, 0, err
		}
		estabelecimentos = append(estabelecimentos, e)
	}

	return estabelecimentos, total, nil
}

// Rubrica methods
func (r *PostgresRepository) CreateRubrica(rb *empresa.Rubrica) error {
	query := `
		INSERT INTO rubricas (
			empresa_id, codigo, descricao, natureza_rubrica, tipo_rubrica,
			cod_inccp, cod_incirrf, cod_incfgts, cod_incsind, data_inicio_validade, data_fim_validade
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
		RETURNING id, data_cadastro`

	err := r.db.QueryRow(query,
		rb.EmpresaID, rb.Codigo, rb.Descricao, rb.NaturezaRubrica, rb.TipoRubrica,
		rb.CodINCCP, rb.CodINCIRRF, rb.CodINCFGTS, rb.CodINCSind, rb.DataInicioValidade, rb.DataFimValidade,
	).Scan(&rb.ID, &rb.DataCadastro)

	return err
}

func (r *PostgresRepository) GetRubricaByID(id string) (*empresa.Rubrica, error) {
	query := `SELECT * FROM rubricas WHERE id = $1`
	rb := &empresa.Rubrica{}
	err := r.db.QueryRow(query, id).Scan(
		&rb.ID, &rb.EmpresaID, &rb.Codigo, &rb.Descricao, &rb.NaturezaRubrica, &rb.TipoRubrica,
		&rb.CodINCCP, &rb.CodINCIRRF, &rb.CodINCFGTS, &rb.CodINCSind, &rb.Ativa,
		&rb.DataCadastro, &rb.DataInicioValidade, &rb.DataFimValidade,
	)
	if err != nil {
		return nil, err
	}
	return rb, nil
}

func (r *PostgresRepository) UpdateRubrica(id string, updates map[string]interface{}) error {
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

	query := fmt.Sprintf("UPDATE rubricas SET %s WHERE id = $%d", strings.Join(setParts, ", "), argCount)
	args = append(args, id)

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *PostgresRepository) DeleteRubrica(id string) error {
	query := `DELETE FROM rubricas WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *PostgresRepository) ListRubricas(filter map[string]interface{}, limit, offset int) ([]*empresa.Rubrica, int64, error) {
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

	if ativa, ok := filter["ativa"]; ok && ativa != nil {
		whereParts = append(whereParts, fmt.Sprintf("ativa = $%d", argCount))
		args = append(args, ativa)
		argCount++
	}

	// Add validity date filter
	now := time.Now()
	whereParts = append(whereParts, fmt.Sprintf("(data_inicio_validade <= $%d AND (data_fim_validade IS NULL OR data_fim_validade >= $%d))", argCount, argCount))
	args = append(args, now, now)
	argCount++

	whereClause := ""
	if len(whereParts) > 0 {
		whereClause = "WHERE " + strings.Join(whereParts, " AND ")
	}

	// Count query
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM rubricas %s", whereClause)
	var total int64
	err := r.db.QueryRow(countQuery, args...).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	// Data query
	query := fmt.Sprintf(`
		SELECT * FROM rubricas %s
		ORDER BY data_cadastro DESC
		LIMIT $%d OFFSET $%d`, whereClause, argCount, argCount+1)
	args = append(args, limit, offset)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var rubricas []*empresa.Rubrica
	for rows.Next() {
		rb := &empresa.Rubrica{}
		err := rows.Scan(
			&rb.ID, &rb.EmpresaID, &rb.Codigo, &rb.Descricao, &rb.NaturezaRubrica, &rb.TipoRubrica,
			&rb.CodINCCP, &rb.CodINCIRRF, &rb.CodINCFGTS, &rb.CodINCSind, &rb.Ativa,
			&rb.DataCadastro, &rb.DataInicioValidade, &rb.DataFimValidade,
		)
		if err != nil {
			return nil, 0, err
		}
		rubricas = append(rubricas, rb)
	}

	return rubricas, total, nil
}

// Lotacao methods
func (r *PostgresRepository) CreateLotacao(l *empresa.Lotacao) error {
	query := `
		INSERT INTO lotacoes (
			empresa_id, codigo, descricao, tipo_lotacao, fpas, cod_terceiros, data_inicio_validade, data_fim_validade
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id, data_cadastro`

	err := r.db.QueryRow(query,
		l.EmpresaID, l.Codigo, l.Descricao, l.TipoLotacao, l.FPAS, l.CodTerceiros, l.DataInicioValidade, l.DataFimValidade,
	).Scan(&l.ID, &l.DataCadastro)

	return err
}

func (r *PostgresRepository) GetLotacaoByID(id string) (*empresa.Lotacao, error) {
	query := `SELECT * FROM lotacoes WHERE id = $1`
	l := &empresa.Lotacao{}
	err := r.db.QueryRow(query, id).Scan(
		&l.ID, &l.EmpresaID, &l.Codigo, &l.Descricao, &l.TipoLotacao, &l.FPAS, &l.CodTerceiros,
		&l.Ativa, &l.DataCadastro, &l.DataInicioValidade, &l.DataFimValidade,
	)
	if err != nil {
		return nil, err
	}
	return l, nil
}

func (r *PostgresRepository) UpdateLotacao(id string, updates map[string]interface{}) error {
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

	query := fmt.Sprintf("UPDATE lotacoes SET %s WHERE id = $%d", strings.Join(setParts, ", "), argCount)
	args = append(args, id)

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *PostgresRepository) DeleteLotacao(id string) error {
	query := `DELETE FROM lotacoes WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *PostgresRepository) ListLotacoes(filter map[string]interface{}, limit, offset int) ([]*empresa.Lotacao, int64, error) {
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

	if ativa, ok := filter["ativa"]; ok && ativa != nil {
		whereParts = append(whereParts, fmt.Sprintf("ativa = $%d", argCount))
		args = append(args, ativa)
		argCount++
	}

	// Add validity date filter
	now := time.Now()
	whereParts = append(whereParts, fmt.Sprintf("(data_inicio_validade <= $%d AND (data_fim_validade IS NULL OR data_fim_validade >= $%d))", argCount, argCount))
	args = append(args, now, now)
	argCount++

	whereClause := ""
	if len(whereParts) > 0 {
		whereClause = "WHERE " + strings.Join(whereParts, " AND ")
	}

	// Count query
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM lotacoes %s", whereClause)
	var total int64
	err := r.db.QueryRow(countQuery, args...).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	// Data query
	query := fmt.Sprintf(`
		SELECT * FROM lotacoes %s
		ORDER BY data_cadastro DESC
		LIMIT $%d OFFSET $%d`, whereClause, argCount, argCount+1)
	args = append(args, limit, offset)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var lotacoes []*empresa.Lotacao
	for rows.Next() {
		l := &empresa.Lotacao{}
		err := rows.Scan(
			&l.ID, &l.EmpresaID, &l.Codigo, &l.Descricao, &l.TipoLotacao, &l.FPAS, &l.CodTerceiros,
			&l.Ativa, &l.DataCadastro, &l.DataInicioValidade, &l.DataFimValidade,
		)
		if err != nil {
			return nil, 0, err
		}
		lotacoes = append(lotacoes, l)
	}

	return lotacoes, total, nil
}

// Cargo methods
func (r *PostgresRepository) CreateCargo(c *empresa.Cargo) error {
	query := `
		INSERT INTO cargos (
			empresa_id, codigo, descricao, cbo, data_inicio_validade, data_fim_validade
		) VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, data_cadastro`

	err := r.db.QueryRow(query,
		c.EmpresaID, c.Codigo, c.Descricao, c.CBO, c.DataInicioValidade, c.DataFimValidade,
	).Scan(&c.ID, &c.DataCadastro)

	return err
}

func (r *PostgresRepository) GetCargoByID(id string) (*empresa.Cargo, error) {
	query := `SELECT * FROM cargos WHERE id = $1`
	c := &empresa.Cargo{}
	err := r.db.QueryRow(query, id).Scan(
		&c.ID, &c.EmpresaID, &c.Codigo, &c.Descricao, &c.CBO, &c.Ativo,
		&c.DataCadastro, &c.DataInicioValidade, &c.DataFimValidade,
	)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *PostgresRepository) UpdateCargo(id string, updates map[string]interface{}) error {
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

	query := fmt.Sprintf("UPDATE cargos SET %s WHERE id = $%d", strings.Join(setParts, ", "), argCount)
	args = append(args, id)

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *PostgresRepository) DeleteCargo(id string) error {
	query := `DELETE FROM cargos WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *PostgresRepository) ListCargos(filter map[string]interface{}, limit, offset int) ([]*empresa.Cargo, int64, error) {
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

	// Add validity date filter
	now := time.Now()
	whereParts = append(whereParts, fmt.Sprintf("(data_inicio_validade <= $%d AND (data_fim_validade IS NULL OR data_fim_validade >= $%d))", argCount, argCount))
	args = append(args, now, now)
	argCount++

	whereClause := ""
	if len(whereParts) > 0 {
		whereClause = "WHERE " + strings.Join(whereParts, " AND ")
	}

	// Count query
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM cargos %s", whereClause)
	var total int64
	err := r.db.QueryRow(countQuery, args...).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	// Data query
	query := fmt.Sprintf(`
		SELECT * FROM cargos %s
		ORDER BY data_cadastro DESC
		LIMIT $%d OFFSET $%d`, whereClause, argCount, argCount+1)
	args = append(args, limit, offset)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var cargos []*empresa.Cargo
	for rows.Next() {
		c := &empresa.Cargo{}
		err := rows.Scan(
			&c.ID, &c.EmpresaID, &c.Codigo, &c.Descricao, &c.CBO, &c.Ativo,
			&c.DataCadastro, &c.DataInicioValidade, &c.DataFimValidade,
		)
		if err != nil {
			return nil, 0, err
		}
		cargos = append(cargos, c)
	}

	return cargos, total, nil
}
