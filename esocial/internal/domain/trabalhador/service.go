package trabalhador

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

// Service defines the interface for trabalhador business logic
type Service interface {
	Create(ctx context.Context, req *CreateTrabalhadorRequest) (*TrabalhadorResponse, error)
	GetByID(ctx context.Context, id string) (*TrabalhadorResponse, error)
	GetByCPF(ctx context.Context, cpf string) (*TrabalhadorResponse, error)
	Update(ctx context.Context, id string, req *UpdateTrabalhadorRequest) (*TrabalhadorResponse, error)
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, req *ListRequest) (*ListResponse[TrabalhadorResponse], error)
}

// TrabalhadorService implements Service interface
type TrabalhadorService struct {
	repo Repository
}

// NewService creates a new trabalhador service
func NewService(repo Repository) Service {
	return &TrabalhadorService{repo: repo}
}

// Create creates a new trabalhador
func (s *TrabalhadorService) Create(ctx context.Context, req *CreateTrabalhadorRequest) (*TrabalhadorResponse, error) {
	// Check if trabalhador with CPF already exists
	existing, err := s.repo.GetByCPF(ctx, req.CPF)
	if err == nil && existing != nil {
		return nil, fmt.Errorf("trabalhador with CPF %s already exists", req.CPF)
	}

	// Create trabalhador entity
	now := time.Now()
	trabalhador := &Trabalhador{
		ID:                    uuid.New().String(),
		EmpresaID:             req.EmpresaID,
		CPF:                   req.CPF,
		NIS:                   req.NIS,
		Nome:                  req.Nome,
		NomeSocial:            req.NomeSocial,
		DataNascimento:        req.DataNascimento,
		Sexo:                  req.Sexo,
		RacaCor:               req.RacaCor,
		EstadoCivil:           req.EstadoCivil,
		GrauInstrucao:         req.GrauInstrucao,
		PaisNascimento:        req.PaisNascimento,
		PaisNacionalidade:     req.PaisNacionalidade,
		NumeroCTPS:            req.NumeroCTPS,
		SerieCTPS:             req.SerieCTPS,
		UFCTPS:                req.UFCTPS,
		DataEmissaoCTPS:       req.DataEmissaoCTPS,
		NumeroRG:              req.NumeroRG,
		OrgaoEmissorRG:        req.OrgaoEmissorRG,
		UFRG:                  req.UFRG,
		DataEmissaoRG:         req.DataEmissaoRG,
		NumeroCNH:             req.NumeroCNH,
		CategoriaCNH:          req.CategoriaCNH,
		ValidadeCNH:           req.ValidadeCNH,
		NumeroRNE:             req.NumeroRNE,
		Telefone:              req.Telefone,
		Celular:               req.Celular,
		Email:                 req.Email,
		CEP:                   req.CEP,
		Logradouro:            req.Logradouro,
		Numero:                req.Numero,
		Complemento:           req.Complemento,
		Bairro:                req.Bairro,
		Cidade:                req.Cidade,
		UF:                    req.UF,
		Banco:                 req.Banco,
		Agencia:               req.Agencia,
		Conta:                 req.Conta,
		TipoConta:             req.TipoConta,
		PossuiDeficiencia:     req.PossuiDeficiencia,
		TipoDeficiencia:       req.TipoDeficiencia,
		ObservacaoDeficiencia: req.ObservacaoDeficiencia,
		Situacao:              "ATIVO",
		DataCadastro:          now,
		DataAtualizacao:       now,
		FotoURL:               req.FotoURL,
	}

	// Save to repository
	err = s.repo.Create(ctx, trabalhador)
	if err != nil {
		return nil, fmt.Errorf("failed to create trabalhador: %w", err)
	}

	// Convert to response
	response := s.entityToResponse(trabalhador)
	return response, nil
}

// GetByID retrieves a trabalhador by ID
func (s *TrabalhadorService) GetByID(ctx context.Context, id string) (*TrabalhadorResponse, error) {
	trabalhador, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	response := s.entityToResponse(trabalhador)
	return response, nil
}

// GetByCPF retrieves a trabalhador by CPF
func (s *TrabalhadorService) GetByCPF(ctx context.Context, cpf string) (*TrabalhadorResponse, error) {
	trabalhador, err := s.repo.GetByCPF(ctx, cpf)
	if err != nil {
		return nil, err
	}

	response := s.entityToResponse(trabalhador)
	return response, nil
}

