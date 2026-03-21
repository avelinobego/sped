package trabalhador

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/lib/pq"
)

// Repository defines the interface for trabalhador data operations
type Repository interface {
	Create(ctx context.Context, trabalhador *Trabalhador) error
	GetByID(ctx context.Context, id string) (*Trabalhador, error)
	GetByCPF(ctx context.Context, cpf string) (*Trabalhador, error)
	Update(ctx context.Context, trabalhador *Trabalhador) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, req *ListRequest) ([]*Trabalhador, int64, error)
}

// PostgresRepository implements Repository interface for PostgreSQL
type PostgresRepository struct {
	db *sql.DB
}

// NewPostgresRepository creates a new PostgreSQL repository
func NewPostgresRepository(db *sql.DB) Repository {
	return &PostgresRepository{db: db}
}

// Create inserts a new trabalhador record
func (r *PostgresRepository) Create(ctx context.Context, trabalhador *Trabalhador) error {
	query := `
		INSERT INTO trabalhador (
			id, empresa_id, cpf, nis, nome, nome_social, data_nascimento, sexo,
			raca_cor, estado_civil, grau_instrucao, pais_nascimento, pais_nacionalidade,
			numero_ctps, serie_ctps, uf_ctps, data_emissao_ctps, numero_rg, orgao_emissor_rg,
			uf_rg, data_emissao_rg, numero_cnh, categoria_cnh, validade_cnh, numero_rne,
			telefone, celular, email, cep, logradouro, numero, complemento, bairro,
			cidade, uf, banco, agencia, conta, tipo_conta, possui_deficiencia,
			tipo_deficiencia, observacao_deficiencia, situacao, data_cadastro,
			data_atualizacao, foto_url
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17,
			$18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32,
			$33, $34, $35, $36, $37, $38, $39, $40, $41, $42, $43, $44, $45, $46, $47
		)`

	_, err := r.db.ExecContext(ctx, query,
		trabalhador.ID, trabalhador.EmpresaID, trabalhador.CPF, trabalhador.NIS,
		trabalhador.Nome, trabalhador.NomeSocial, trabalhador.DataNascimento,
		trabalhador.Sexo, trabalhador.RacaCor, trabalhador.EstadoCivil,
		trabalhador.GrauInstrucao, trabalhador.PaisNascimento, trabalhador.PaisNacionalidade,
		trabalhador.NumeroCTPS, trabalhador.SerieCTPS, trabalhador.UFCTPS,
		trabalhador.DataEmissaoCTPS, trabalhador.NumeroRG, trabalhador.OrgaoEmissorRG,
		trabalhador.UFRG, trabalhador.DataEmissaoRG, trabalhador.NumeroCNH,
		trabalhador.CategoriaCNH, trabalhador.ValidadeCNH, trabalhador.NumeroRNE,
		trabalhador.Telefone, trabalhador.Celular, trabalhador.Email, trabalhador.CEP,
		trabalhador.Logradouro, trabalhador.Numero, trabalhador.Complemento,
		trabalhador.Bairro, trabalhador.Cidade, trabalhador.UF, trabalhador.Banco,
		trabalhador.Agencia, trabalhador.Conta, trabalhador.TipoConta,
		trabalhador.PossuiDeficiencia, trabalhador.TipoDeficiencia,
		trabalhador.ObservacaoDeficiencia, trabalhador.Situacao,
		trabalhador.DataCadastro, trabalhador.DataAtualizacao, trabalhador.FotoURL,
	)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code {
			case "23505": // unique_violation
				return fmt.Errorf("trabalhador with CPF %s already exists", trabalhador.CPF)
			case "23503": // foreign_key_violation
				return fmt.Errorf("empresa with ID %s does not exist", trabalhador.EmpresaID)
			}
		}
		return fmt.Errorf("failed to create trabalhador: %w", err)
	}

	return nil
}

