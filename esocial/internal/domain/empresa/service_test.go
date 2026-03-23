package empresa

import (
	"errors"
	"fmt"
	"io"
	"testing"
	"time"

	"github.com/avelinobego/esocial/configs"
)

type mockRepository struct {
	empresas         map[string]*Empresa
	estabelecimentos map[string]*Estabelecimento
	rubricas         map[string]*Rubrica
	lotacoes         map[string]*Lotacao
	cargos           map[string]*Cargo
	errors           map[string]error
}

func newMockRepository() *mockRepository {
	return &mockRepository{
		empresas:         make(map[string]*Empresa),
		estabelecimentos: make(map[string]*Estabelecimento),
		rubricas:         make(map[string]*Rubrica),
		lotacoes:         make(map[string]*Lotacao),
		cargos:           make(map[string]*Cargo),
		errors:           make(map[string]error),
	}
}

func (m *mockRepository) errFor(method string) error {
	if m.errors == nil {
		return nil
	}
	if err, ok := m.errors[method]; ok {
		return err
	}
	return nil
}

func (m *mockRepository) CreateEmpresa(e *Empresa) error {
	if err := m.errFor("CreateEmpresa"); err != nil {
		return err
	}

	if e.ID == "" {
		e.ID = fmt.Sprintf("empresa-%d", len(m.empresas)+1)
	}
	e.DataCadastro = time.Now().UTC()
	e.DataAtualizacao = e.DataCadastro
	m.empresas[e.ID] = e
	return nil
}

func (m *mockRepository) GetEmpresaByID(id string) (*Empresa, error) {
	if err := m.errFor("GetEmpresaByID"); err != nil {
		return nil, err
	}

	if e, ok := m.empresas[id]; ok {
		return e, nil
	}
	return nil, nil
}

func (m *mockRepository) GetEmpresaByCNPJ(cnpj string) (*Empresa, error) {
	if err := m.errFor("GetEmpresaByCNPJ"); err != nil {
		return nil, err
	}

	for _, e := range m.empresas {
		if e.CNPJ == cnpj {
			return e, nil
		}
	}
	return nil, nil
}

func (m *mockRepository) UpdateEmpresa(id string, updates map[string]interface{}) error {
	if err := m.errFor("UpdateEmpresa"); err != nil {
		return err
	}

	if e, ok := m.empresas[id]; ok {
		if v, ok := updates["razao_social"]; ok {
			switch v2 := v.(type) {
			case string:
				if v2 != "" {
					e.RazaoSocial = v2
				}
			case *string:
				if v2 != nil {
					e.RazaoSocial = *v2
				}
			}
		}
		if v, ok := updates["situacao"]; ok {
			switch v2 := v.(type) {
			case string:
				if v2 != "" {
					e.Situacao = v2
				}
			case *string:
				if v2 != nil {
					e.Situacao = *v2
				}
			}
		}
		e.DataAtualizacao = time.Now().UTC()
		m.empresas[id] = e
	}
	return nil
}

func (m *mockRepository) DeleteEmpresa(id string) error {
	if err := m.errFor("DeleteEmpresa"); err != nil {
		return err
	}

	delete(m.empresas, id)
	return nil
}

func (m *mockRepository) ListEmpresas(filter map[string]interface{}, limit int, offset int) ([]*Empresa, int64, error) {
	if err := m.errFor("ListEmpresas"); err != nil {
		return nil, 0, err
	}

	list := make([]*Empresa, 0, len(m.empresas))
	for _, e := range m.empresas {
		list = append(list, e)
	}
	return list, int64(len(list)), nil
}

func (m *mockRepository) CreateEstabelecimento(e *Estabelecimento) error {
	if err := m.errFor("CreateEstabelecimento"); err != nil {
		return err
	}

	if e.ID == "" {
		e.ID = fmt.Sprintf("est-%d", len(m.estabelecimentos)+1)
	}
	e.DataCadastro = time.Now().UTC()
	m.estabelecimentos[e.ID] = e
	return nil
}

func (m *mockRepository) GetEstabelecimentoByID(id string) (*Estabelecimento, error) {
	if err := m.errFor("GetEstabelecimentoByID"); err != nil {
		return nil, err
	}

	if e, ok := m.estabelecimentos[id]; ok {
		return e, nil
	}
	return nil, nil
}

func (m *mockRepository) UpdateEstabelecimento(id string, updates map[string]interface{}) error {
	if err := m.errFor("UpdateEstabelecimento"); err != nil {
		return err
	}

	if e, ok := m.estabelecimentos[id]; ok {
		if v, ok := updates["situacao"]; ok {
			switch v2 := v.(type) {
			case string:
				if v2 != "" {
					e.Situacao = v2
				}
			case *string:
				if v2 != nil {
					e.Situacao = *v2
				}
			}
		}
		m.estabelecimentos[id] = e
	}
	return nil
}

func (m *mockRepository) DeleteEstabelecimento(id string) error {
	if err := m.errFor("DeleteEstabelecimento"); err != nil {
		return err
	}

	delete(m.estabelecimentos, id)
	return nil
}

func (m *mockRepository) ListEstabelecimentos(filter map[string]interface{}, limit int, offset int) ([]*Estabelecimento, int64, error) {
	if err := m.errFor("ListEstabelecimentos"); err != nil {
		return nil, 0, err
	}

	list := make([]*Estabelecimento, 0, len(m.estabelecimentos))
	for _, e := range m.estabelecimentos {
		list = append(list, e)
	}
	return list, int64(len(list)), nil
}

func (m *mockRepository) CreateRubrica(r *Rubrica) error {
	if err := m.errFor("CreateRubrica"); err != nil {
		return err
	}

	if r.ID == "" {
		r.ID = fmt.Sprintf("rub-%d", len(m.rubricas)+1)
	}
	r.DataCadastro = time.Now().UTC()
	m.rubricas[r.ID] = r
	return nil
}