// Update modifies an existing trabalhador
func (s *TrabalhadorService) Update(ctx context.Context, id string, req *UpdateTrabalhadorRequest) (*TrabalhadorResponse, error) {
	// Get existing trabalhador
	trabalhador, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// Apply updates
	if req.NIS != nil {
		trabalhador.NIS = req.NIS
	}
	if req.Nome != nil {
		trabalhador.Nome = *req.Nome
	}
	if req.NomeSocial != nil {
		trabalhador.NomeSocial = req.NomeSocial
	}
	if req.DataNascimento != nil {
		trabalhador.DataNascimento = *req.DataNascimento
	}
	if req.Sexo != nil {
		trabalhador.Sexo = *req.Sexo
	}
	if req.RacaCor != nil {
		trabalhador.RacaCor = req.RacaCor
	}
	if req.EstadoCivil != nil {
		trabalhador.EstadoCivil = req.EstadoCivil
	}
	if req.GrauInstrucao != nil {
		trabalhador.GrauInstrucao = req.GrauInstrucao
	}
	if req.PaisNascimento != nil {
		trabalhador.PaisNascimento = *req.PaisNascimento
	}
	if req.PaisNacionalidade != nil {
		trabalhador.PaisNacionalidade = *req.PaisNacionalidade
	}
	if req.NumeroCTPS != nil {
		trabalhador.NumeroCTPS = req.NumeroCTPS
	}
	if req.SerieCTPS != nil {
		trabalhador.SerieCTPS = req.SerieCTPS
	}
	if req.UFCTPS != nil {
		trabalhador.UFCTPS = req.UFCTPS
	}
	if req.DataEmissaoCTPS != nil {
		trabalhador.DataEmissaoCTPS = req.DataEmissaoCTPS
	}
	if req.NumeroRG != nil {
		trabalhador.NumeroRG = req.NumeroRG
	}
	if req.OrgaoEmissorRG != nil {
		trabalhador.OrgaoEmissorRG = req.OrgaoEmissorRG
	}
	if req.UFRG != nil {
		trabalhador.UFRG = req.UFRG
	}
	if req.DataEmissaoRG != nil {
		trabalhador.DataEmissaoRG = req.DataEmissaoRG
	}
	if req.NumeroCNH != nil {
		trabalhador.NumeroCNH = req.NumeroCNH
	}
	if req.CategoriaCNH != nil {
		trabalhador.CategoriaCNH = req.CategoriaCNH
	}
	if req.ValidadeCNH != nil {
		trabalhador.ValidadeCNH = req.ValidadeCNH
	}
	if req.NumeroRNE != nil {
		trabalhador.NumeroRNE = req.NumeroRNE
	}
	if req.Telefone != nil {
		trabalhador.Telefone = req.Telefone
	}
	if req.Celular != nil {
		trabalhador.Celular = req.Celular
	}
	if req.Email != nil {
		trabalhador.Email = req.Email
	}
	if req.CEP != nil {
		trabalhador.CEP = req.CEP
	}
	if req.Logradouro != nil {
		trabalhador.Logradouro = req.Logradouro
	}
	if req.Numero != nil {
		trabalhador.Numero = req.Numero
	}
	if req.Complemento != nil {
		trabalhador.Complemento = req.Complemento
	}
	if req.Bairro != nil {
		trabalhador.Bairro = req.Bairro
	}
	if req.Cidade != nil {
		trabalhador.Cidade = req.Cidade
	}
	if req.UF != nil {
		trabalhador.UF = req.UF
	}
	if req.Banco != nil {
		trabalhador.Banco = req.Banco
	}
	if req.Agencia != nil {
		trabalhador.Agencia = req.Agencia
	}
	if req.Conta != nil {
		trabalhador.Conta = req.Conta
	}
	if req.TipoConta != nil {
		trabalhador.TipoConta = req.TipoConta
	}
	if req.PossuiDeficiencia != nil {
		trabalhador.PossuiDeficiencia = *req.PossuiDeficiencia
	}
	if req.TipoDeficiencia != nil {
		trabalhador.TipoDeficiencia = req.TipoDeficiencia
	}
	if req.ObservacaoDeficiencia != nil {
		trabalhador.ObservacaoDeficiencia = req.ObservacaoDeficiencia
	}
	if req.Situacao != nil {
		trabalhador.Situacao = *req.Situacao
	}
	if req.FotoURL != nil {
		trabalhador.FotoURL = req.FotoURL
	}

	trabalhador.DataAtualizacao = time.Now()

	// Save to repository
	err = s.repo.Update(ctx, trabalhador)
	if err != nil {
		return nil, fmt.Errorf("failed to update trabalhador: %w", err)
	}

	// Convert to response
	response := s.entityToResponse(trabalhador)
	return response, nil
}

