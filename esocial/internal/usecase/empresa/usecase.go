package empresa

import (
	"github.com/avelinobego/esocial/configs"
	"github.com/avelinobego/esocial/internal/domain/empresa"
)

type Usecase struct {
	service *empresa.Service
	logger  *configs.Logger
}

func NewUsecase(service *empresa.Service, logger *configs.Logger) *Usecase {
	return &Usecase{
		service: service,
		logger:  logger,
	}
}

// Empresa Usecase Methods

func (u *Usecase) CreateEmpresa(req *empresa.CreateEmpresaRequest) (*empresa.EmpresaResponse, error) {
	u.logger.Info("Creating empresa", "cnpj", req.CNPJ)
	return u.service.CreateEmpresa(req)
}

func (u *Usecase) GetEmpresaByID(id string) (*empresa.EmpresaResponse, error) {
	u.logger.Info("Getting empresa by ID", "id", id)
	return u.service.GetEmpresaByID(id)
}

func (u *Usecase) UpdateEmpresa(id string, req *empresa.UpdateEmpresaRequest) (*empresa.EmpresaResponse, error) {
	u.logger.Info("Updating empresa", "id", id)
	return u.service.UpdateEmpresa(id, req)
}

func (u *Usecase) DeleteEmpresa(id string) error {
	u.logger.Info("Deleting empresa", "id", id)
	return u.service.DeleteEmpresa(id)
}

func (u *Usecase) ListEmpresas(req *empresa.ListRequest) (*empresa.ListResponse[empresa.EmpresaResponse], error) {
	u.logger.Info("Listing empresas", "page", req.Page, "limit", req.Limit)
	return u.service.ListEmpresas(req)
}

// Estabelecimento Usecase Methods

func (u *Usecase) CreateEstabelecimento(req *empresa.CreateEstabelecimentoRequest) (*empresa.EstabelecimentoResponse, error) {
	u.logger.Info("Creating estabelecimento", "empresa_id", req.EmpresaID)
	return u.service.CreateEstabelecimento(req)
}

func (u *Usecase) GetEstabelecimentoByID(id string) (*empresa.EstabelecimentoResponse, error) {
	u.logger.Info("Getting estabelecimento by ID", "id", id)
	return u.service.GetEstabelecimentoByID(id)
}

func (u *Usecase) UpdateEstabelecimento(id string, req *empresa.UpdateEstabelecimentoRequest) (*empresa.EstabelecimentoResponse, error) {
	u.logger.Info("Updating estabelecimento", "id", id)
	return u.service.UpdateEstabelecimento(id, req)
}

func (u *Usecase) DeleteEstabelecimento(id string) error {
	u.logger.Info("Deleting estabelecimento", "id", id)
	return u.service.DeleteEstabelecimento(id)
}

func (u *Usecase) ListEstabelecimentos(req *empresa.ListRequest) (*empresa.ListResponse[empresa.EstabelecimentoResponse], error) {
	u.logger.Info("Listing estabelecimentos", "page", req.Page, "limit", req.Limit)
	return u.service.ListEstabelecimentos(req)
}

// Rubrica Usecase Methods

func (u *Usecase) CreateRubrica(req *empresa.CreateRubricaRequest) (*empresa.RubricaResponse, error) {
	u.logger.Info("Creating rubrica", "empresa_id", req.EmpresaID, "codigo", req.Codigo)
	return u.service.CreateRubrica(req)
}

func (u *Usecase) GetRubricaByID(id string) (*empresa.RubricaResponse, error) {
	u.logger.Info("Getting rubrica by ID", "id", id)
	return u.service.GetRubricaByID(id)
}

func (u *Usecase) UpdateRubrica(id string, req *empresa.UpdateRubricaRequest) (*empresa.RubricaResponse, error) {
	u.logger.Info("Updating rubrica", "id", id)
	return u.service.UpdateRubrica(id, req)
}

func (u *Usecase) DeleteRubrica(id string) error {
	u.logger.Info("Deleting rubrica", "id", id)
	return u.service.DeleteRubrica(id)
}

func (u *Usecase) ListRubricas(req *empresa.ListRequest) (*empresa.ListResponse[empresa.RubricaResponse], error) {
	u.logger.Info("Listing rubricas", "page", req.Page, "limit", req.Limit)
	return u.service.ListRubricas(req)
}

// Lotacao Usecase Methods

func (u *Usecase) CreateLotacao(req *empresa.CreateLotacaoRequest) (*empresa.LotacaoResponse, error) {
	u.logger.Info("Creating lotacao", "empresa_id", req.EmpresaID, "codigo", req.Codigo)
	return u.service.CreateLotacao(req)
}

func (u *Usecase) GetLotacaoByID(id string) (*empresa.LotacaoResponse, error) {
	u.logger.Info("Getting lotacao by ID", "id", id)
	return u.service.GetLotacaoByID(id)
}

func (u *Usecase) UpdateLotacao(id string, req *empresa.UpdateLotacaoRequest) (*empresa.LotacaoResponse, error) {
	u.logger.Info("Updating lotacao", "id", id)
	return u.service.UpdateLotacao(id, req)
}

func (u *Usecase) DeleteLotacao(id string) error {
	u.logger.Info("Deleting lotacao", "id", id)
	return u.service.DeleteLotacao(id)
}

func (u *Usecase) ListLotacoes(req *empresa.ListRequest) (*empresa.ListResponse[empresa.LotacaoResponse], error) {
	u.logger.Info("Listing lotacoes", "page", req.Page, "limit", req.Limit)
	return u.service.ListLotacoes(req)
}

// Cargo Usecase Methods

func (u *Usecase) CreateCargo(req *empresa.CreateCargoRequest) (*empresa.CargoResponse, error) {
	u.logger.Info("Creating cargo", "empresa_id", req.EmpresaID, "codigo", req.Codigo)
	return u.service.CreateCargo(req)
}

func (u *Usecase) GetCargoByID(id string) (*empresa.CargoResponse, error) {
	u.logger.Info("Getting cargo by ID", "id", id)
	return u.service.GetCargoByID(id)
}

func (u *Usecase) UpdateCargo(id string, req *empresa.UpdateCargoRequest) (*empresa.CargoResponse, error) {
	u.logger.Info("Updating cargo", "id", id)
	return u.service.UpdateCargo(id, req)
}

func (u *Usecase) DeleteCargo(id string) error {
	u.logger.Info("Deleting cargo", "id", id)
	return u.service.DeleteCargo(id)
}

func (u *Usecase) ListCargos(req *empresa.ListRequest) (*empresa.ListResponse[empresa.CargoResponse], error) {
	u.logger.Info("Listing cargos", "page", req.Page, "limit", req.Limit)
	return u.service.ListCargos(req)
}
