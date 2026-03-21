package empresa

import (
	"errors"

	"github.com/avelinobego/esocial/configs"
	"github.com/avelinobego/esocial/pkg/utils"
)

var (
	ErrEmpresaNotFound         = errors.New("empresa not found")
	ErrEstabelecimentoNotFound = errors.New("estabelecimento not found")
	ErrRubricaNotFound         = errors.New("rubrica not found")
	ErrLotacaoNotFound         = errors.New("lotacao not found")
	ErrCargoNotFound           = errors.New("cargo not found")
	ErrEmpresaAlreadyExists    = errors.New("empresa already exists")
	ErrInvalidCNPJ             = errors.New("invalid CNPJ")
)

type Repository interface {
	// Empresa methods
	CreateEmpresa(*Empresa) error
	GetEmpresaByID(string) (*Empresa, error)
	GetEmpresaByCNPJ(string) (*Empresa, error)
	UpdateEmpresa(string, map[string]interface{}) error
	DeleteEmpresa(string) error
	ListEmpresas(map[string]interface{}, int, int) ([]*Empresa, int64, error)

	// Estabelecimento methods
	CreateEstabelecimento(*Estabelecimento) error
	GetEstabelecimentoByID(string) (*Estabelecimento, error)
	UpdateEstabelecimento(string, map[string]interface{}) error
	DeleteEstabelecimento(string) error
	ListEstabelecimentos(map[string]interface{}, int, int) ([]*Estabelecimento, int64, error)

	// Rubrica methods
	CreateRubrica(*Rubrica) error
	GetRubricaByID(string) (*Rubrica, error)
	UpdateRubrica(string, map[string]interface{}) error
	DeleteRubrica(string) error
	ListRubricas(map[string]interface{}, int, int) ([]*Rubrica, int64, error)

	// Lotacao methods
	CreateLotacao(*Lotacao) error
	GetLotacaoByID(string) (*Lotacao, error)
	UpdateLotacao(string, map[string]interface{}) error
	DeleteLotacao(string) error
	ListLotacoes(map[string]interface{}, int, int) ([]*Lotacao, int64, error)

	// Cargo methods
	CreateCargo(*Cargo) error
	GetCargoByID(string) (*Cargo, error)
	UpdateCargo(string, map[string]interface{}) error
	DeleteCargo(string) error
	ListCargos(map[string]interface{}, int, int) ([]*Cargo, int64, error)
}

type Service struct {
	repo   Repository
	logger *configs.Logger
}

func NewService(repo Repository, logger *configs.Logger) *Service {
	return &Service{
		repo:   repo,
		logger: logger,
	}
}

// Empresa Service Methods
func (s *Service) CreateEmpresa(req *CreateEmpresaRequest) (*EmpresaResponse, error) {
	// Validate CNPJ format
	if !s.isValidCNPJ(req.CNPJ) {
		return nil, ErrInvalidCNPJ
	}

	// Check if empresa already exists
	existing, err := s.repo.GetEmpresaByCNPJ(req.CNPJ)
	if err == nil && existing != nil {
		return nil, ErrEmpresaAlreadyExists
	}

	// Create empresa entity
	e := &Empresa{
		CNPJ:                    req.CNPJ,
		RazaoSocial:             req.RazaoSocial,
		NomeFantasia:            req.NomeFantasia,
		TipoInscricao:           req.TipoInscricao,
		ClassificacaoTributaria: req.ClassificacaoTributaria,
		NaturezaJuridica:        req.NaturezaJuridica,
		IndCooperativa:          req.IndCooperativa,
		IndConstrutora:          req.IndConstrutora,
		IndDesoneracao:          req.IndDesoneracao,
		Situacao:                "ATIVA",
		UsuarioCadastro:         req.UsuarioCadastro,
		Telefone:                req.Telefone,
		Email:                   req.Email,
		Logradouro:              req.Logradouro,
		Numero:                  req.Numero,
		Complemento:             req.Complemento,
		Bairro:                  req.Bairro,
		Cidade:                  req.Cidade,
		UF:                      req.UF,
		CEP:                     req.CEP,
	}

	err = s.repo.CreateEmpresa(e)
	if err != nil {
		s.logger.Error("Failed to create empresa", "error", err, "cnpj", req.CNPJ)
		return nil, err
	}

	s.logger.Info("Empresa created successfully", "id", e.ID, "cnpj", req.CNPJ)
	return s.mapEmpresaToResponse(e), nil
}

