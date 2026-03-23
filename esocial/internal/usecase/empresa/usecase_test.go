package empresa_test

import (
	"fmt"
	"io"
	"testing"
	"time"

	config "github.com/avelinobego/esocial/configs"
	domain "github.com/avelinobego/esocial/internal/domain/empresa"
	usecase "github.com/avelinobego/esocial/internal/usecase/empresa"
)

type mockRepository struct {
	empresas         map[string]*domain.Empresa
	estabelecimentos map[string]*domain.Estabelecimento
	rubricas         map[string]*domain.Rubrica
	lotacoes         map[string]*domain.Lotacao
	cargos           map[string]*domain.Cargo
}

func newMockRepository() *mockRepository {
	return &mockRepository{
		empresas:         make(map[string]*domain.Empresa),
		estabelecimentos: make(map[string]*domain.Estabelecimento),
		rubricas:         make(map[string]*domain.Rubrica),
		lotacoes:         make(map[string]*domain.Lotacao),
		cargos:           make(map[string]*domain.Cargo),
	}
}

func (m *mockRepository) CreateEmpresa(e *domain.Empresa) error {
	if e.ID == "" {
		e.ID = fmt.Sprintf("empresa-%d", len(m.empresas)+1)
	}
	e.DataCadastro = time.Now().UTC()
	e.DataAtualizacao = e.DataCadastro
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

func (m *mockRepository) CreateEstabelecimento(e *domain.Estabelecimento) error {
	if e.ID == "" {
		e.ID = fmt.Sprintf("est-%d", len(m.estabelecimentos)+1)
	}
	e.DataCadastro = time.Now().UTC()
	m.estabelecimentos[e.ID] = e
	return nil
}

func (m *mockRepository) GetEstabelecimentoByID(id string) (*domain.Estabelecimento, error) {
	if e, ok := m.estabelecimentos[id]; ok {
		return e, nil
	}
	return nil, nil
}

func (m *mockRepository) UpdateEstabelecimento(id string, updates map[string]interface{}) error {
	return nil
}
func (m *mockRepository) DeleteEstabelecimento(id string) error {
	delete(m.estabelecimentos, id)
	return nil
}
func (m *mockRepository) ListEstabelecimentos(filter map[string]interface{}, limit int, offset int) ([]*domain.Estabelecimento, int64, error) {
	list := make([]*domain.Estabelecimento, 0, len(m.estabelecimentos))
	for _, e := range m.estabelecimentos {
		list = append(list, e)
	}
	return list, int64(len(list)), nil
}

func (m *mockRepository) CreateRubrica(r *domain.Rubrica) error {
	if r.ID == "" {
		r.ID = fmt.Sprintf("rub-%d", len(m.rubricas)+1)
	}
	r.DataCadastro = time.Now().UTC()
	m.rubricas[r.ID] = r
	return nil
}

func (m *mockRepository) GetRubricaByID(id string) (*domain.Rubrica, error) {
	if r, ok := m.rubricas[id]; ok {
		return r, nil
	}
	return nil, nil
}

func (m *mockRepository) UpdateRubrica(id string, updates map[string]interface{}) error { return nil }
func (m *mockRepository) DeleteRubrica(id string) error                                 { delete(m.rubricas, id); return nil }
func (m *mockRepository) ListRubricas(filter map[string]interface{}, limit int, offset int) ([]*domain.Rubrica, int64, error) {
	list := make([]*domain.Rubrica, 0, len(m.rubricas))
	for _, r := range m.rubricas {
		list = append(list, r)
	}
	return list, int64(len(list)), nil
}

func (m *mockRepository) CreateLotacao(l *domain.Lotacao) error {
	if l.ID == "" {
		l.ID = fmt.Sprintf("lot-%d", len(m.lotacoes)+1)
	}
	l.DataCadastro = time.Now().UTC()
	m.lotacoes[l.ID] = l
	return nil
}

func (m *mockRepository) GetLotacaoByID(id string) (*domain.Lotacao, error) {
	if l, ok := m.lotacoes[id]; ok {
		return l, nil
	}
	return nil, nil
}

func (m *mockRepository) UpdateLotacao(id string, updates map[string]interface{}) error { return nil }
func (m *mockRepository) DeleteLotacao(id string) error                                 { delete(m.lotacoes, id); return nil }
func (m *mockRepository) ListLotacoes(filter map[string]interface{}, limit int, offset int) ([]*domain.Lotacao, int64, error) {
	list := make([]*domain.Lotacao, 0, len(m.lotacoes))
	for _, l := range m.lotacoes {
		list = append(list, l)
	}
	return list, int64(len(list)), nil
}

func (m *mockRepository) CreateCargo(c *domain.Cargo) error {
	if c.ID == "" {
		c.ID = fmt.Sprintf("cargo-%d", len(m.cargos)+1)
	}
	c.DataCadastro = time.Now().UTC()
	m.cargos[c.ID] = c
	return nil
}

func (m *mockRepository) GetCargoByID(id string) (*domain.Cargo, error) {
	if c, ok := m.cargos[id]; ok {
		return c, nil
	}
	return nil, nil
}

func (m *mockRepository) UpdateCargo(id string, updates map[string]interface{}) error { return nil }
func (m *mockRepository) DeleteCargo(id string) error                                 { delete(m.cargos, id); return nil }
func (m *mockRepository) ListCargos(filter map[string]interface{}, limit int, offset int) ([]*domain.Cargo, int64, error) {
	list := make([]*domain.Cargo, 0, len(m.cargos))
	for _, c := range m.cargos {
		list = append(list, c)
	}
	return list, int64(len(list)), nil
}

func TestUsecaseDelegatesToService(t *testing.T) {
	logger := config.NewLogger(io.Discard)
	repo := newMockRepository()
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
	repo := newMockRepository()
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

func TestUsecaseEstabelecimentoCRUD(t *testing.T) {
	logger := config.NewLogger(io.Discard)
	repo := newMockRepository()
	repo.empresas["emp1"] = &domain.Empresa{
		ID: "emp1", CNPJ: "11111111000111", RazaoSocial: "E",
		TipoInscricao: "1", ClassificacaoTributaria: "01",
		IndCooperativa: "N", IndConstrutora: "N", IndDesoneracao: "N",
		Situacao: "ATIVA", DataCadastro: time.Now().UTC(), DataAtualizacao: time.Now().UTC(),
	}
	service := domain.NewService(repo, logger)
	uc := usecase.NewUsecase(service, logger)

	t.Run("create estabelecimento", func(t *testing.T) {
		req := &domain.CreateEstabelecimentoRequest{
			EmpresaID: "emp1", TipoInscricao: "1",
			NumeroInscricao: "12345", IndObra: "N",
		}
		created, err := uc.CreateEstabelecimento(req)
		if err != nil {
			t.Fatalf("CreateEstabelecimento failed: %v", err)
		}
		if created.ID == "" {
			t.Fatal("expected ID to be assigned")
		}
		if created.EmpresaID != "emp1" {
			t.Fatalf("expected empresa_id emp1, got %s", created.EmpresaID)
		}
	})

	t.Run("get estabelecimento", func(t *testing.T) {
		// First create one
		req := &domain.CreateEstabelecimentoRequest{
			EmpresaID: "emp1", TipoInscricao: "1",
			NumeroInscricao: "12346", IndObra: "N",
		}
		created, err := uc.CreateEstabelecimento(req)
		if err != nil {
			t.Fatalf("CreateEstabelecimento failed: %v", err)
		}

		// Then get it
		fetched, err := uc.GetEstabelecimentoByID(created.ID)
		if err != nil {
			t.Fatalf("GetEstabelecimentoByID failed: %v", err)
		}
		if fetched.ID != created.ID {
			t.Fatalf("expected id %s, got %s", created.ID, fetched.ID)
		}
	})

	t.Run("get missing estabelecimento", func(t *testing.T) {
		_, err := uc.GetEstabelecimentoByID("not-found")
		if err != domain.ErrEstabelecimentoNotFound {
			t.Fatalf("expected ErrEstabelecimentoNotFound, got %v", err)
		}
	})
}

func TestUsecaseRubricaCRUD(t *testing.T) {
	logger := config.NewLogger(io.Discard)
	repo := newMockRepository()
	repo.empresas["emp2"] = &domain.Empresa{
		ID: "emp2", CNPJ: "22222222000122", RazaoSocial: "E2",
		TipoInscricao: "1", ClassificacaoTributaria: "01",
		IndCooperativa: "N", IndConstrutora: "N", IndDesoneracao: "N",
		Situacao: "ATIVA", DataCadastro: time.Now().UTC(), DataAtualizacao: time.Now().UTC(),
	}
	service := domain.NewService(repo, logger)
	uc := usecase.NewUsecase(service, logger)

	t.Run("create rubrica", func(t *testing.T) {
		req := &domain.CreateRubricaRequest{
			EmpresaID: "emp2", Codigo: "R01", Descricao: "Rubrica 1",
			NaturezaRubrica: "0001", TipoRubrica: "D",
			DataInicioValidade: time.Now().UTC(),
		}
		created, err := uc.CreateRubrica(req)
		if err != nil {
			t.Fatalf("CreateRubrica failed: %v", err)
		}
		if created.ID == "" {
			t.Fatal("expected ID to be assigned")
		}
		if created.Codigo != "R01" {
			t.Fatalf("expected codigo R01, got %s", created.Codigo)
		}
	})

	t.Run("get rubrica", func(t *testing.T) {
		req := &domain.CreateRubricaRequest{
			EmpresaID: "emp2", Codigo: "R02", Descricao: "Rubrica 2",
			NaturezaRubrica: "0001", TipoRubrica: "D",
			DataInicioValidade: time.Now().UTC(),
		}
		created, err := uc.CreateRubrica(req)
		if err != nil {
			t.Fatalf("CreateRubrica failed: %v", err)
		}

		fetched, err := uc.GetRubricaByID(created.ID)
		if err != nil {
			t.Fatalf("GetRubricaByID failed: %v", err)
		}
		if fetched.ID != created.ID {
			t.Fatalf("expected id %s, got %s", created.ID, fetched.ID)
		}
	})

	t.Run("get missing rubrica", func(t *testing.T) {
		_, err := uc.GetRubricaByID("not-found")
		if err != domain.ErrRubricaNotFound {
			t.Fatalf("expected ErrRubricaNotFound, got %v", err)
		}
	})
}

func TestUsecaseLotacaoCRUD(t *testing.T) {
	logger := config.NewLogger(io.Discard)
	repo := newMockRepository()
	repo.empresas["emp3"] = &domain.Empresa{
		ID: "emp3", CNPJ: "33333333000133", RazaoSocial: "E3",
		TipoInscricao: "1", ClassificacaoTributaria: "01",
		IndCooperativa: "N", IndConstrutora: "N", IndDesoneracao: "N",
		Situacao: "ATIVA", DataCadastro: time.Now().UTC(), DataAtualizacao: time.Now().UTC(),
	}
	service := domain.NewService(repo, logger)
	uc := usecase.NewUsecase(service, logger)

	t.Run("create lotacao", func(t *testing.T) {
		req := &domain.CreateLotacaoRequest{
			EmpresaID: "emp3", Codigo: "L01", Descricao: "Lotacao 1",
			TipoLotacao: "01", DataInicioValidade: time.Now().UTC(),
		}
		created, err := uc.CreateLotacao(req)
		if err != nil {
			t.Fatalf("CreateLotacao failed: %v", err)
		}
		if created.ID == "" {
			t.Fatal("expected ID to be assigned")
		}
		if created.Codigo != "L01" {
			t.Fatalf("expected codigo L01, got %s", created.Codigo)
		}
	})

	t.Run("get lotacao", func(t *testing.T) {
		req := &domain.CreateLotacaoRequest{
			EmpresaID: "emp3", Codigo: "L02", Descricao: "Lotacao 2",
			TipoLotacao: "01", DataInicioValidade: time.Now().UTC(),
		}
		created, err := uc.CreateLotacao(req)
		if err != nil {
			t.Fatalf("CreateLotacao failed: %v", err)
		}

		fetched, err := uc.GetLotacaoByID(created.ID)
		if err != nil {
			t.Fatalf("GetLotacaoByID failed: %v", err)
		}
		if fetched.ID != created.ID {
			t.Fatalf("expected id %s, got %s", created.ID, fetched.ID)
		}
	})

	t.Run("get missing lotacao", func(t *testing.T) {
		_, err := uc.GetLotacaoByID("not-found")
		if err != domain.ErrLotacaoNotFound {
			t.Fatalf("expected ErrLotacaoNotFound, got %v", err)
		}
	})
}

func TestUsecaseCargoCRUD(t *testing.T) {
	logger := config.NewLogger(io.Discard)
	repo := newMockRepository()
	repo.empresas["emp4"] = &domain.Empresa{
		ID: "emp4", CNPJ: "44444444000144", RazaoSocial: "E4",
		TipoInscricao: "1", ClassificacaoTributaria: "01",
		IndCooperativa: "N", IndConstrutora: "N", IndDesoneracao: "N",
		Situacao: "ATIVA", DataCadastro: time.Now().UTC(), DataAtualizacao: time.Now().UTC(),
	}
	service := domain.NewService(repo, logger)
	uc := usecase.NewUsecase(service, logger)

	t.Run("create cargo", func(t *testing.T) {
		req := &domain.CreateCargoRequest{
			EmpresaID: "emp4", Codigo: "C01", Descricao: "Cargo 1",
			CBO: "123456", DataInicioValidade: time.Now().UTC(),
		}
		created, err := uc.CreateCargo(req)
		if err != nil {
			t.Fatalf("CreateCargo failed: %v", err)
		}
		if created.ID == "" {
			t.Fatal("expected ID to be assigned")
		}
		if created.Codigo != "C01" {
			t.Fatalf("expected codigo C01, got %s", created.Codigo)
		}
	})

	t.Run("get cargo", func(t *testing.T) {
		req := &domain.CreateCargoRequest{
			EmpresaID: "emp4", Codigo: "C02", Descricao: "Cargo 2",
			CBO: "654321", DataInicioValidade: time.Now().UTC(),
		}
		created, err := uc.CreateCargo(req)
		if err != nil {
			t.Fatalf("CreateCargo failed: %v", err)
		}

		fetched, err := uc.GetCargoByID(created.ID)
		if err != nil {
			t.Fatalf("GetCargoByID failed: %v", err)
		}
		if fetched.ID != created.ID {
			t.Fatalf("expected id %s, got %s", created.ID, fetched.ID)
		}
	})

	t.Run("get missing cargo", func(t *testing.T) {
		_, err := uc.GetCargoByID("not-found")
		if err != domain.ErrCargoNotFound {
			t.Fatalf("expected ErrCargoNotFound, got %v", err)
		}
	})
}