// Delete removes a trabalhador
func (s *TrabalhadorService) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}

// List retrieves trabalhadores with pagination and filtering
func (s *TrabalhadorService) List(ctx context.Context, req *ListRequest) (*ListResponse[TrabalhadorResponse], error) {
	trabalhadores, total, err := s.repo.List(ctx, req)
	if err != nil {
		return nil, err
	}

	// Convert entities to responses
	responses := make([]TrabalhadorResponse, len(trabalhadores))
	for i, trabalhador := range trabalhadores {
		responses[i] = *s.entityToResponse(trabalhador)
	}

	// Calculate total pages
	totalPages := int((total + int64(req.Limit) - 1) / int64(req.Limit))
	if totalPages == 0 {
		totalPages = 1
	}

	return &ListResponse[TrabalhadorResponse]{
		Data:       responses,
		Total:      total,
		Page:       req.Page,
		Limit:      req.Limit,
		TotalPages: totalPages,
	}, nil
}

// entityToResponse converts entity to response DTO
func (s *TrabalhadorService) entityToResponse(trabalhador *Trabalhador) *TrabalhadorResponse {
	return &TrabalhadorResponse{
		ID:                    trabalhador.ID,
		EmpresaID:             trabalhador.EmpresaID,
		CPF:                   trabalhador.CPF,
		NIS:                   trabalhador.NIS,
		Nome:                  trabalhador.Nome,
		NomeSocial:            trabalhador.NomeSocial,
		DataNascimento:        trabalhador.DataNascimento,
		Sexo:                  trabalhador.Sexo,
		RacaCor:               trabalhador.RacaCor,
		EstadoCivil:           trabalhador.EstadoCivil,
		GrauInstrucao:         trabalhador.GrauInstrucao,
		PaisNascimento:        trabalhador.PaisNascimento,
		PaisNacionalidade:     trabalhador.PaisNacionalidade,
		NumeroCTPS:            trabalhador.NumeroCTPS,
		SerieCTPS:             trabalhador.SerieCTPS,
		UFCTPS:                trabalhador.UFCTPS,
		DataEmissaoCTPS:       trabalhador.DataEmissaoCTPS,
		NumeroRG:              trabalhador.NumeroRG,
		OrgaoEmissorRG:        trabalhador.OrgaoEmissorRG,
		UFRG:                  trabalhador.UFRG,
		DataEmissaoRG:         trabalhador.DataEmissaoRG,
		NumeroCNH:             trabalhador.NumeroCNH,
		CategoriaCNH:          trabalhador.CategoriaCNH,
		ValidadeCNH:           trabalhador.ValidadeCNH,
		NumeroRNE:             trabalhador.NumeroRNE,
		Telefone:              trabalhador.Telefone,
		Celular:               trabalhador.Celular,
		Email:                 trabalhador.Email,
		CEP:                   trabalhador.CEP,
		Logradouro:            trabalhador.Logradouro,
		Numero:                trabalhador.Numero,
		Complemento:           trabalhador.Complemento,
		Bairro:                trabalhador.Bairro,
		Cidade:                trabalhador.Cidade,
		UF:                    trabalhador.UF,
		Banco:                 trabalhador.Banco,
		Agencia:               trabalhador.Agencia,
		Conta:                 trabalhador.Conta,
		TipoConta:             trabalhador.TipoConta,
		PossuiDeficiencia:     trabalhador.PossuiDeficiencia,
		TipoDeficiencia:       trabalhador.TipoDeficiencia,
		ObservacaoDeficiencia: trabalhador.ObservacaoDeficiencia,
		Situacao:              trabalhador.Situacao,
		DataCadastro:          trabalhador.DataCadastro,
		DataAtualizacao:       trabalhador.DataAtualizacao,
		FotoURL:               trabalhador.FotoURL,
	}
}