func (s *Service) GetEmpresaByID(id string) (*EmpresaResponse, error) {
	e, err := s.repo.GetEmpresaByID(id)
	if err != nil {
		s.logger.Error("Failed to get empresa", "error", err, "id", id)
		return nil, err
	}
	if e == nil {
		return nil, ErrEmpresaNotFound
	}
	return s.mapEmpresaToResponse(e), nil
}

func (s *Service) UpdateEmpresa(id string, req *UpdateEmpresaRequest) (*EmpresaResponse, error) {
	// Check if empresa exists
	existing, err := s.repo.GetEmpresaByID(id)
	if err != nil {
		return nil, err
	}
	if existing == nil {
		return nil, ErrEmpresaNotFound
	}

	// Build updates map
	updates := utils.MapFromArgs(
		"razao_social", req.RazaoSocial,
		"nome_fantasia", req.NomeFantasia,
		"tipo_inscricao", req.TipoInscricao,
		"classificacao_tributaria", req.ClassificacaoTributaria,
		"natureza_juridica", req.NaturezaJuridica,
		"ind_cooperativa", req.IndCooperativa,
		"ind_construtora", req.IndConstrutora,
		"ind_desoneracao", req.IndDesoneracao,
		"situacao", req.Situacao,
		"telefone", req.Telefone,
		"email", req.Email,
		"logradouro", req.Logradouro,
		"numero", req.Numero,
		"complemento", req.Complemento,
		"bairro", req.Bairro,
		"cidade", req.Cidade,
		"uf", req.UF,
		"cep", req.CEP,
		"protocolo_esocial", req.ProtocoloESocial,
		"recibo_esocial", req.ReciboESocial,
		"data_envio_esocial", req.DataEnvioESocial,
		"status_esocial", req.StatusESocial,
	)

	err = s.repo.UpdateEmpresa(id, updates)
	if err != nil {
		s.logger.Error("Failed to update empresa", "error", err, "id", id)
		return nil, err
	}

	// Get updated empresa
	updated, err := s.repo.GetEmpresaByID(id)
	if err != nil {
		return nil, err
	}

	s.logger.Info("Empresa updated successfully", "id", id)
	return s.mapEmpresaToResponse(updated), nil
}

func (s *Service) DeleteEmpresa(id string) error {
	// Check if empresa exists
	existing, err := s.repo.GetEmpresaByID(id)
	if err != nil {
		return err
	}
	if existing == nil {
		return ErrEmpresaNotFound
	}

	err = s.repo.DeleteEmpresa(id)
	if err != nil {
		s.logger.Error("Failed to delete empresa", "error", err, "id", id)
		return err
	}

	s.logger.Info("Empresa deleted successfully", "id", id)
	return nil
}

func (s *Service) ListEmpresas(req *ListRequest) (*ListResponse[EmpresaResponse], error) {
	page := 1
	limit := 10
	if req.Page > 0 {
		page = req.Page
	}
	if req.Limit > 0 && req.Limit <= 100 {
		limit = req.Limit
	}
	offset := (page - 1) * limit

	filter := utils.MapFromArgs(
		"empresa_id", req.EmpresaID,
		"search", req.Search,
		"situacao", req.Ativo, // reusing Ativo for situacao
	)

	empresas, total, err := s.repo.ListEmpresas(filter, limit, offset)
	if err != nil {
		s.logger.Error("Failed to list empresas", "error", err)
		return nil, err
	}

	responses := make([]EmpresaResponse, len(empresas))
	for i, e := range empresas {
		responses[i] = *s.mapEmpresaToResponse(e)
	}

	totalPages := (int(total) + limit - 1) / limit

	return &ListResponse[EmpresaResponse]{
		Data:       responses,
		Total:      total,
		Page:       page,
		Limit:      limit,
		TotalPages: totalPages,
	}, nil
}