func (m *mockRepository) GetRubricaByID(id string) (*Rubrica, error) {
	if err := m.errFor("GetRubricaByID"); err != nil {
		return nil, err
	}

	if r, ok := m.rubricas[id]; ok {
		return r, nil
	}
	return nil, nil
}

func (m *mockRepository) UpdateRubrica(id string, updates map[string]interface{}) error {
	if err := m.errFor("UpdateRubrica"); err != nil {
		return err
	}

	if r, ok := m.rubricas[id]; ok {
		if v, ok := updates["descricao"]; ok {
			switch v2 := v.(type) {
			case string:
				if v2 != "" {
					r.Descricao = v2
				}
			case *string:
				if v2 != nil {
					r.Descricao = *v2
				}
			}
		}
		m.rubricas[id] = r
	}
	return nil
}

func (m *mockRepository) DeleteRubrica(id string) error {
	if err := m.errFor("DeleteRubrica"); err != nil {
		return err
	}

	delete(m.rubricas, id)
	return nil
}

func (m *mockRepository) ListRubricas(filter map[string]interface{}, limit int, offset int) ([]*Rubrica, int64, error) {
	if err := m.errFor("ListRubricas"); err != nil {
		return nil, 0, err
	}

	list := make([]*Rubrica, 0, len(m.rubricas))
	for _, r := range m.rubricas {
		list = append(list, r)
	}
	return list, int64(len(list)), nil
}

func (m *mockRepository) CreateLotacao(l *Lotacao) error {
	if err := m.errFor("CreateLotacao"); err != nil {
		return err
	}

	if l.ID == "" {
		l.ID = fmt.Sprintf("lot-%d", len(m.lotacoes)+1)
	}
	l.DataCadastro = time.Now().UTC()
	m.lotacoes[l.ID] = l
	return nil
}

func (m *mockRepository) GetLotacaoByID(id string) (*Lotacao, error) {
	if err := m.errFor("GetLotacaoByID"); err != nil {
		return nil, err
	}

	if l, ok := m.lotacoes[id]; ok {
		return l, nil
	}
	return nil, nil
}

func (m *mockRepository) UpdateLotacao(id string, updates map[string]interface{}) error {
	if err := m.errFor("UpdateLotacao"); err != nil {
		return err
	}

	if l, ok := m.lotacoes[id]; ok {
		if v, ok := updates["descricao"]; ok {
			switch vv := v.(type) {
			case string:
				if vv != "" {
					l.Descricao = vv
				}
			case *string:
				if vv != nil {
					l.Descricao = *vv
				}
			}
		}
		m.lotacoes[id] = l
	}
	return nil
}

func (m *mockRepository) DeleteLotacao(id string) error {
	if err := m.errFor("DeleteLotacao"); err != nil {
		return err
	}

	delete(m.lotacoes, id)
	return nil
}

func (m *mockRepository) ListLotacoes(filter map[string]interface{}, limit int, offset int) ([]*Lotacao, int64, error) {
	if err := m.errFor("ListLotacoes"); err != nil {
		return nil, 0, err
	}

	list := make([]*Lotacao, 0, len(m.lotacoes))
	for _, l := range m.lotacoes {
		list = append(list, l)
	}
	return list, int64(len(list)), nil
}

func (m *mockRepository) CreateCargo(c *Cargo) error {
	if err := m.errFor("CreateCargo"); err != nil {
		return err
	}

	if c.ID == "" {
		c.ID = fmt.Sprintf("cargo-%d", len(m.cargos)+1)
	}
	c.DataCadastro = time.Now().UTC()
	m.cargos[c.ID] = c
	return nil
}

func (m *mockRepository) GetCargoByID(id string) (*Cargo, error) {
	if err := m.errFor("GetCargoByID"); err != nil {
		return nil, err
	}

	if c, ok := m.cargos[id]; ok {
		return c, nil
	}
	return nil, nil
}

func (m *mockRepository) UpdateCargo(id string, updates map[string]interface{}) error {
	if err := m.errFor("UpdateCargo"); err != nil {
		return err
	}

	if c, ok := m.cargos[id]; ok {
		if v, ok := updates["descricao"]; ok {
			switch vv := v.(type) {
			case string:
				if vv != "" {
					c.Descricao = vv
				}
			case *string:
				if vv != nil {
					c.Descricao = *vv
				}
			}
		}
		m.cargos[id] = c
	}
	return nil
}

func (m *mockRepository) DeleteCargo(id string) error {
	if err := m.errFor("DeleteCargo"); err != nil {
		return err
	}

	delete(m.cargos, id)
	return nil
}

func (m *mockRepository) ListCargos(filter map[string]interface{}, limit int, offset int) ([]*Cargo, int64, error) {
	if err := m.errFor("ListCargos"); err != nil {
		return nil, 0, err
	}

	list := make([]*Cargo, 0, len(m.cargos))
	for _, c := range m.cargos {
		list = append(list, c)
	}
	return list, int64(len(list)), nil
}

func setupService() (*Service, *mockRepository) {
	logger := configs.NewLogger(io.Discard)
	repo := newMockRepository()
	return NewService(repo, logger), repo
}