// GetByID retrieves a trabalhador by ID
func (r *PostgresRepository) GetByID(ctx context.Context, id string) (*Trabalhador, error) {
	query := `
		SELECT id, empresa_id, cpf, nis, nome, nome_social, data_nascimento, sexo,
			   raca_cor, estado_civil, grau_instrucao, pais_nascimento, pais_nacionalidade,
			   numero_ctps, serie_ctps, uf_ctps, data_emissao_ctps, numero_rg, orgao_emissor_rg,
			   uf_rg, data_emissao_rg, numero_cnh, categoria_cnh, validade_cnh, numero_rne,
			   telefone, celular, email, cep, logradouro, numero, complemento, bairro,
			   cidade, uf, banco, agencia, conta, tipo_conta, possui_deficiencia,
			   tipo_deficiencia, observacao_deficiencia, situacao, data_cadastro,
			   data_atualizacao, foto_url
		FROM trabalhador
		WHERE id = $1`

	var trabalhador Trabalhador
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&trabalhador.ID, &trabalhador.EmpresaID, &trabalhador.CPF, &trabalhador.NIS,
		&trabalhador.Nome, &trabalhador.NomeSocial, &trabalhador.DataNascimento,
		&trabalhador.Sexo, &trabalhador.RacaCor, &trabalhador.EstadoCivil,
		&trabalhador.GrauInstrucao, &trabalhador.PaisNascimento, &trabalhador.PaisNacionalidade,
		&trabalhador.NumeroCTPS, &trabalhador.SerieCTPS, &trabalhador.UFCTPS,
		&trabalhador.DataEmissaoCTPS, &trabalhador.NumeroRG, &trabalhador.OrgaoEmissorRG,
		&trabalhador.UFRG, &trabalhador.DataEmissaoRG, &trabalhador.NumeroCNH,
		&trabalhador.CategoriaCNH, &trabalhador.ValidadeCNH, &trabalhador.NumeroRNE,
		&trabalhador.Telefone, &trabalhador.Celular, &trabalhador.Email, &trabalhador.CEP,
		&trabalhador.Logradouro, &trabalhador.Numero, &trabalhador.Complemento,
		&trabalhador.Bairro, &trabalhador.Cidade, &trabalhador.UF, &trabalhador.Banco,
		&trabalhador.Agencia, &trabalhador.Conta, &trabalhador.TipoConta,
		&trabalhador.PossuiDeficiencia, &trabalhador.TipoDeficiencia,
		&trabalhador.ObservacaoDeficiencia, &trabalhador.Situacao,
		&trabalhador.DataCadastro, &trabalhador.DataAtualizacao, &trabalhador.FotoURL,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("trabalhador with ID %s not found", id)
		}
		return nil, fmt.Errorf("failed to get trabalhador: %w", err)
	}

	return &trabalhador, nil
}

// GetByCPF retrieves a trabalhador by CPF
func (r *PostgresRepository) GetByCPF(ctx context.Context, cpf string) (*Trabalhador, error) {
	query := `
		SELECT id, empresa_id, cpf, nis, nome, nome_social, data_nascimento, sexo,
			   raca_cor, estado_civil, grau_instrucao, pais_nascimento, pais_nacionalidade,
			   numero_ctps, serie_ctps, uf_ctps, data_emissao_ctps, numero_rg, orgao_emissor_rg,
			   uf_rg, data_emissao_rg, numero_cnh, categoria_cnh, validade_cnh, numero_rne,
			   telefone, celular, email, cep, logradouro, numero, complemento, bairro,
			   cidade, uf, banco, agencia, conta, tipo_conta, possui_deficiencia,
			   tipo_deficiencia, observacao_deficiencia, situacao, data_cadastro,
			   data_atualizacao, foto_url
		FROM trabalhador
		WHERE cpf = $1`

	var trabalhador Trabalhador
	err := r.db.QueryRowContext(ctx, query, cpf).Scan(
		&trabalhador.ID, &trabalhador.EmpresaID, &trabalhador.CPF, &trabalhador.NIS,
		&trabalhador.Nome, &trabalhador.NomeSocial, &trabalhador.DataNascimento,
		&trabalhador.Sexo, &trabalhador.RacaCor, &trabalhador.EstadoCivil,
		&trabalhador.GrauInstrucao, &trabalhador.PaisNascimento, &trabalhador.PaisNacionalidade,
		&trabalhador.NumeroCTPS, &trabalhador.SerieCTPS, &trabalhador.UFCTPS,
		&trabalhador.DataEmissaoCTPS, &trabalhador.NumeroRG, &trabalhador.OrgaoEmissorRG,
		&trabalhador.UFRG, &trabalhador.DataEmissaoRG, &trabalhador.NumeroCNH,
		&trabalhador.CategoriaCNH, &trabalhador.ValidadeCNH, &trabalhador.NumeroRNE,
		&trabalhador.Telefone, &trabalhador.Celular, &trabalhador.Email, &trabalhador.CEP,
		&trabalhador.Logradouro, &trabalhador.Numero, &trabalhador.Complemento,
		&trabalhador.Bairro, &trabalhador.Cidade, &trabalhador.UF, &trabalhador.Banco,
		&trabalhador.Agencia, &trabalhador.Conta, &trabalhador.TipoConta,
		&trabalhador.PossuiDeficiencia, &trabalhador.TipoDeficiencia,
		&trabalhador.ObservacaoDeficiencia, &trabalhador.Situacao,
		&trabalhador.DataCadastro, &trabalhador.DataAtualizacao, &trabalhador.FotoURL,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("trabalhador with CPF %s not found", cpf)
		}
		return nil, fmt.Errorf("failed to get trabalhador: %w", err)
	}

	return &trabalhador, nil
}