// Estabelecimento Service Methods
func (s *Service) CreateEstabelecimento(req *CreateEstabelecimentoRequest) (*EstabelecimentoResponse, error) {
	// Check if empresa exists
	_, err := s.repo.GetEmpresaByID(req.EmpresaID)
	if err != nil {
		return nil, err
	}

	e := &Estabelecimento{
		EmpresaID:            req.EmpresaID,
		TipoInscricao:        req.TipoInscricao,
		NumeroInscricao:      req.NumeroInscricao,
		CNAEPrincipal:        req.CNAEPrincipal,
		CNAEPreparatorio:     req.CNAEPreparatorio,
		AlvaraFuncionamento:  req.AlvaraFuncionamento,
		DataInicioAtividades: req.DataInicioAtividades,
		IndObra:              req.IndObra,
		Logradouro:           req.Logradouro,
		Numero:               req.Numero,
		Complemento:          req.Complemento,
		Bairro:               req.Bairro,
		Cidade:               req.Cidade,
		UF:                   req.UF,
		CEP:                  req.CEP,
		Situacao:             "ATIVO",
	}

	err = s.repo.CreateEstabelecimento(e)
	if err != nil {
		s.logger.Error("Failed to create estabelecimento", "error", err, "empresa_id", req.EmpresaID)
		return nil, err
	}

	s.logger.Info("Estabelecimento created successfully", "id", e.ID)
	return s.mapEstabelecimentoToResponse(e), nil
}

func (s *Service) GetEstabelecimentoByID(id string) (*EstabelecimentoResponse, error) {
	e, err := s.repo.GetEstabelecimentoByID(id)
	if err != nil {
		s.logger.Error("Failed to get estabelecimento", "error", err, "id", id)
		return nil, err
	}
	if e == nil {
		return nil, ErrEstabelecimentoNotFound
	}
	return s.mapEstabelecimentoToResponse(e), nil
}

func (s *Service) UpdateEstabelecimento(id string, req *UpdateEstabelecimentoRequest) (*EstabelecimentoResponse, error) {
	existing, err := s.repo.GetEstabelecimentoByID(id)
	if err != nil {
		return nil, err
	}
	if existing == nil {
		return nil, ErrEstabelecimentoNotFound
	}

	updates := utils.MapFromArgs(
		"tipo_inscricao", req.TipoInscricao,
		"numero_inscricao", req.NumeroInscricao,
		"cnae_principal", req.CNAEPrincipal,
		"cnae_preparatorio", req.CNAEPreparatorio,
		"alvara_funcionamento", req.AlvaraFuncionamento,
		"data_inicio_atividades", req.DataInicioAtividades,
		"ind_obra", req.IndObra,
		"logradouro", req.Logradouro,
		"numero", req.Numero,
		"complemento", req.Complemento,
		"bairro", req.Bairro,
		"cidade", req.Cidade,
		"uf", req.UF,
		"cep", req.CEP,
		"situacao", req.Situacao,
	)

	err = s.repo.UpdateEstabelecimento(id, updates)
	if err != nil {
		s.logger.Error("Failed to update estabelecimento", "error", err, "id", id)
		return nil, err
	}

	updated, err := s.repo.GetEstabelecimentoByID(id)
	if err != nil {
		return nil, err
	}

	s.logger.Info("Estabelecimento updated successfully", "id", id)
	return s.mapEstabelecimentoToResponse(updated), nil
}

func (s *Service) DeleteEstabelecimento(id string) error {
	existing, err := s.repo.GetEstabelecimentoByID(id)
	if err != nil {
		return err
	}
	if existing == nil {
		return ErrEstabelecimentoNotFound
	}

	err = s.repo.DeleteEstabelecimento(id)
	if err != nil {
		s.logger.Error("Failed to delete estabelecimento", "error", err, "id", id)
		return err
	}

	s.logger.Info("Estabelecimento deleted successfully", "id", id)
	return nil
}

func (s *Service) ListEstabelecimentos(req *ListRequest) (*ListResponse[EstabelecimentoResponse], error) {
	page := 1
	limit := 10
	if req.Page > 0 {
		page = req.Page
	}
	if req.Limit > 0 && req.Limit <= 100 {
		limit = req.Limit
	}
	offset := (page - 1) * limit

	filter := utils.MapFromArgs(
		"empresa_id", req.EmpresaID,
		"search", req.Search,
		"situacao", req.Ativo,
	)

	estabelecimentos, total, err := s.repo.ListEstabelecimentos(filter, limit, offset)
	if err != nil {
		s.logger.Error("Failed to list estabelecimentos", "error", err)
		return nil, err
	}

	responses := make([]EstabelecimentoResponse, len(estabelecimentos))
	for i, e := range estabelecimentos {
		responses[i] = *s.mapEstabelecimentoToResponse(e)
	}

	totalPages := (int(total) + limit - 1) / limit

	return &ListResponse[EstabelecimentoResponse]{
		Data:       responses,
		Total:      total,
		Page:       page,
		Limit:      limit,
		TotalPages: totalPages,
	}, nil
}