func TestServiceEmpresaCRUD(t *testing.T) {
	s, repo := setupService()

	t.Run("invalid cnpj", func(t *testing.T) {
		_, err := s.CreateEmpresa(&CreateEmpresaRequest{CNPJ: "123"})
		if err != ErrInvalidCNPJ {
			t.Fatalf("expected ErrInvalidCNPJ, got %v", err)
		}
	})

	t.Run("create and retrieve empresa", func(t *testing.T) {
		req := &CreateEmpresaRequest{
			CNPJ:                    "12345678000195",
			RazaoSocial:             "Test Company",
			TipoInscricao:           "1",
			ClassificacaoTributaria: "01",
			IndCooperativa:          "N",
			IndConstrutora:          "N",
			IndDesoneracao:          "N",
		}

		resp, err := s.CreateEmpresa(req)
		if err != nil {
			t.Fatalf("failed to create empresa: %v", err)
		}
		if resp.CNPJ != req.CNPJ {
			t.Fatalf("expected cnpj %s, got %s", req.CNPJ, resp.CNPJ)
		}

		_, err = s.CreateEmpresa(req)
		if err != ErrEmpresaAlreadyExists {
			t.Fatalf("expected ErrEmpresaAlreadyExists, got %v", err)
		}

		found, err := s.GetEmpresaByID(resp.ID)
		if err != nil {
			t.Fatalf("GetEmpresaByID returned error: %v", err)
		}
		if found == nil || found.ID != resp.ID {
			t.Fatalf("expected empresa id %s, got %v", resp.ID, found)
		}

		updatedData := &UpdateEmpresaRequest{RazaoSocial: ptrString("Updated")}
		updated, err := s.UpdateEmpresa(resp.ID, updatedData)
		if err != nil {
			t.Fatalf("UpdateEmpresa returned error: %v", err)
		}
		if updated.RazaoSocial != "Updated" {
			t.Fatalf("expected updated RazaoSocial, got %s", updated.RazaoSocial)
		}

		err = s.DeleteEmpresa(resp.ID)
		if err != nil {
			t.Fatalf("DeleteEmpresa returned error: %v", err)
		}

		_, err = s.GetEmpresaByID(resp.ID)
		if err != ErrEmpresaNotFound {
			t.Fatalf("expected ErrEmpresaNotFound after delete, got %v", err)
		}

		if len(repo.empresas) != 0 {
			t.Fatalf("expected empty empresa map after delete, got %d", len(repo.empresas))
		}
	})

	t.Run("list empresas", func(t *testing.T) {
		repo.empresas = make(map[string]*Empresa)
		repo.CreateEmpresa(&Empresa{ID: "1", CNPJ: "11111111111111", RazaoSocial: "A", TipoInscricao: "1", ClassificacaoTributaria: "01", IndCooperativa: "N", IndConstrutora: "N", IndDesoneracao: "N", Situacao: "ATIVA", DataCadastro: time.Now().UTC(), DataAtualizacao: time.Now().UTC()})
		repo.CreateEmpresa(&Empresa{ID: "2", CNPJ: "22222222000122", RazaoSocial: "B", TipoInscricao: "1", ClassificacaoTributaria: "01", IndCooperativa: "N", IndConstrutora: "N", IndDesoneracao: "N", Situacao: "ATIVA", DataCadastro: time.Now().UTC(), DataAtualizacao: time.Now().UTC()})

		listResp, err := s.ListEmpresas(&ListRequest{Page: 1, Limit: 10})
		if err != nil {
			t.Fatalf("ListEmpresas returned error: %v", err)
		}
		if listResp.Total != 2 {
			t.Fatalf("expected total 2, got %d", listResp.Total)
		}
	})
}

func TestServiceEstabelecimentoAndRubrica(t *testing.T) {
	s, repo := setupService()

	// Set up empresa required for nested entities
	repo.CreateEmpresa(&Empresa{ID: "e1", CNPJ: "33333333000133", RazaoSocial: "E", TipoInscricao: "1", ClassificacaoTributaria: "01", IndCooperativa: "N", IndConstrutora: "N", IndDesoneracao: "N", Situacao: "ATIVA", DataCadastro: time.Now().UTC(), DataAtualizacao: time.Now().UTC()})

	t.Run("create and get estabelecimento", func(t *testing.T) {
		estReq := &CreateEstabelecimentoRequest{EmpresaID: "e1", TipoInscricao: "1", NumeroInscricao: "12345", IndObra: "N"}
		created, err := s.CreateEstabelecimento(estReq)
		if err != nil {
			t.Fatalf("CreateEstabelecimento failed: %v", err)
		}
		if created.EmpresaID != "e1" {
			t.Fatalf("expected empresa_id e1, got %s", created.EmpresaID)
		}

		fetched, err := s.GetEstabelecimentoByID(created.ID)
		if err != nil || fetched == nil {
			t.Fatalf("GetEstabelecimentoByID failed: %v", err)
		}
	})

	t.Run("update estabelecimento", func(t *testing.T) {
		all := []string{}
		for id := range repo.estabelecimentos {
			all = append(all, id)
		}
		if len(all) == 0 {
			t.Fatal("expected one estabelecimento")
		}
		id := all[0]
		_, err := s.UpdateEstabelecimento(id, &UpdateEstabelecimentoRequest{Situacao: ptrString("INATIVO")})
		if err != nil {
			t.Fatalf("UpdateEstabelecimento failed: %v", err)
		}
		updated, _ := s.GetEstabelecimentoByID(id)
		if updated.Situacao != "INATIVO" {
			t.Fatalf("expected INATIVO situacao, got %s", updated.Situacao)
		}
	})

	t.Run("delete estabelecimento", func(t *testing.T) {
		var id string
		for k := range repo.estabelecimentos {
			id = k
			break
		}
		if id == "" {
			t.Fatal("no estabelecimento to delete")
		}
		err := s.DeleteEstabelecimento(id)
		if err != nil {
			t.Fatalf("DeleteEstabelecimento failed: %v", err)
		}
		_, err = s.GetEstabelecimentoByID(id)
		if err != ErrEstabelecimentoNotFound {
			t.Fatalf("expected ErrEstabelecimentoNotFound after delete, got %v", err)
		}
	})

	t.Run("create and get rubrica", func(t *testing.T) {
		rubReq := &CreateRubricaRequest{EmpresaID: "e1", Codigo: "R1", Descricao: "R1", NaturezaRubrica: "0001", TipoRubrica: "D", DataInicioValidade: time.Now().UTC()}
		created, err := s.CreateRubrica(rubReq)
		if err != nil {
			t.Fatalf("CreateRubrica failed: %v", err)
		}
		if created.EmpresaID != "e1" {
			t.Fatalf("expected empresa_id e1, got %s", created.EmpresaID)
		}

		fetched, err := s.GetRubricaByID(created.ID)
		if err != nil || fetched == nil {
			t.Fatalf("GetRubricaByID failed: %v", err)
		}

		_, err = s.UpdateRubrica(created.ID, &UpdateRubricaRequest{Descricao: ptrString("R1-up")})
		if err != nil {
			t.Fatalf("UpdateRubrica failed: %v", err)
		}
		updated, err := s.GetRubricaByID(created.ID)
		if err != nil || updated == nil {
			t.Fatalf("GetRubricaByID after update failed: %v", err)
		}
		if updated.Descricao != "R1-up" {
			t.Fatalf("expected updated rubrica descricao, got %s", updated.Descricao)
		}
	})

	t.Run("delete rubrica", func(t *testing.T) {
		var id string
		for k := range repo.rubricas {
			id = k
			break
		}
		err := s.DeleteRubrica(id)
		if err != nil {
			t.Fatalf("DeleteRubrica failed: %v", err)
		}
		_, err = s.GetRubricaByID(id)
		if err != ErrRubricaNotFound {
			t.Fatalf("expected ErrRubricaNotFound after delete, got %v", err)
		}
	})
}

