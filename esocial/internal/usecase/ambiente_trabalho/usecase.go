package ambiente_trabalho

import (
	"github.com/avelinobego/esocial/configs"
	"github.com/avelinobego/esocial/internal/domain/ambiente_trabalho"
)

type Usecase struct {
	service *ambiente_trabalho.Service
	logger  *configs.Logger
}

func NewUsecase(service *ambiente_trabalho.Service, logger *configs.Logger) *Usecase {
	return &Usecase{
		service: service,
		logger:  logger,
	}
}

// Horario Usecase Methods
func (u *Usecase) CreateHorario(req *ambiente_trabalho.CreateHorarioRequest) (*ambiente_trabalho.HorarioResponse, error) {
	u.logger.Info("Creating horario", "empresa_id", req.EmpresaID, "codigo", req.Codigo)
	return u.service.CreateHorario(req)
}

func (u *Usecase) GetHorarioByID(id string) (*ambiente_trabalho.HorarioResponse, error) {
	u.logger.Info("Getting horario by ID", "id", id)
	return u.service.GetHorarioByID(id)
}

func (u *Usecase) UpdateHorario(id string, req *ambiente_trabalho.UpdateHorarioRequest) (*ambiente_trabalho.HorarioResponse, error) {
	u.logger.Info("Updating horario", "id", id)
	return u.service.UpdateHorario(id, req)
}

func (u *Usecase) DeleteHorario(id string) error {
	u.logger.Info("Deleting horario", "id", id)
	return u.service.DeleteHorario(id)
}

func (u *Usecase) ListHorarios(req *ambiente_trabalho.ListRequest) (*ambiente_trabalho.ListResponse[ambiente_trabalho.HorarioResponse], error) {
	u.logger.Info("Listing horarios", "page", req.Page, "limit", req.Limit)
	return u.service.ListHorarios(req)
}

// HorarioDetalhe Usecase Methods
func (u *Usecase) CreateHorarioDetalhe(req *ambiente_trabalho.CreateHorarioDetalheRequest) (*ambiente_trabalho.HorarioDetalheResponse, error) {
	u.logger.Info("Creating horario detalhe", "horario_id", req.HorarioID, "dia_semana", req.DiaSemana)
	return u.service.CreateHorarioDetalhe(req)
}

func (u *Usecase) UpdateHorarioDetalhe(id string, req *ambiente_trabalho.UpdateHorarioDetalheRequest) (*ambiente_trabalho.HorarioDetalheResponse, error) {
	u.logger.Info("Updating horario detalhe", "id", id)
	return u.service.UpdateHorarioDetalhe(id, req)
}

func (u *Usecase) DeleteHorarioDetalhe(id string) error {
	u.logger.Info("Deleting horario detalhe", "id", id)
	return u.service.DeleteHorarioDetalhe(id)
}

// AmbienteTrabalho Usecase Methods
func (u *Usecase) CreateAmbienteTrabalho(req *ambiente_trabalho.CreateAmbienteTrabalhoRequest) (*ambiente_trabalho.AmbienteTrabalhoResponse, error) {
	u.logger.Info("Creating ambiente trabalho", "empresa_id", req.EmpresaID, "codigo", req.Codigo)
	return u.service.CreateAmbienteTrabalho(req)
}

func (u *Usecase) GetAmbienteTrabalhoByID(id string) (*ambiente_trabalho.AmbienteTrabalhoResponse, error) {
	u.logger.Info("Getting ambiente trabalho by ID", "id", id)
	return u.service.GetAmbienteTrabalhoByID(id)
}

func (u *Usecase) UpdateAmbienteTrabalho(id string, req *ambiente_trabalho.UpdateAmbienteTrabalhoRequest) (*ambiente_trabalho.AmbienteTrabalhoResponse, error) {
	u.logger.Info("Updating ambiente trabalho", "id", id)
	return u.service.UpdateAmbienteTrabalho(id, req)
}

func (u *Usecase) DeleteAmbienteTrabalho(id string) error {
	u.logger.Info("Deleting ambiente trabalho", "id", id)
	return u.service.DeleteAmbienteTrabalho(id)
}

func (u *Usecase) ListAmbientesTrabalho(req *ambiente_trabalho.ListRequest) (*ambiente_trabalho.ListResponse[ambiente_trabalho.AmbienteTrabalhoResponse], error) {
	u.logger.Info("Listing ambientes trabalho", "page", req.Page, "limit", req.Limit)
	return u.service.ListAmbientesTrabalho(req)
}