// Rubrica Service Methods
func (s *Service) CreateRubrica(req *CreateRubricaRequest) (*RubricaResponse, error) {
	_, err := s.repo.GetEmpresaByID(req.EmpresaID)
	if err != nil {
		return nil, err
	}

	rb := &Rubrica{
		EmpresaID:          req.EmpresaID,
		Codigo:             req.Codigo,
		Descricao:          req.Descricao,
		NaturezaRubrica:    req.NaturezaRubrica,
		TipoRubrica:        req.TipoRubrica,
		CodINCCP:           req.CodINCCP,
		CodINCIRRF:         req.CodINCIRRF,
		CodINCFGTS:         req.CodINCFGTS,
		CodINCSind:         req.CodINCSind,
		Ativa:              true,
		DataInicioValidade: req.DataInicioValidade,
		DataFimValidade:    req.DataFimValidade,
	}

	err = s.repo.CreateRubrica(rb)
	if err != nil {
		s.logger.Error("Failed to create rubrica", "error", err, "empresa_id", req.EmpresaID)
		return nil, err
	}

	s.logger.Info("Rubrica created successfully", "id", rb.ID)
	return s.mapRubricaToResponse(rb), nil
}

func (s *Service) GetRubricaByID(id string) (*RubricaResponse, error) {
	rb, err := s.repo.GetRubricaByID(id)
	if err != nil {
		s.logger.Error("Failed to get rubrica", "error", err, "id", id)
		return nil, err
	}
	if rb == nil {
		return nil, ErrRubricaNotFound
	}
	return s.mapRubricaToResponse(rb), nil
}

func (s *Service) UpdateRubrica(id string, req *UpdateRubricaRequest) (*RubricaResponse, error) {
	existing, err := s.repo.GetRubricaByID(id)
	if err != nil {
		return nil, err
	}
	if existing == nil {
		return nil, ErrRubricaNotFound
	}

	updates := utils.MapFromArgs(
		"codigo", req.Codigo,
		"descricao", req.Descricao,
		"natureza_rubrica", req.NaturezaRubrica,
		"tipo_rubrica", req.TipoRubrica,
		"cod_inccp", req.CodINCCP,
		"cod_incirrf", req.CodINCIRRF,
		"cod_incfgts", req.CodINCFGTS,
		"cod_incsind", req.CodINCSind,
		"ativa", req.Ativa,
		"data_inicio_validade", req.DataInicioValidade,
		"data_fim_validade", req.DataFimValidade,
	)

	err = s.repo.UpdateRubrica(id, updates)
	if err != nil {
		s.logger.Error("Failed to update rubrica", "error", err, "id", id)
		return nil, err
	}

	updated, err := s.repo.GetRubricaByID(id)
	if err != nil {
		return nil, err
	}

	s.logger.Info("Rubrica updated successfully", "id", id)
	return s.mapRubricaToResponse(updated), nil
}

func (s *Service) DeleteRubrica(id string) error {
	existing, err := s.repo.GetRubricaByID(id)
	if err != nil {
		return err
	}
	if existing == nil {
		return ErrRubricaNotFound
	}

	err = s.repo.DeleteRubrica(id)
	if err != nil {
		s.logger.Error("Failed to delete rubrica", "error", err, "id", id)
		return err
	}

	s.logger.Info("Rubrica deleted successfully", "id", id)
	return nil
}

func (s *Service) ListRubricas(req *ListRequest) (*ListResponse[RubricaResponse], error) {
	page := 1
	limit := 10
	if req.Page > 0 {
		page = req.Page
	}
	if req.Limit > 0 && req.Limit <= 100 {
		limit = req.Limit
	}
	offset := (page - 1) * limit

	filter := utils.MapFromArgs(
		"empresa_id", req.EmpresaID,
		"search", req.Search,
		"ativa", req.Ativo,
	)

	rubricas, total, err := s.repo.ListRubricas(filter, limit, offset)
	if err != nil {
		s.logger.Error("Failed to list rubricas", "error", err)
		return nil, err
	}

	responses := make([]RubricaResponse, len(rubricas))
	for i, rb := range rubricas {
		responses[i] = *s.mapRubricaToResponse(rb)
	}

	totalPages := (int(total) + limit - 1) / limit

	return &ListResponse[RubricaResponse]{
		Data:       responses,
		Total:      total,
		Page:       page,
		Limit:      limit,
		TotalPages: totalPages,
	}, nil
}