func TestServiceMapAndListing(t *testing.T) {
	s, repo := setupService()

	// Setup base empresa
	repo.CreateEmpresa(&Empresa{ID: "e1", CNPJ: "99999999000199", RazaoSocial: "X", TipoInscricao: "1", ClassificacaoTributaria: "01", IndCooperativa: "N", IndConstrutora: "N", IndDesoneracao: "N", Situacao: "ATIVA", DataCadastro: time.Now().UTC(), DataAtualizacao: time.Now().UTC()})
	// Add other entities
	repo.CreateEstabelecimento(&Estabelecimento{ID: "est1", EmpresaID: "e1", TipoInscricao: "1", NumeroInscricao: "22222", IndObra: "N", Situacao: "ATIVO", DataCadastro: time.Now().UTC()})
	repo.CreateRubrica(&Rubrica{ID: "rub1", EmpresaID: "e1", Codigo: "R01", Descricao: "desc", NaturezaRubrica: "0001", TipoRubrica: "D", Ativa: true, DataCadastro: time.Now().UTC(), DataInicioValidade: time.Now().UTC()})
	repo.CreateLotacao(&Lotacao{ID: "lot1", EmpresaID: "e1", Codigo: "L01", Descricao: "Ldesc", TipoLotacao: "01", Ativa: true, DataCadastro: time.Now().UTC(), DataInicioValidade: time.Now().UTC()})
	repo.CreateCargo(&Cargo{ID: "car1", EmpresaID: "e1", Codigo: "C01", Descricao: "Cdesc", CBO: "123456", Ativo: true, DataCadastro: time.Now().UTC(), DataInicioValidade: time.Now().UTC()})

	if _, err := s.ListEmpresas(&ListRequest{Page: 1, Limit: 10}); err != nil {
		t.Fatalf("ListEmpresas failed: %v", err)
	}
	if _, err := s.ListEstabelecimentos(&ListRequest{Page: 1, Limit: 10}); err != nil {
		t.Fatalf("ListEstabelecimentos failed: %v", err)
	}
	if _, err := s.ListRubricas(&ListRequest{Page: 1, Limit: 10}); err != nil {
		t.Fatalf("ListRubricas failed: %v", err)
	}
	if _, err := s.ListLotacoes(&ListRequest{Page: 1, Limit: 10}); err != nil {
		t.Fatalf("ListLotacoes failed: %v", err)
	}
	if _, err := s.ListCargos(&ListRequest{Page: 1, Limit: 10}); err != nil {
		t.Fatalf("ListCargos failed: %v", err)
	}

	if _, err := s.GetEmpresaByID("notfound"); err != ErrEmpresaNotFound {
		t.Fatalf("expected ErrEmpresaNotFound, got %v", err)
	}
	if _, err := s.UpdateEmpresa("notfound", &UpdateEmpresaRequest{RazaoSocial: ptrString("X")}); err != ErrEmpresaNotFound {
		t.Fatalf("expected ErrEmpresaNotFound on update, got %v", err)
	}
	if err := s.DeleteEmpresa("notfound"); err != ErrEmpresaNotFound {
		t.Fatalf("expected ErrEmpresaNotFound on delete, got %v", err)
	}

	if _, err := s.GetEstabelecimentoByID("notfound"); err != ErrEstabelecimentoNotFound {
		t.Fatalf("expected ErrEstabelecimentoNotFound, got %v", err)
	}
	if _, err := s.UpdateEstabelecimento("notfound", &UpdateEstabelecimentoRequest{Situacao: ptrString("X")}); err != ErrEstabelecimentoNotFound {
		t.Fatalf("expected ErrEstabelecimentoNotFound on update, got %v", err)
	}
	if err := s.DeleteEstabelecimento("notfound"); err != ErrEstabelecimentoNotFound {
		t.Fatalf("expected ErrEstabelecimentoNotFound on delete, got %v", err)
	}

	if _, err := s.GetRubricaByID("notfound"); err != ErrRubricaNotFound {
		t.Fatalf("expected ErrRubricaNotFound, got %v", err)
	}
	if _, err := s.UpdateRubrica("notfound", &UpdateRubricaRequest{Descricao: ptrString("X")}); err != ErrRubricaNotFound {
		t.Fatalf("expected ErrRubricaNotFound on update, got %v", err)
	}
	if err := s.DeleteRubrica("notfound"); err != ErrRubricaNotFound {
		t.Fatalf("expected ErrRubricaNotFound on delete, got %v", err)
	}

	if _, err := s.GetLotacaoByID("notfound"); err != ErrLotacaoNotFound {
		t.Fatalf("expected ErrLotacaoNotFound, got %v", err)
	}
	if _, err := s.UpdateLotacao("notfound", &UpdateLotacaoRequest{Descricao: ptrString("X")}); err != ErrLotacaoNotFound {
		t.Fatalf("expected ErrLotacaoNotFound on update, got %v", err)
	}
	if err := s.DeleteLotacao("notfound"); err != ErrLotacaoNotFound {
		t.Fatalf("expected ErrLotacaoNotFound on delete, got %v", err)
	}

	if _, err := s.GetCargoByID("notfound"); err != ErrCargoNotFound {
		t.Fatalf("expected ErrCargoNotFound, got %v", err)
	}
	if _, err := s.UpdateCargo("notfound", &UpdateCargoRequest{Descricao: ptrString("X")}); err != ErrCargoNotFound {
		t.Fatalf("expected ErrCargoNotFound on update, got %v", err)
	}
	if err := s.DeleteCargo("notfound"); err != ErrCargoNotFound {
		t.Fatalf("expected ErrCargoNotFound on delete, got %v", err)
	}

	x := s.mapEmpresaToResponse(&Empresa{ID: "xx", CNPJ: "11111111000111", RazaoSocial: "Z"})
	if x.ID != "xx" {
		t.Fatalf("unexpected maps value")
	}
	y := s.mapEstabelecimentoToResponse(&Estabelecimento{ID: "e1", EmpresaID: "e1", TipoInscricao: "1", NumeroInscricao: "22222", Situacao: "ATIVO", DataCadastro: time.Now().UTC()})
	if y.ID != "e1" {
		t.Fatalf("unexpected maps value")
	}
	z := s.mapRubricaToResponse(&Rubrica{ID: "r1", EmpresaID: "e1", Codigo: "R1", Descricao: "desc", NaturezaRubrica: "0001", TipoRubrica: "D", Ativa: true, DataCadastro: time.Now().UTC(), DataInicioValidade: time.Now().UTC()})
	if z.ID != "r1" {
		t.Fatalf("unexpected maps value")
	}
	l := s.mapLotacaoToResponse(&Lotacao{ID: "l1", EmpresaID: "e1", Codigo: "L1", Descricao: "l", TipoLotacao: "01", Ativa: true, DataCadastro: time.Now().UTC(), DataInicioValidade: time.Now().UTC()})
	if l.ID != "l1" {
		t.Fatalf("unexpected maps value")
	}
	c := s.mapCargoToResponse(&Cargo{ID: "c1", EmpresaID: "e1", Codigo: "C1", Descricao: "c", CBO: "123456", Ativo: true, DataCadastro: time.Now().UTC(), DataInicioValidade: time.Now().UTC()})
	if c.ID != "c1" {
		t.Fatalf("unexpected maps value")
	}
}

