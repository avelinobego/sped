package empresa_test

import (
	"io"
	"testing"
	"time"

	config "github.com/avelinobego/esocial/configs"
	domain "github.com/avelinobego/esocial/internal/domain/empresa"
	usecase "github.com/avelinobego/esocial/internal/usecase/empresa"
)

type mockRepository struct {
	empresas map[string]*domain.Empresa
}

func (m *mockRepository) CreateEmpresa(e *domain.Empresa) error {
	if e.ID == "" {
		e.ID = "empresa-1"
	}
	m.empresas[e.ID] = e
	return nil
}

func (m *mockRepository) GetEmpresaByID(id string) (*domain.Empresa, error) {
	if e, ok := m.empresas[id]; ok {
		return e, nil
	}
	return nil, nil
}

func (m *mockRepository) GetEmpresaByCNPJ(cnpj string) (*domain.Empresa, error) {
	for _, e := range m.empresas {
		if e.CNPJ == cnpj {
			return e, nil
		}
	}
	return nil, nil
}

func (m *mockRepository) UpdateEmpresa(id string, updates map[string]interface{}) error { return nil }
func (m *mockRepository) DeleteEmpresa(id string) error                                 { delete(m.empresas, id); return nil }
func (m *mockRepository) ListEmpresas(filter map[string]interface{}, limit int, offset int) ([]*domain.Empresa, int64, error) {
	list := make([]*domain.Empresa, 0, len(m.empresas))
	for _, e := range m.empresas {
		list = append(list, e)
	}
	return list, int64(len(list)), nil
}
func (m *mockRepository) CreateEstabelecimento(*domain.Estabelecimento) error { return nil }
func (m *mockRepository) GetEstabelecimentoByID(string) (*domain.Estabelecimento, error) {
	return nil, nil
}
func (m *mockRepository) UpdateEstabelecimento(string, map[string]interface{}) error { return nil }
func (m *mockRepository) DeleteEstabelecimento(string) error                         { return nil }
func (m *mockRepository) ListEstabelecimentos(map[string]interface{}, int, int) ([]*domain.Estabelecimento, int64, error) {
	return nil, 0, nil
}
func (m *mockRepository) CreateRubrica(*domain.Rubrica) error                { return nil }
func (m *mockRepository) GetRubricaByID(string) (*domain.Rubrica, error)     { return nil, nil }
func (m *mockRepository) UpdateRubrica(string, map[string]interface{}) error { return nil }
func (m *mockRepository) DeleteRubrica(string) error                         { return nil }
func (m *mockRepository) ListRubricas(map[string]interface{}, int, int) ([]*domain.Rubrica, int64, error) {
	return nil, 0, nil
}
func (m *mockRepository) CreateLotacao(*domain.Lotacao) error                { return nil }
func (m *mockRepository) GetLotacaoByID(string) (*domain.Lotacao, error)     { return nil, nil }
func (m *mockRepository) UpdateLotacao(string, map[string]interface{}) error { return nil }
func (m *mockRepository) DeleteLotacao(string) error                         { return nil }
func (m *mockRepository) ListLotacoes(map[string]interface{}, int, int) ([]*domain.Lotacao, int64, error) {
	return nil, 0, nil
}
func (m *mockRepository) CreateCargo(*domain.Cargo) error                  { return nil }
func (m *mockRepository) GetCargoByID(string) (*domain.Cargo, error)       { return nil, nil }
func (m *mockRepository) UpdateCargo(string, map[string]interface{}) error { return nil }
func (m *mockRepository) DeleteCargo(string) error                         { return nil }
func (m *mockRepository) ListCargos(map[string]interface{}, int, int) ([]*domain.Cargo, int64, error) {
	return nil, 0, nil
}

func TestUsecaseDelegatesToService(t *testing.T) {
	logger := config.NewLogger(io.Discard)
	repo := &mockRepository{empresas: make(map[string]*domain.Empresa)}
	svc := domain.NewService(repo, logger)
	uc := usecase.NewUsecase(svc, logger)

	req := &domain.CreateEmpresaRequest{
		CNPJ:                    "12345678000195",
		RazaoSocial:             "Teste Societe",
		TipoInscricao:           "1",
		ClassificacaoTributaria: "01",
		IndCooperativa:          "N",
		IndConstrutora:          "N",
		IndDesoneracao:          "N",
	}

	created, err := uc.CreateEmpresa(req)
	if err != nil {
		t.Fatalf("CreateEmpresa failed: %v", err)
	}
	if created.CNPJ != req.CNPJ {
		t.Fatalf("unexpected cnpj: got %s", created.CNPJ)
	}

	fetched, err := uc.GetEmpresaByID(created.ID)
	if err != nil {
		t.Fatalf("GetEmpresaByID failed: %v", err)
	}
	if fetched == nil || fetched.ID != created.ID {
		t.Fatalf("unexpected returned empresa")
	}

	_, err = uc.GetEmpresaByID("not-found")
	if err != domain.ErrEmpresaNotFound {
		t.Fatalf("expected ErrEmpresaNotFound, got %v", err)
	}
}

func TestUsecaseListEmpresas(t *testing.T) {
	logger := config.NewLogger(io.Discard)
	repo := &mockRepository{empresas: make(map[string]*domain.Empresa)}
	repo.empresas["e1"] = &domain.Empresa{ID: "e1", CNPJ: "11111111000111", RazaoSocial: "A", TipoInscricao: "1", ClassificacaoTributaria: "01", IndCooperativa: "N", IndConstrutora: "N", IndDesoneracao: "N", Situacao: "ATIVA", DataCadastro: time.Now().UTC(), DataAtualizacao: time.Now().UTC()}
	service := domain.NewService(repo, logger)
	uc := usecase.NewUsecase(service, logger)

	lst, err := uc.ListEmpresas(&domain.ListRequest{Page: 1, Limit: 10})
	if err != nil {
		t.Fatalf("ListEmpresas failed: %v", err)
	}
	if lst.Total != 1 || len(lst.Data) != 1 {
		t.Fatalf("unexpected list output %v", lst)
	}
}