// Lotacao Service Methods
func (s *Service) CreateLotacao(req *CreateLotacaoRequest) (*LotacaoResponse, error) {
	_, err := s.repo.GetEmpresaByID(req.EmpresaID)
	if err != nil {
		return nil, err
	}

	l := &Lotacao{
		EmpresaID:          req.EmpresaID,
		Codigo:             req.Codigo,
		Descricao:          req.Descricao,
		TipoLotacao:        req.TipoLotacao,
		FPAS:               req.FPAS,
		CodTerceiros:       req.CodTerceiros,
		Ativa:              true,
		DataInicioValidade: req.DataInicioValidade,
		DataFimValidade:    req.DataFimValidade,
	}

	err = s.repo.CreateLotacao(l)
	if err != nil {
		s.logger.Error("Failed to create lotacao", "error", err, "empresa_id", req.EmpresaID)
		return nil, err
	}

	s.logger.Info("Lotacao created successfully", "id", l.ID)
	return s.mapLotacaoToResponse(l), nil
}

func (s *Service) GetLotacaoByID(id string) (*LotacaoResponse, error) {
	l, err := s.repo.GetLotacaoByID(id)
	if err != nil {
		s.logger.Error("Failed to get lotacao", "error", err, "id", id)
		return nil, err
	}
	if l == nil {
		return nil, ErrLotacaoNotFound
	}
	return s.mapLotacaoToResponse(l), nil
}

func (s *Service) UpdateLotacao(id string, req *UpdateLotacaoRequest) (*LotacaoResponse, error) {
	existing, err := s.repo.GetLotacaoByID(id)
	if err != nil {
		return nil, err
	}
	if existing == nil {
		return nil, ErrLotacaoNotFound
	}

	updates := utils.MapFromArgs(
		"codigo", req.Codigo,
		"descricao", req.Descricao,
		"tipo_lotacao", req.TipoLotacao,
		"fpas", req.FPAS,
		"cod_terceiros", req.CodTerceiros,
		"ativa", req.Ativa,
		"data_inicio_validade", req.DataInicioValidade,
		"data_fim_validade", req.DataFimValidade,
	)

	err = s.repo.UpdateLotacao(id, updates)
	if err != nil {
		s.logger.Error("Failed to update lotacao", "error", err, "id", id)
		return nil, err
	}

	updated, err := s.repo.GetLotacaoByID(id)
	if err != nil {
		return nil, err
	}

	s.logger.Info("Lotacao updated successfully", "id", id)
	return s.mapLotacaoToResponse(updated), nil
}

func (s *Service) DeleteLotacao(id string) error {
	existing, err := s.repo.GetLotacaoByID(id)
	if err != nil {
		return err
	}
	if existing == nil {
		return ErrLotacaoNotFound
	}

	err = s.repo.DeleteLotacao(id)
	if err != nil {
		s.logger.Error("Failed to delete lotacao", "error", err, "id", id)
		return err
	}

	s.logger.Info("Lotacao deleted successfully", "id", id)
	return nil
}

func (s *Service) ListLotacoes(req *ListRequest) (*ListResponse[LotacaoResponse], error) {
	page := 1
	limit := 10
	if req.Page > 0 {
		page = req.Page
	}
	if req.Limit > 0 && req.Limit <= 100 {
		limit = req.Limit
	}
	offset := (page - 1) * limit

	filter := utils.MapFromArgs(
		"empresa_id", req.EmpresaID,
		"search", req.Search,
		"ativa", req.Ativo,
	)

	lotacoes, total, err := s.repo.ListLotacoes(filter, limit, offset)
	if err != nil {
		s.logger.Error("Failed to list lotacoes", "error", err)
		return nil, err
	}

	responses := make([]LotacaoResponse, len(lotacoes))
	for i, l := range lotacoes {
		responses[i] = *s.mapLotacaoToResponse(l)
	}

	totalPages := (int(total) + limit - 1) / limit

	return &ListResponse[LotacaoResponse]{
		Data:       responses,
		Total:      total,
		Page:       page,
		Limit:      limit,
		TotalPages: totalPages,
	}, nil
}