func ptrString(value string) *string {
	return &value
}

func TestServiceLotacaoAndCargo(t *testing.T) {
	s, repo := setupService()

	// create company for foreign key
	repo.CreateEmpresa(&Empresa{ID: "e2", CNPJ: "44444444000144", RazaoSocial: "Z", TipoInscricao: "1", ClassificacaoTributaria: "01", IndCooperativa: "N", IndConstrutora: "N", IndDesoneracao: "N", Situacao: "ATIVA", DataCadastro: time.Now().UTC(), DataAtualizacao: time.Now().UTC()})

	t.Run("create and get lotacao", func(t *testing.T) {
		lotReq := &CreateLotacaoRequest{EmpresaID: "e2", Codigo: "L1", Descricao: "L1", TipoLotacao: "01", DataInicioValidade: time.Now().UTC()}
		created, err := s.CreateLotacao(lotReq)
		if err != nil {
			t.Fatalf("CreateLotacao failed: %v", err)
		}
		if created.EmpresaID != "e2" {
			t.Fatalf("expected empresa_id e2, got %s", created.EmpresaID)
		}

		fetched, err := s.GetLotacaoByID(created.ID)
		if err != nil || fetched == nil {
			t.Fatalf("GetLotacaoByID failed: %v", err)
		}

		_, err = s.UpdateLotacao(created.ID, &UpdateLotacaoRequest{Descricao: ptrString("L1-up")})
		if err != nil {
			t.Fatalf("UpdateLotacao failed: %v", err)
		}
		updated, err := s.GetLotacaoByID(created.ID)
		if err != nil || updated == nil {
			t.Fatalf("GetLotacaoByID after update failed: %v", err)
		}
		if updated.Descricao != "L1-up" {
			t.Fatalf("expected updated lotacao descricao, got %s", updated.Descricao)
		}
	})

	t.Run("delete lotacao", func(t *testing.T) {
		var id string
		for k := range repo.lotacoes {
			id = k
			break
		}
		err := s.DeleteLotacao(id)
		if err != nil {
			t.Fatalf("DeleteLotacao failed: %v", err)
		}
		_, err = s.GetLotacaoByID(id)
		if err != ErrLotacaoNotFound {
			t.Fatalf("expected ErrLotacaoNotFound after delete, got %v", err)
		}
	})

	t.Run("create and get cargo", func(t *testing.T) {
		carReq := &CreateCargoRequest{EmpresaID: "e2", Codigo: "C1", Descricao: "C1", CBO: "123456", DataInicioValidade: time.Now().UTC()}
		created, err := s.CreateCargo(carReq)
		if err != nil {
			t.Fatalf("CreateCargo failed: %v", err)
		}
		if created.EmpresaID != "e2" {
			t.Fatalf("expected empresa_id e2, got %s", created.EmpresaID)
		}

		fetched, err := s.GetCargoByID(created.ID)
		if err != nil || fetched == nil {
			t.Fatalf("GetCargoByID failed: %v", err)
		}

		_, err = s.UpdateCargo(created.ID, &UpdateCargoRequest{Descricao: ptrString("C1-up")})
		if err != nil {
			t.Fatalf("UpdateCargo failed: %v", err)
		}
		updated, err := s.GetCargoByID(created.ID)
		if err != nil || updated == nil {
			t.Fatalf("GetCargoByID after update failed: %v", err)
		}
		if updated.Descricao != "C1-up" {
			t.Fatalf("expected updated cargo descricao, got %s", updated.Descricao)
		}
	})

	t.Run("delete cargo", func(t *testing.T) {
		var id string
		for k := range repo.cargos {
			id = k
			break
		}
		err := s.DeleteCargo(id)
		if err != nil {
			t.Fatalf("DeleteCargo failed: %v", err)
		}
		_, err = s.GetCargoByID(id)
		if err != ErrCargoNotFound {
			t.Fatalf("expected ErrCargoNotFound after delete, got %v", err)
		}
	})
}

