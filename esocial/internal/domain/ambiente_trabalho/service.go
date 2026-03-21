package ambiente_trabalho

import (
	"errors"

	"github.com/avelinobego/esocial/configs"
	"github.com/avelinobego/esocial/pkg/utils"
)

var (
	ErrHorarioNotFound             = errors.New("horario not found")
	ErrHorarioDetalheNotFound      = errors.New("horario detalhe not found")
	ErrAmbienteTrabalhoNotFound    = errors.New("ambiente trabalho not found")
	ErrHorarioDetalheAlreadyExists = errors.New("horario detalhe already exists for this day")
)

type Repository interface {
	// Horario methods
	CreateHorario(*Horario) error
	GetHorarioByID(string) (*Horario, error)
	UpdateHorario(string, map[string]interface{}) error
	DeleteHorario(string) error
	ListHorarios(map[string]interface{}, int, int) ([]*Horario, int64, error)

	// HorarioDetalhe methods
	CreateHorarioDetalhe(*HorarioDetalhe) error
	GetHorarioDetalheByID(string) (*HorarioDetalhe, error)
	GetHorarioDetalhesByHorarioID(string) ([]*HorarioDetalhe, error)
	UpdateHorarioDetalhe(string, map[string]interface{}) error
	DeleteHorarioDetalhe(string) error
	DeleteHorarioDetalhesByHorarioID(string) error

	// AmbienteTrabalho methods
	CreateAmbienteTrabalho(*AmbienteTrabalho) error
	GetAmbienteTrabalhoByID(string) (*AmbienteTrabalho, error)
	UpdateAmbienteTrabalho(string, map[string]interface{}) error
	DeleteAmbienteTrabalho(string) error
	ListAmbientesTrabalho(map[string]interface{}, int, int) ([]*AmbienteTrabalho, int64, error)
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

// Horario Service Methods
func (s *Service) CreateHorario(req *CreateHorarioRequest) (*HorarioResponse, error) {
	h := &Horario{
		EmpresaID:             req.EmpresaID,
		Codigo:                req.Codigo,
		Descricao:             req.Descricao,
		TipoJornada:           req.TipoJornada,
		DuracaoJornadaSemanal: req.DuracaoJornadaSemanal,
		TipoIntervalo:         req.TipoIntervalo,
		DuracaoIntervalo:      req.DuracaoIntervalo,
		Ativo:                 true,
	}

	err := s.repo.CreateHorario(h)
	if err != nil {
		s.logger.Error("Failed to create horario", "error", err, "empresa_id", req.EmpresaID)
		return nil, err
	}

	s.logger.Info("Horario created successfully", "id", h.ID)
	return s.mapHorarioToResponse(h), nil
}

func (s *Service) GetHorarioByID(id string) (*HorarioResponse, error) {
	h, err := s.repo.GetHorarioByID(id)
	if err != nil {
		s.logger.Error("Failed to get horario", "error", err, "id", id)
		return nil, err
	}
	if h == nil {
		return nil, ErrHorarioNotFound
	}

	detalhes, err := s.repo.GetHorarioDetalhesByHorarioID(id)
	if err != nil {
		s.logger.Error("Failed to get horario detalhes", "error", err, "horario_id", id)
		return nil, err
	}

	response := s.mapHorarioToResponse(h)
	response.Detalhes = make([]HorarioDetalheResponse, len(detalhes))
	for i, d := range detalhes {
		response.Detalhes[i] = *s.mapHorarioDetalheToResponse(d)
	}

	return response, nil
}

func (s *Service) UpdateHorario(id string, req *UpdateHorarioRequest) (*HorarioResponse, error) {
	existing, err := s.repo.GetHorarioByID(id)
	if err != nil {
		return nil, err
	}
	if existing == nil {
		return nil, ErrHorarioNotFound
	}

	updates := utils.MapFromArgs(
		"codigo", req.Codigo,
		"descricao", req.Descricao,
		"tipo_jornada", req.TipoJornada,
		"duracao_jornada_semanal", req.DuracaoJornadaSemanal,
		"tipo_intervalo", req.TipoIntervalo,
		"duracao_intervalo", req.DuracaoIntervalo,
		"ativo", req.Ativo,
	)

	err = s.repo.UpdateHorario(id, updates)
	if err != nil {
		s.logger.Error("Failed to update horario", "error", err, "id", id)
		return nil, err
	}

	updated, err := s.repo.GetHorarioByID(id)
	if err != nil {
		return nil, err
	}

	s.logger.Info("Horario updated successfully", "id", id)
	return s.mapHorarioToResponse(updated), nil
}

func (s *Service) DeleteHorario(id string) error {
	existing, err := s.repo.GetHorarioByID(id)
	if err != nil {
		return err
	}
	if existing == nil {
		return ErrHorarioNotFound
	}

	// Delete all detalhes first
	err = s.repo.DeleteHorarioDetalhesByHorarioID(id)
	if err != nil {
		s.logger.Error("Failed to delete horario detalhes", "error", err, "horario_id", id)
		return err
	}

	err = s.repo.DeleteHorario(id)
	if err != nil {
		s.logger.Error("Failed to delete horario", "error", err, "id", id)
		return err
	}

	s.logger.Info("Horario deleted successfully", "id", id)
	return nil
}

func (s *Service) ListHorarios(req *ListRequest) (*ListResponse[HorarioResponse], error) {
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

	horarios, total, err := s.repo.ListHorarios(filter, limit, offset)
	if err != nil {
		s.logger.Error("Failed to list horarios", "error", err)
		return nil, err
	}

	responses := make([]HorarioResponse, len(horarios))
	for i, h := range horarios {
		responses[i] = *s.mapHorarioToResponse(h)
	}

	totalPages := (int(total) + limit - 1) / limit

	return &ListResponse[HorarioResponse]{
		Data:       responses,
		Total:      total,
		Page:       page,
		Limit:      limit,
		TotalPages: totalPages,
	}, nil
}

// HorarioDetalhe Service Methods
func (s *Service) CreateHorarioDetalhe(req *CreateHorarioDetalheRequest) (*HorarioDetalheResponse, error) {
	// Check if horario exists
	_, err := s.repo.GetHorarioByID(req.HorarioID)
	if err != nil {
		return nil, err
	}

	// Check if detalhe already exists for this day
	detalhes, err := s.repo.GetHorarioDetalhesByHorarioID(req.HorarioID)
	if err != nil {
		return nil, err
	}

	for _, d := range detalhes {
		if d.DiaSemana == req.DiaSemana {
			return nil, ErrHorarioDetalheAlreadyExists
		}
	}

	hd := &HorarioDetalhe{
		HorarioID:      req.HorarioID,
		DiaSemana:      req.DiaSemana,
		HorarioEntrada: req.HorarioEntrada,
		HorarioSaida:   req.HorarioSaida,
		DuracaoJornada: req.DuracaoJornada,
	}

	err = s.repo.CreateHorarioDetalhe(hd)
	if err != nil {
		s.logger.Error("Failed to create horario detalhe", "error", err, "horario_id", req.HorarioID)
		return nil, err
	}

	s.logger.Info("Horario detalhe created successfully", "id", hd.ID)
	return s.mapHorarioDetalheToResponse(hd), nil
}

func (s *Service) UpdateHorarioDetalhe(id string, req *UpdateHorarioDetalheRequest) (*HorarioDetalheResponse, error) {
	existing, err := s.repo.GetHorarioDetalheByID(id)
	if err != nil {
		return nil, err
	}
	if existing == nil {
		return nil, ErrHorarioDetalheNotFound
	}

	updates := utils.MapFromArgs(
		"dia_semana", req.DiaSemana,
		"horario_entrada", req.HorarioEntrada,
		"horario_saida", req.HorarioSaida,
		"duracao_jornada", req.DuracaoJornada,
	)

	err = s.repo.UpdateHorarioDetalhe(id, updates)
	if err != nil {
		s.logger.Error("Failed to update horario detalhe", "error", err, "id", id)
		return nil, err
	}

	updated, err := s.repo.GetHorarioDetalheByID(id)
	if err != nil {
		return nil, err
	}

	s.logger.Info("Horario detalhe updated successfully", "id", id)
	return s.mapHorarioDetalheToResponse(updated), nil
}

func (s *Service) DeleteHorarioDetalhe(id string) error {
	existing, err := s.repo.GetHorarioDetalheByID(id)
	if err != nil {
		return err
	}
	if existing == nil {
		return ErrHorarioDetalheNotFound
	}

	err = s.repo.DeleteHorarioDetalhe(id)
	if err != nil {
		s.logger.Error("Failed to delete horario detalhe", "error", err, "id", id)
		return err
	}

	s.logger.Info("Horario detalhe deleted successfully", "id", id)
	return nil
}

// AmbienteTrabalho Service Methods
func (s *Service) CreateAmbienteTrabalho(req *CreateAmbienteTrabalhoRequest) (*AmbienteTrabalhoResponse, error) {
	at := &AmbienteTrabalho{
		EmpresaID:       req.EmpresaID,
		Codigo:          req.Codigo,
		Descricao:       req.Descricao,
		LocalAmbiente:   req.LocalAmbiente,
		TipoInscricao:   req.TipoInscricao,
		NumeroInscricao: req.NumeroInscricao,
		Ativo:           true,
	}

	err := s.repo.CreateAmbienteTrabalho(at)
	if err != nil {
		s.logger.Error("Failed to create ambiente trabalho", "error", err, "empresa_id", req.EmpresaID)
		return nil, err
	}

	s.logger.Info("Ambiente trabalho created successfully", "id", at.ID)
	return s.mapAmbienteTrabalhoToResponse(at), nil
}

func (s *Service) GetAmbienteTrabalhoByID(id string) (*AmbienteTrabalhoResponse, error) {
	at, err := s.repo.GetAmbienteTrabalhoByID(id)
	if err != nil {
		s.logger.Error("Failed to get ambiente trabalho", "error", err, "id", id)
		return nil, err
	}
	if at == nil {
		return nil, ErrAmbienteTrabalhoNotFound
	}
	return s.mapAmbienteTrabalhoToResponse(at), nil
}

func (s *Service) UpdateAmbienteTrabalho(id string, req *UpdateAmbienteTrabalhoRequest) (*AmbienteTrabalhoResponse, error) {
	existing, err := s.repo.GetAmbienteTrabalhoByID(id)
	if err != nil {
		return nil, err
	}
	if existing == nil {
		return nil, ErrAmbienteTrabalhoNotFound
	}

	updates := utils.MapFromArgs(
		"codigo", req.Codigo,
		"descricao", req.Descricao,
		"local_ambiente", req.LocalAmbiente,
		"tipo_inscricao", req.TipoInscricao,
		"numero_inscricao", req.NumeroInscricao,
		"ativo", req.Ativo,
	)

	err = s.repo.UpdateAmbienteTrabalho(id, updates)
	if err != nil {
		s.logger.Error("Failed to update ambiente trabalho", "error", err, "id", id)
		return nil, err
	}

	updated, err := s.repo.GetAmbienteTrabalhoByID(id)
	if err != nil {
		return nil, err
	}

	s.logger.Info("Ambiente trabalho updated successfully", "id", id)
	return s.mapAmbienteTrabalhoToResponse(updated), nil
}

func (s *Service) DeleteAmbienteTrabalho(id string) error {
	existing, err := s.repo.GetAmbienteTrabalhoByID(id)
	if err != nil {
		return err
	}
	if existing == nil {
		return ErrAmbienteTrabalhoNotFound
	}

	err = s.repo.DeleteAmbienteTrabalho(id)
	if err != nil {
		s.logger.Error("Failed to delete ambiente trabalho", "error", err, "id", id)
		return err
	}

	s.logger.Info("Ambiente trabalho deleted successfully", "id", id)
	return nil
}

func (s *Service) ListAmbientesTrabalho(req *ListRequest) (*ListResponse[AmbienteTrabalhoResponse], error) {
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

	ambientes, total, err := s.repo.ListAmbientesTrabalho(filter, limit, offset)
	if err != nil {
		s.logger.Error("Failed to list ambientes trabalho", "error", err)
		return nil, err
	}

	responses := make([]AmbienteTrabalhoResponse, len(ambientes))
	for i, at := range ambientes {
		responses[i] = *s.mapAmbienteTrabalhoToResponse(at)
	}

	totalPages := (int(total) + limit - 1) / limit

	return &ListResponse[AmbienteTrabalhoResponse]{
		Data:       responses,
		Total:      total,
		Page:       page,
		Limit:      limit,
		TotalPages: totalPages,
	}, nil
}

// Helper methods
func (s *Service) mapHorarioToResponse(h *Horario) *HorarioResponse {
	return &HorarioResponse{
		ID:                    h.ID,
		EmpresaID:             h.EmpresaID,
		Codigo:                h.Codigo,
		Descricao:             h.Descricao,
		TipoJornada:           h.TipoJornada,
		DuracaoJornadaSemanal: h.DuracaoJornadaSemanal,
		TipoIntervalo:         h.TipoIntervalo,
		DuracaoIntervalo:      h.DuracaoIntervalo,
		Ativo:                 h.Ativo,
		DataCadastro:          h.DataCadastro,
	}
}

func (s *Service) mapHorarioDetalheToResponse(hd *HorarioDetalhe) *HorarioDetalheResponse {
	return &HorarioDetalheResponse{
		ID:             hd.ID,
		HorarioID:      hd.HorarioID,
		DiaSemana:      hd.DiaSemana,
		HorarioEntrada: hd.HorarioEntrada,
		HorarioSaida:   hd.HorarioSaida,
		DuracaoJornada: hd.DuracaoJornada,
	}
}

func (s *Service) mapAmbienteTrabalhoToResponse(at *AmbienteTrabalho) *AmbienteTrabalhoResponse {
	return &AmbienteTrabalhoResponse{
		ID:              at.ID,
		EmpresaID:       at.EmpresaID,
		Codigo:          at.Codigo,
		Descricao:       at.Descricao,
		LocalAmbiente:   at.LocalAmbiente,
		TipoInscricao:   at.TipoInscricao,
		NumeroInscricao: at.NumeroInscricao,
		Ativo:           at.Ativo,
		DataCadastro:    at.DataCadastro,
	}
}