// Cargo Service Methods
func (s *Service) CreateCargo(req *CreateCargoRequest) (*CargoResponse, error) {
	_, err := s.repo.GetEmpresaByID(req.EmpresaID)
	if err != nil {
		return nil, err
	}

	c := &Cargo{
		EmpresaID:          req.EmpresaID,
		Codigo:             req.Codigo,
		Descricao:          req.Descricao,
		CBO:                req.CBO,
		Ativo:              true,
		DataInicioValidade: req.DataInicioValidade,
		DataFimValidade:    req.DataFimValidade,
	}

	err = s.repo.CreateCargo(c)
	if err != nil {
		s.logger.Error("Failed to create cargo", "error", err, "empresa_id", req.EmpresaID)
		return nil, err
	}

	s.logger.Info("Cargo created successfully", "id", c.ID)
	return s.mapCargoToResponse(c), nil
}

func (s *Service) GetCargoByID(id string) (*CargoResponse, error) {
	c, err := s.repo.GetCargoByID(id)
	if err != nil {
		s.logger.Error("Failed to get cargo", "error", err, "id", id)
		return nil, err
	}
	if c == nil {
		return nil, ErrCargoNotFound
	}
	return s.mapCargoToResponse(c), nil
}

func (s *Service) UpdateCargo(id string, req *UpdateCargoRequest) (*CargoResponse, error) {
	existing, err := s.repo.GetCargoByID(id)
	if err != nil {
		return nil, err
	}
	if existing == nil {
		return nil, ErrCargoNotFound
	}

	updates := utils.MapFromArgs(
		"codigo", req.Codigo,
		"descricao", req.Descricao,
		"cbo", req.CBO,
		"ativo", req.Ativo,
		"data_inicio_validade", req.DataInicioValidade,
		"data_fim_validade", req.DataFimValidade,
	)

	err = s.repo.UpdateCargo(id, updates)
	if err != nil {
		s.logger.Error("Failed to update cargo", "error", err, "id", id)
		return nil, err
	}

	updated, err := s.repo.GetCargoByID(id)
	if err != nil {
		return nil, err
	}

	s.logger.Info("Cargo updated successfully", "id", id)
	return s.mapCargoToResponse(updated), nil
}

func (s *Service) DeleteCargo(id string) error {
	existing, err := s.repo.GetCargoByID(id)
	if err != nil {
		return err
	}
	if existing == nil {
		return ErrCargoNotFound
	}

	err = s.repo.DeleteCargo(id)
	if err != nil {
		s.logger.Error("Failed to delete cargo", "error", err, "id", id)
		return err
	}

	s.logger.Info("Cargo deleted successfully", "id", id)
	return nil
}

func (s *Service) ListCargos(req *ListRequest) (*ListResponse[CargoResponse], error) {
	page := 1
	limit := 10
	if req.Page > 0 {
		page = req.Page
	}
	if req.Limit > 0 && req.Limit <= 100 {
		limit = req.Limit
	}
	offset := (page - 1) * limit

	filter := utils.MapFromArgs(
		"empresa_id", req.EmpresaID,
		"search", req.Search,
		"ativo", req.Ativo,
	)

	cargos, total, err := s.repo.ListCargos(filter, limit, offset)
	if err != nil {
		s.logger.Error("Failed to list cargos", "error", err)
		return nil, err
	}

	responses := make([]CargoResponse, len(cargos))
	for i, c := range cargos {
		responses[i] = *s.mapCargoToResponse(c)
	}

	totalPages := (int(total) + limit - 1) / limit

	return &ListResponse[CargoResponse]{
		Data:       responses,
		Total:      total,
		Page:       page,
		Limit:      limit,
		TotalPages: totalPages,
	}, nil
}

// Helper methods
func (s *Service) isValidCNPJ(cnpj string) bool {
	// Basic CNPJ validation - should be 14 digits
	if len(cnpj) != 14 {
		return false
	}
	for _, char := range cnpj {
		if char < '0' || char > '9' {
			return false
		}
	}
	return true
}