func TestServiceRepoErrorBranches(t *testing.T) {
	s, repo := setupService()

	repo.errors["GetEmpresaByID"] = errors.New("repo failure")
	if _, err := s.CreateEstabelecimento(&CreateEstabelecimentoRequest{EmpresaID: "e1"}); err == nil {
		t.Fatal("expected error from CreateEstabelecimento when GetEmpresaByID fails")
	}
	if _, err := s.CreateRubrica(&CreateRubricaRequest{EmpresaID: "e1"}); err == nil {
		t.Fatal("expected error from CreateRubrica when GetEmpresaByID fails")
	}
	if _, err := s.CreateLotacao(&CreateLotacaoRequest{EmpresaID: "e1"}); err == nil {
		t.Fatal("expected error from CreateLotacao when GetEmpresaByID fails")
	}
	if _, err := s.CreateCargo(&CreateCargoRequest{EmpresaID: "e1"}); err == nil {
		t.Fatal("expected error from CreateCargo when GetEmpresaByID fails")
	}

	repo.errors = map[string]error{"GetEstabelecimentoByID": errors.New("repo failure")}
	if _, err := s.GetEstabelecimentoByID("x"); err == nil {
		t.Fatal("expected error from GetEstabelecimentoByID when repo fails")
	}

	repo.errors = map[string]error{"GetRubricaByID": errors.New("repo failure")}
	if _, err := s.GetRubricaByID("x"); err == nil {
		t.Fatal("expected error from GetRubricaByID when repo fails")
	}

	repo.errors = map[string]error{"GetLotacaoByID": errors.New("repo failure")}
	if _, err := s.GetLotacaoByID("x"); err == nil {
		t.Fatal("expected error from GetLotacaoByID when repo fails")
	}

	repo.errors = map[string]error{"GetCargoByID": errors.New("repo failure")}
	if _, err := s.GetCargoByID("x"); err == nil {
		t.Fatal("expected error from GetCargoByID when repo fails")
	}

	repo.errors = map[string]error{"ListEmpresas": errors.New("repo failure")}
	if _, err := s.ListEmpresas(&ListRequest{Page: 1, Limit: 10}); err == nil {
		t.Fatal("expected error from ListEmpresas when repo fails")
	}

	repo.errors = map[string]error{"ListEstabelecimentos": errors.New("repo failure")}
	if _, err := s.ListEstabelecimentos(&ListRequest{Page: 1, Limit: 10}); err == nil {
		t.Fatal("expected error from ListEstabelecimentos when repo fails")
	}

	repo.errors = map[string]error{"ListRubricas": errors.New("repo failure")}
	if _, err := s.ListRubricas(&ListRequest{Page: 1, Limit: 10}); err == nil {
		t.Fatal("expected error from ListRubricas when repo fails")
	}

	repo.errors = map[string]error{"ListLotacoes": errors.New("repo failure")}
	if _, err := s.ListLotacoes(&ListRequest{Page: 1, Limit: 10}); err == nil {
		t.Fatal("expected error from ListLotacoes when repo fails")
	}

	repo.errors = map[string]error{"ListCargos": errors.New("repo failure")}
	if _, err := s.ListCargos(&ListRequest{Page: 1, Limit: 10}); err == nil {
		t.Fatal("expected error from ListCargos when repo fails")
	}
}

func TestServiceFetchErrorPaths(t *testing.T) {
	s, repo := setupService()

	// Set up a empresa for testing
	repo.CreateEmpresa(&Empresa{
		ID: "test-emp", CNPJ: "55555555000155", RazaoSocial: "Test",
		TipoInscricao: "1", ClassificacaoTributaria: "01",
		IndCooperativa: "N", IndConstrutora: "N", IndDesoneracao: "N",
		Situacao: "ATIVA", DataCadastro: time.Now().UTC(), DataAtualizacao: time.Now().UTC(),
	})

	// Test GetEmpresaByID error path (repo error)
	repo.errors = map[string]error{"GetEmpresaByID": errors.New("db connection error")}
	if _, err := s.GetEmpresaByID("test-emp"); err == nil {
		t.Fatal("expected error from GetEmpresaByID when repo fails")
	}

	// Test UpdateEmpresa error path: GetEmpresaByID fails on initial check
	repo.errors = map[string]error{"GetEmpresaByID": errors.New("db connection error")}
	if _, err := s.UpdateEmpresa("test-emp", &UpdateEmpresaRequest{RazaoSocial: ptrString("Updated")}); err == nil {
		t.Fatal("expected error from UpdateEmpresa when initial GetEmpresaByID fails")
	}

	// Test DeleteEmpresa error path: GetEmpresaByID fails on initial check
	repo.errors = map[string]error{"GetEmpresaByID": errors.New("db connection error")}
	if err := s.DeleteEmpresa("test-emp"); err == nil {
		t.Fatal("expected error from DeleteEmpresa when GetEmpresaByID fails")
	}

	// Test UpdateEmpresa error path: UpdateEmpresa repo call fails
	repo.errors = map[string]error{"UpdateEmpresa": errors.New("db update error")}
	if _, err := s.UpdateEmpresa("test-emp", &UpdateEmpresaRequest{RazaoSocial: ptrString("Updated")}); err == nil {
		t.Fatal("expected error from UpdateEmpresa when repo update fails")
	}

	// Test UpdateEmpresa error path: Fetch after update fails
	repo.errors = map[string]error{"GetEmpresaByID": errors.New("db fetch error after update")}
	if _, err := s.UpdateEmpresa("test-emp", &UpdateEmpresaRequest{RazaoSocial: ptrString("Updated")}); err == nil {
		t.Fatal("expected error from UpdateEmpresa when fetch after update fails")
	}

	// Test DeleteEmpresa error path: Delete call fails
	repo.errors = map[string]error{"DeleteEmpresa": errors.New("db delete error")}
	if err := s.DeleteEmpresa("test-emp"); err == nil {
		t.Fatal("expected error from DeleteEmpresa when repo delete fails")
	}
}