// Update modifies an existing trabalhador record
func (r *PostgresRepository) Update(ctx context.Context, trabalhador *Trabalhador) error {
	query := `
		UPDATE trabalhador SET
			nis = $2, nome = $3, nome_social = $4, data_nascimento = $5, sexo = $6,
			raca_cor = $7, estado_civil = $8, grau_instrucao = $9, pais_nascimento = $10,
			pais_nacionalidade = $11, numero_ctps = $12, serie_ctps = $13, uf_ctps = $14,
			data_emissao_ctps = $15, numero_rg = $16, orgao_emissor_rg = $17, uf_rg = $18,
			data_emissao_rg = $19, numero_cnh = $20, categoria_cnh = $21, validade_cnh = $22,
			numero_rne = $23, telefone = $24, celular = $25, email = $26, cep = $27,
			logradouro = $28, numero = $29, complemento = $30, bairro = $31, cidade = $32,
			uf = $33, banco = $34, agencia = $35, conta = $36, tipo_conta = $37,
			possui_deficiencia = $38, tipo_deficiencia = $39, observacao_deficiencia = $40,
			situacao = $41, data_atualizacao = $42, foto_url = $43
		WHERE id = $1`

	result, err := r.db.ExecContext(ctx, query,
		trabalhador.ID, trabalhador.NIS, trabalhador.Nome, trabalhador.NomeSocial,
		trabalhador.DataNascimento, trabalhador.Sexo, trabalhador.RacaCor,
		trabalhador.EstadoCivil, trabalhador.GrauInstrucao, trabalhador.PaisNascimento,
		trabalhador.PaisNacionalidade, trabalhador.NumeroCTPS, trabalhador.SerieCTPS,
		trabalhador.UFCTPS, trabalhador.DataEmissaoCTPS, trabalhador.NumeroRG,
		trabalhador.OrgaoEmissorRG, trabalhador.UFRG, trabalhador.DataEmissaoRG,
		trabalhador.NumeroCNH, trabalhador.CategoriaCNH, trabalhador.ValidadeCNH,
		trabalhador.NumeroRNE, trabalhador.Telefone, trabalhador.Celular,
		trabalhador.Email, trabalhador.CEP, trabalhador.Logradouro, trabalhador.Numero,
		trabalhador.Complemento, trabalhador.Bairro, trabalhador.Cidade, trabalhador.UF,
		trabalhador.Banco, trabalhador.Agencia, trabalhador.Conta, trabalhador.TipoConta,
		trabalhador.PossuiDeficiencia, trabalhador.TipoDeficiencia,
		trabalhador.ObservacaoDeficiencia, trabalhador.Situacao,
		trabalhador.DataAtualizacao, trabalhador.FotoURL,
	)

	if err != nil {
		return fmt.Errorf("failed to update trabalhador: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("trabalhador with ID %s not found", trabalhador.ID)
	}

	return nil
}

// Delete removes a trabalhador record
func (r *PostgresRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM trabalhador WHERE id = $1`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete trabalhador: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("trabalhador with ID %s not found", id)
	}

	return nil
}

// List retrieves trabalhadores with pagination and filtering
func (r *PostgresRepository) List(ctx context.Context, req *ListRequest) ([]*Trabalhador, int64, error) {
	// Build the base query
	selectQuery := `
		SELECT id, empresa_id, cpf, nis, nome, nome_social, data_nascimento, sexo,
			   raca_cor, estado_civil, grau_instrucao, pais_nascimento, pais_nacionalidade,
			   numero_ctps, serie_ctps, uf_ctps, data_emissao_ctps, numero_rg, orgao_emissor_rg,
			   uf_rg, data_emissao_rg, numero_cnh, categoria_cnh, validade_cnh, numero_rne,
			   telefone, celular, email, cep, logradouro, numero, complemento, bairro,
			   cidade, uf, banco, agencia, conta, tipo_conta, possui_deficiencia,
			   tipo_deficiencia, observacao_deficiencia, situacao, data_cadastro,
			   data_atualizacao, foto_url
		FROM trabalhador`

	countQuery := `SELECT COUNT(*) FROM trabalhador`

	var conditions []string
	var args []interface{}
	argIndex := 1

	// Add filters
	if req.EmpresaID != nil {
		conditions = append(conditions, fmt.Sprintf("empresa_id = $%d", argIndex))
		args = append(args, *req.EmpresaID)
		argIndex++
	}

	if req.Situacao != nil {
		conditions = append(conditions, fmt.Sprintf("situacao = $%d", argIndex))
		args = append(args, *req.Situacao)
		argIndex++
	}

	if req.Sexo != nil {
		conditions = append(conditions, fmt.Sprintf("sexo = $%d", argIndex))
		args = append(args, *req.Sexo)
		argIndex++
	}

	if req.PossuiDeficiencia != nil {
		conditions = append(conditions, fmt.Sprintf("possui_deficiencia = $%d", argIndex))
		args = append(args, *req.PossuiDeficiencia)
		argIndex++
	}

	if req.Search != nil && *req.Search != "" {
		searchTerm := "%" + *req.Search + "%"
		conditions = append(conditions, fmt.Sprintf("(nome ILIKE $%d OR cpf ILIKE $%d)", argIndex, argIndex+1))
		args = append(args, searchTerm, searchTerm)
		argIndex += 2
	}

	// Add WHERE clause if there are conditions
	if len(conditions) > 0 {
		whereClause := " WHERE " + strings.Join(conditions, " AND ")
		selectQuery += whereClause
		countQuery += whereClause
	}

	// Add ordering
	selectQuery += " ORDER BY nome ASC"

	// Add pagination
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.Limit <= 0 {
		req.Limit = 10
	}
	offset := (req.Page - 1) * req.Limit
	selectQuery += fmt.Sprintf(" LIMIT %d OFFSET %d", req.Limit, offset)

	// Execute count query
	var total int64
	err := r.db.QueryRowContext(ctx, countQuery, args...).Scan(&total)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to count trabalhadores: %w", err)
	}

	// Execute select query
	rows, err := r.db.QueryContext(ctx, selectQuery, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to list trabalhadores: %w", err)
	}
	defer rows.Close()

	var trabalhadores []*Trabalhador
	for rows.Next() {
		var trabalhador Trabalhador
		err := rows.Scan(
			&trabalhador.ID, &trabalhador.EmpresaID, &trabalhador.CPF, &trabalhador.NIS,
			&trabalhador.Nome, &trabalhador.NomeSocial, &trabalhador.DataNascimento,
			&trabalhador.Sexo, &trabalhador.RacaCor, &trabalhador.EstadoCivil,
			&trabalhador.GrauInstrucao, &trabalhador.PaisNascimento, &trabalhador.PaisNacionalidade,
			&trabalhador.NumeroCTPS, &trabalhador.SerieCTPS, &trabalhador.UFCTPS,
			&trabalhador.DataEmissaoCTPS, &trabalhador.NumeroRG, &trabalhador.OrgaoEmissorRG,
			&trabalhador.UFRG, &trabalhador.DataEmissaoRG, &trabalhador.NumeroCNH,
			&trabalhador.CategoriaCNH, &trabalhador.ValidadeCNH, &trabalhador.NumeroRNE,
			&trabalhador.Telefone, &trabalhador.Celular, &trabalhador.Email, &trabalhador.CEP,
			&trabalhador.Logradouro, &trabalhador.Numero, &trabalhador.Complemento,
			&trabalhador.Bairro, &trabalhador.Cidade, &trabalhador.UF, &trabalhador.Banco,
			&trabalhador.Agencia, &trabalhador.Conta, &trabalhador.TipoConta,
			&trabalhador.PossuiDeficiencia, &trabalhador.TipoDeficiencia,
			&trabalhador.ObservacaoDeficiencia, &trabalhador.Situacao,
			&trabalhador.DataCadastro, &trabalhador.DataAtualizacao, &trabalhador.FotoURL,
		)
		if err != nil {
			return nil, 0, fmt.Errorf("failed to scan trabalhador: %w", err)
		}
		trabalhadores = append(trabalhadores, &trabalhador)
	}

	if err = rows.Err(); err != nil {
		return nil, 0, fmt.Errorf("error iterating trabalhadores: %w", err)
	}

	return trabalhadores, total, nil
}