func (s *Service) mapEmpresaToResponse(e *Empresa) *EmpresaResponse {
	return &EmpresaResponse{
		ID:                      e.ID,
		CNPJ:                    e.CNPJ,
		RazaoSocial:             e.RazaoSocial,
		NomeFantasia:            e.NomeFantasia,
		TipoInscricao:           e.TipoInscricao,
		ClassificacaoTributaria: e.ClassificacaoTributaria,
		NaturezaJuridica:        e.NaturezaJuridica,
		IndCooperativa:          e.IndCooperativa,
		IndConstrutora:          e.IndConstrutora,
		IndDesoneracao:          e.IndDesoneracao,
		Situacao:                e.Situacao,
		DataCadastro:            e.DataCadastro,
		DataAtualizacao:         e.DataAtualizacao,
		UsuarioCadastro:         e.UsuarioCadastro,
		Telefone:                e.Telefone,
		Email:                   e.Email,
		Logradouro:              e.Logradouro,
		Numero:                  e.Numero,
		Complemento:             e.Complemento,
		Bairro:                  e.Bairro,
		Cidade:                  e.Cidade,
		UF:                      e.UF,
		CEP:                     e.CEP,
		ProtocoloESocial:        e.ProtocoloESocial,
		ReciboESocial:           e.ReciboESocial,
		DataEnvioESocial:        e.DataEnvioESocial,
		StatusESocial:           e.StatusESocial,
	}
}

func (s *Service) mapEstabelecimentoToResponse(e *Estabelecimento) *EstabelecimentoResponse {
	return &EstabelecimentoResponse{
		ID:                   e.ID,
		EmpresaID:            e.EmpresaID,
		TipoInscricao:        e.TipoInscricao,
		NumeroInscricao:      e.NumeroInscricao,
		CNAEPrincipal:        e.CNAEPrincipal,
		CNAEPreparatorio:     e.CNAEPreparatorio,
		AlvaraFuncionamento:  e.AlvaraFuncionamento,
		DataInicioAtividades: e.DataInicioAtividades,
		IndObra:              e.IndObra,
		Logradouro:           e.Logradouro,
		Numero:               e.Numero,
		Complemento:          e.Complemento,
		Bairro:               e.Bairro,
		Cidade:               e.Cidade,
		UF:                   e.UF,
		CEP:                  e.CEP,
		Situacao:             e.Situacao,
		DataCadastro:         e.DataCadastro,
	}
}

func (s *Service) mapRubricaToResponse(rb *Rubrica) *RubricaResponse {
	return &RubricaResponse{
		ID:                 rb.ID,
		EmpresaID:          rb.EmpresaID,
		Codigo:             rb.Codigo,
		Descricao:          rb.Descricao,
		NaturezaRubrica:    rb.NaturezaRubrica,
		TipoRubrica:        rb.TipoRubrica,
		CodINCCP:           rb.CodINCCP,
		CodINCIRRF:         rb.CodINCIRRF,
		CodINCFGTS:         rb.CodINCFGTS,
		CodINCSind:         rb.CodINCSind,
		Ativa:              rb.Ativa,
		DataCadastro:       rb.DataCadastro,
		DataInicioValidade: rb.DataInicioValidade,
		DataFimValidade:    rb.DataFimValidade,
	}
}

func (s *Service) mapLotacaoToResponse(l *Lotacao) *LotacaoResponse {
	return &LotacaoResponse{
		ID:                 l.ID,
		EmpresaID:          l.EmpresaID,
		Codigo:             l.Codigo,
		Descricao:          l.Descricao,
		TipoLotacao:        l.TipoLotacao,
		FPAS:               l.FPAS,
		CodTerceiros:       l.CodTerceiros,
		Ativa:              l.Ativa,
		DataCadastro:       l.DataCadastro,
		DataInicioValidade: l.DataInicioValidade,
		DataFimValidade:    l.DataFimValidade,
	}
}

func (s *Service) mapCargoToResponse(c *Cargo) *CargoResponse {
	return &CargoResponse{
		ID:                 c.ID,
		EmpresaID:          c.EmpresaID,
		Codigo:             c.Codigo,
		Descricao:          c.Descricao,
		CBO:                c.CBO,
		Ativo:              c.Ativo,
		DataCadastro:       c.DataCadastro,
		DataInicioValidade: c.DataInicioValidade,
		DataFimValidade:    c.DataFimValidade,
	}
}