func TestServiceEstabelecimentoErrorPaths(t *testing.T) {
	s, repo := setupService()

	// Set up empresa and estabelecimento
	repo.CreateEmpresa(&Empresa{
		ID: "emp1", CNPJ: "77777777000177", RazaoSocial: "E",
		TipoInscricao: "1", ClassificacaoTributaria: "01",
		IndCooperativa: "N", IndConstrutora: "N", IndDesoneracao: "N",
		Situacao: "ATIVA", DataCadastro: time.Now().UTC(), DataAtualizacao: time.Now().UTC(),
	})

	repo.CreateEstabelecimento(&Estabelecimento{
		ID: "est1", EmpresaID: "emp1", TipoInscricao: "1",
		NumeroInscricao: "12345", IndObra: "N",
		Situacao: "ATIVO", DataCadastro: time.Now().UTC(),
	})

	// Test GetEstabelecimentoByID error path
	repo.errors = map[string]error{"GetEstabelecimentoByID": errors.New("db error")}
	if _, err := s.GetEstabelecimentoByID("est1"); err == nil {
		t.Fatal("expected error from GetEstabelecimentoByID when repo fails")
	}

	// Test UpdateEstabelecimento error path: GetEstabelecimentoByID fails
	repo.errors = map[string]error{"GetEstabelecimentoByID": errors.New("db error")}
	if _, err := s.UpdateEstabelecimento("est1", &UpdateEstabelecimentoRequest{Situacao: ptrString("INATIVO")}); err == nil {
		t.Fatal("expected error from UpdateEstabelecimento when initial get fails")
	}

	// Test UpdateEstabelecimento error path: UpdateEstabelecimento fails
	repo.errors = map[string]error{"UpdateEstabelecimento": errors.New("db update error")}
	if _, err := s.UpdateEstabelecimento("est1", &UpdateEstabelecimentoRequest{Situacao: ptrString("INATIVO")}); err == nil {
		t.Fatal("expected error from UpdateEstabelecimento when repo update fails")
	}

	// Test DeleteEstabelecimento error path: GetEstabelecimentoByID fails
	repo.errors = map[string]error{"GetEstabelecimentoByID": errors.New("db error")}
	if err := s.DeleteEstabelecimento("est1"); err == nil {
		t.Fatal("expected error from DeleteEstabelecimento when get fails")
	}

	// Test DeleteEstabelecimento error path: DeleteEstabelecimento fails
	repo.errors = map[string]error{"DeleteEstabelecimento": errors.New("db delete error")}
	if err := s.DeleteEstabelecimento("est1"); err == nil {
		t.Fatal("expected error from DeleteEstabelecimento when delete fails")
	}
}

func TestServiceRubricaErrorPaths(t *testing.T) {
	s, repo := setupService()

	repo.CreateEmpresa(&Empresa{
		ID: "emp2", CNPJ: "88888888000188", RazaoSocial: "E2",
		TipoInscricao: "1", ClassificacaoTributaria: "01",
		IndCooperativa: "N", IndConstrutora: "N", IndDesoneracao: "N",
		Situacao: "ATIVA", DataCadastro: time.Now().UTC(), DataAtualizacao: time.Now().UTC(),
	})

	repo.CreateRubrica(&Rubrica{
		ID: "rub1", EmpresaID: "emp2", Codigo: "R01", Descricao: "desc",
		NaturezaRubrica: "0001", TipoRubrica: "D", Ativa: true,
		DataCadastro: time.Now().UTC(), DataInicioValidade: time.Now().UTC(),
	})

	// Test GetRubricaByID error path
	repo.errors = map[string]error{"GetRubricaByID": errors.New("db error")}
	if _, err := s.GetRubricaByID("rub1"); err == nil {
		t.Fatal("expected error from GetRubricaByID when repo fails")
	}

	// Test UpdateRubrica error path: GetRubricaByID fails initially
	repo.errors = map[string]error{"GetRubricaByID": errors.New("db error")}
	if _, err := s.UpdateRubrica("rub1", &UpdateRubricaRequest{Descricao: ptrString("new")}); err == nil {
		t.Fatal("expected error from UpdateRubrica when initial get fails")
	}

	// Test UpdateRubrica error path: UpdateRubrica fails
	repo.errors = map[string]error{"UpdateRubrica": errors.New("db update error")}
	if _, err := s.UpdateRubrica("rub1", &UpdateRubricaRequest{Descricao: ptrString("new")}); err == nil {
		t.Fatal("expected error from UpdateRubrica when repo update fails")
	}

	// Test DeleteRubrica error path: GetRubricaByID fails
	repo.errors = map[string]error{"GetRubricaByID": errors.New("db error")}
	if err := s.DeleteRubrica("rub1"); err == nil {
		t.Fatal("expected error from DeleteRubrica when get fails")
	}

	// Test DeleteRubrica error path: DeleteRubrica fails
	repo.errors = map[string]error{"DeleteRubrica": errors.New("db delete error")}
	if err := s.DeleteRubrica("rub1"); err == nil {
		t.Fatal("expected error from DeleteRubrica when delete fails")
	}
}

func TestServiceLotacaoErrorPaths(t *testing.T) {
	s, repo := setupService()

	repo.CreateEmpresa(&Empresa{
		ID: "emp3", CNPJ: "99999999000199", RazaoSocial: "E3",
		TipoInscricao: "1", ClassificacaoTributaria: "01",
		IndCooperativa: "N", IndConstrutora: "N", IndDesoneracao: "N",
		Situacao: "ATIVA", DataCadastro: time.Now().UTC(), DataAtualizacao: time.Now().UTC(),
	})

	repo.CreateLotacao(&Lotacao{
		ID: "lot1", EmpresaID: "emp3", Codigo: "L01", Descricao: "desc",
		TipoLotacao: "01", Ativa: true,
		DataCadastro: time.Now().UTC(), DataInicioValidade: time.Now().UTC(),
	})

	// Test GetLotacaoByID error path
	repo.errors = map[string]error{"GetLotacaoByID": errors.New("db error")}
	if _, err := s.GetLotacaoByID("lot1"); err == nil {
		t.Fatal("expected error from GetLotacaoByID when repo fails")
	}

	// Test UpdateLotacao error path: GetLotacaoByID fails initially
	repo.errors = map[string]error{"GetLotacaoByID": errors.New("db error")}
	if _, err := s.UpdateLotacao("lot1", &UpdateLotacaoRequest{Descricao: ptrString("new")}); err == nil {
		t.Fatal("expected error from UpdateLotacao when initial get fails")
	}

	// Test UpdateLotacao error path: UpdateLotacao fails
	repo.errors = map[string]error{"UpdateLotacao": errors.New("db update error")}
	if _, err := s.UpdateLotacao("lot1", &UpdateLotacaoRequest{Descricao: ptrString("new")}); err == nil {
		t.Fatal("expected error from UpdateLotacao when repo update fails")
	}

	// Test DeleteLotacao error path: GetLotacaoByID fails
	repo.errors = map[string]error{"GetLotacaoByID": errors.New("db error")}
	if err := s.DeleteLotacao("lot1"); err == nil {
		t.Fatal("expected error from DeleteLotacao when get fails")
	}

	// Test DeleteLotacao error path: DeleteLotacao fails
	repo.errors = map[string]error{"DeleteLotacao": errors.New("db delete error")}
	if err := s.DeleteLotacao("lot1"); err == nil {
		t.Fatal("expected error from DeleteLotacao when delete fails")
	}
}

func TestServiceCargoErrorPaths(t *testing.T) {
	s, repo := setupService()

	repo.CreateEmpresa(&Empresa{
		ID: "emp4", CNPJ: "11111111111111", RazaoSocial: "E4",
		TipoInscricao: "1", ClassificacaoTributaria: "01",
		IndCooperativa: "N", IndConstrutora: "N", IndDesoneracao: "N",
		Situacao: "ATIVA", DataCadastro: time.Now().UTC(), DataAtualizacao: time.Now().UTC(),
	})

	repo.CreateCargo(&Cargo{
		ID: "car1", EmpresaID: "emp4", Codigo: "C01", Descricao: "desc",
		CBO: "123456", Ativo: true,
		DataCadastro: time.Now().UTC(), DataInicioValidade: time.Now().UTC(),
	})

	// Test GetCargoByID error path
	repo.errors = map[string]error{"GetCargoByID": errors.New("db error")}
	if _, err := s.GetCargoByID("car1"); err == nil {
		t.Fatal("expected error from GetCargoByID when repo fails")
	}

	// Test UpdateCargo error path: GetCargoByID fails initially
	repo.errors = map[string]error{"GetCargoByID": errors.New("db error")}
	if _, err := s.UpdateCargo("car1", &UpdateCargoRequest{Descricao: ptrString("new")}); err == nil {
		t.Fatal("expected error from UpdateCargo when initial get fails")
	}

	// Test UpdateCargo error path: UpdateCargo fails
	repo.errors = map[string]error{"UpdateCargo": errors.New("db update error")}
	if _, err := s.UpdateCargo("car1", &UpdateCargoRequest{Descricao: ptrString("new")}); err == nil {
		t.Fatal("expected error from UpdateCargo when repo update fails")
	}

	// Test DeleteCargo error path: GetCargoByID fails
	repo.errors = map[string]error{"GetCargoByID": errors.New("db error")}
	if err := s.DeleteCargo("car1"); err == nil {
		t.Fatal("expected error from DeleteCargo when get fails")
	}

	// Test DeleteCargo error path: DeleteCargo fails
	repo.errors = map[string]error{"DeleteCargo": errors.New("db delete error")}
	if err := s.DeleteCargo("car1"); err == nil {
		t.Fatal("expected error from DeleteCargo when delete fails")
	}
}

func TestServiceCNPJValidation(t *testing.T) {
	s, _ := setupService()

	tests := []struct {
		name    string
		cnpj    string
		wantErr bool
	}{
		{"valid cnpj format", "11222333000181", false},
		{"invalid cnpj format too short", "12345", true},
		{"invalid cnpj text", "abcdefghijklmn", true},
		{"invalid cnpj contains letters", "1234567800a199", true},
		{"empty cnpj", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := s.CreateEmpresa(&CreateEmpresaRequest{
				CNPJ:                    tt.cnpj,
				RazaoSocial:             "Test",
				TipoInscricao:           "1",
				ClassificacaoTributaria: "01",
				IndCooperativa:          "N",
				IndConstrutora:          "N",
				IndDesoneracao:          "N",
			})
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateEmpresa() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
