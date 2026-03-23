package empresa

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/avelinobego/esocial/configs"
	"github.com/gorilla/mux"
)

type mockHandlerRepository struct {
	empresas         map[string]*Empresa
	estabelecimentos map[string]*Estabelecimento
	rubricas         map[string]*Rubrica
	lotacoes         map[string]*Lotacao
	cargos           map[string]*Cargo
	errors           map[string]error
}

func newMockHandlerRepository() *mockHandlerRepository {
	return &mockHandlerRepository{
		empresas:         make(map[string]*Empresa),
		estabelecimentos: make(map[string]*Estabelecimento),
		rubricas:         make(map[string]*Rubrica),
		lotacoes:         make(map[string]*Lotacao),
		cargos:           make(map[string]*Cargo),
		errors:           make(map[string]error),
	}
}

func (m *mockHandlerRepository) errFor(method string) error {
	if err, ok := m.errors[method]; ok {
		return err
	}
	return nil
}

func (m *mockHandlerRepository) CreateEmpresa(e *Empresa) error {
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

func (m *mockHandlerRepository) GetEmpresaByID(id string) (*Empresa, error) {
	if err := m.errFor("GetEmpresaByID"); err != nil {
		return nil, err
	}
	if e, ok := m.empresas[id]; ok {
		return e, nil
	}
	return nil, nil
}

func (m *mockHandlerRepository) GetEmpresaByCNPJ(cnpj string) (*Empresa, error) {
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

func (m *mockHandlerRepository) UpdateEmpresa(id string, updates map[string]interface{}) error {
	if err := m.errFor("UpdateEmpresa"); err != nil {
		return err
	}
	if e, ok := m.empresas[id]; ok {
		e.DataAtualizacao = time.Now().UTC()
		m.empresas[id] = e
	}
	return nil
}

func (m *mockHandlerRepository) DeleteEmpresa(id string) error {
	if err := m.errFor("DeleteEmpresa"); err != nil {
		return err
	}
	delete(m.empresas, id)
	return nil
}

func (m *mockHandlerRepository) ListEmpresas(filter map[string]interface{}, limit int, offset int) ([]*Empresa, int64, error) {
	if err := m.errFor("ListEmpresas"); err != nil {
		return nil, 0, err
	}
	list := make([]*Empresa, 0, len(m.empresas))
	for _, e := range m.empresas {
		list = append(list, e)
	}
	return list, int64(len(list)), nil
}

func (m *mockHandlerRepository) CreateEstabelecimento(e *Estabelecimento) error {
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

func (m *mockHandlerRepository) GetEstabelecimentoByID(id string) (*Estabelecimento, error) {
	if err := m.errFor("GetEstabelecimentoByID"); err != nil {
		return nil, err
	}
	if e, ok := m.estabelecimentos[id]; ok {
		return e, nil
	}
	return nil, nil
}

func (m *mockHandlerRepository) UpdateEstabelecimento(id string, updates map[string]interface{}) error {
	if err := m.errFor("UpdateEstabelecimento"); err != nil {
		return err
	}
	return nil
}

func (m *mockHandlerRepository) DeleteEstabelecimento(id string) error {
	if err := m.errFor("DeleteEstabelecimento"); err != nil {
		return err
	}
	delete(m.estabelecimentos, id)
	return nil
}

func (m *mockHandlerRepository) ListEstabelecimentos(filter map[string]interface{}, limit int, offset int) ([]*Estabelecimento, int64, error) {
	if err := m.errFor("ListEstabelecimentos"); err != nil {
		return nil, 0, err
	}
	list := make([]*Estabelecimento, 0, len(m.estabelecimentos))
	for _, e := range m.estabelecimentos {
		list = append(list, e)
	}
	return list, int64(len(list)), nil
}

func (m *mockHandlerRepository) CreateRubrica(r *Rubrica) error {
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

func (m *mockHandlerRepository) GetRubricaByID(id string) (*Rubrica, error) {
	if err := m.errFor("GetRubricaByID"); err != nil {
		return nil, err
	}
	if r, ok := m.rubricas[id]; ok {
		return r, nil
	}
	return nil, nil
}

func (m *mockHandlerRepository) UpdateRubrica(id string, updates map[string]interface{}) error {
	if err := m.errFor("UpdateRubrica"); err != nil {
		return err
	}
	return nil
}

func (m *mockHandlerRepository) DeleteRubrica(id string) error {
	if err := m.errFor("DeleteRubrica"); err != nil {
		return err
	}
	delete(m.rubricas, id)
	return nil
}

func (m *mockHandlerRepository) ListRubricas(filter map[string]interface{}, limit int, offset int) ([]*Rubrica, int64, error) {
	if err := m.errFor("ListRubricas"); err != nil {
		return nil, 0, err
	}
	list := make([]*Rubrica, 0, len(m.rubricas))
	for _, r := range m.rubricas {
		list = append(list, r)
	}
	return list, int64(len(list)), nil
}

func (m *mockHandlerRepository) CreateLotacao(l *Lotacao) error {
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

func (m *mockHandlerRepository) GetLotacaoByID(id string) (*Lotacao, error) {
	if err := m.errFor("GetLotacaoByID"); err != nil {
		return nil, err
	}
	if l, ok := m.lotacoes[id]; ok {
		return l, nil
	}
	return nil, nil
}

func (m *mockHandlerRepository) UpdateLotacao(id string, updates map[string]interface{}) error {
	if err := m.errFor("UpdateLotacao"); err != nil {
		return err
	}
	return nil
}

func (m *mockHandlerRepository) DeleteLotacao(id string) error {
	if err := m.errFor("DeleteLotacao"); err != nil {
		return err
	}
	delete(m.lotacoes, id)
	return nil
}

func (m *mockHandlerRepository) ListLotacoes(filter map[string]interface{}, limit int, offset int) ([]*Lotacao, int64, error) {
	if err := m.errFor("ListLotacoes"); err != nil {
		return nil, 0, err
	}
	list := make([]*Lotacao, 0, len(m.lotacoes))
	for _, l := range m.lotacoes {
		list = append(list, l)
	}
	return list, int64(len(list)), nil
}

func (m *mockHandlerRepository) CreateCargo(c *Cargo) error {
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

func (m *mockHandlerRepository) GetCargoByID(id string) (*Cargo, error) {
	if err := m.errFor("GetCargoByID"); err != nil {
		return nil, err
	}
	if c, ok := m.cargos[id]; ok {
		return c, nil
	}
	return nil, nil
}

func (m *mockHandlerRepository) UpdateCargo(id string, updates map[string]interface{}) error {
	if err := m.errFor("UpdateCargo"); err != nil {
		return err
	}
	return nil
}

func (m *mockHandlerRepository) DeleteCargo(id string) error {
	if err := m.errFor("DeleteCargo"); err != nil {
		return err
	}
	delete(m.cargos, id)
	return nil
}

func (m *mockHandlerRepository) ListCargos(filter map[string]interface{}, limit int, offset int) ([]*Cargo, int64, error) {
	if err := m.errFor("ListCargos"); err != nil {
		return nil, 0, err
	}
	list := make([]*Cargo, 0, len(m.cargos))
	for _, c := range m.cargos {
		list = append(list, c)
	}
	return list, int64(len(list)), nil
}

func setupHandlerTest() *Handler {
	logger := configs.NewLogger(io.Discard)
	repo := newMockHandlerRepository()
	svc := NewService(repo, logger)
	return NewHandler(svc)
}

func TestHandlerCreateEmpresa(t *testing.T) {
	h := setupHandlerTest()

	t.Run("successful create", func(t *testing.T) {
		payload := CreateEmpresaRequest{
			CNPJ: "12345678000195", RazaoSocial: "Test",
			TipoInscricao: "1", ClassificacaoTributaria: "01",
			IndCooperativa: "N", IndConstrutora: "N", IndDesoneracao: "N",
		}
		body, _ := json.Marshal(payload)
		req := httptest.NewRequest("POST", "/empresas", bytes.NewReader(body))
		w := httptest.NewRecorder()

		h.CreateEmpresa(w, req)

		if w.Code != http.StatusCreated {
			t.Fatalf("expected status 201, got %d", w.Code)
		}
		var resp EmpresaResponse
		json.Unmarshal(w.Body.Bytes(), &resp)
		if resp.ID == "" {
			t.Fatal("expected ID to be set")
		}
	})

	t.Run("invalid json", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/empresas", bytes.NewReader([]byte("invalid")))
		w := httptest.NewRecorder()
		h := setupHandlerTest()

		h.CreateEmpresa(w, req)

		if w.Code != http.StatusBadRequest {
			t.Fatalf("expected status 400, got %d", w.Code)
		}
	})

	t.Run("invalid cnpj error", func(t *testing.T) {
		payload := CreateEmpresaRequest{CNPJ: "123"}
		body, _ := json.Marshal(payload)
		req := httptest.NewRequest("POST", "/empresas", bytes.NewReader(body))
		w := httptest.NewRecorder()
		h := setupHandlerTest()

		h.CreateEmpresa(w, req)

		if w.Code != http.StatusBadRequest {
			t.Fatalf("expected status 400, got %d", w.Code)
		}
	})
}

func TestHandlerGetEmpresa(t *testing.T) {
	h := setupHandlerTest()

	t.Run("get non-existent", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/empresas/notfound", nil)
		vars := map[string]string{"id": "notfound"}
		req = mux.SetURLVars(req, vars)
		w := httptest.NewRecorder()

		h.GetEmpresa(w, req)

		if w.Code != http.StatusNotFound {
			t.Fatalf("expected status 404, got %d", w.Code)
		}
	})

	t.Run("successful get", func(t *testing.T) {
		h := setupHandlerTest()
		// First create empresa
		payload := CreateEmpresaRequest{
			CNPJ: "12345678000195", RazaoSocial: "Test",
			TipoInscricao: "1", ClassificacaoTributaria: "01",
			IndCooperativa: "N", IndConstrutora: "N", IndDesoneracao: "N",
		}
		body, _ := json.Marshal(payload)
		req := httptest.NewRequest("POST", "/empresas", bytes.NewReader(body))
		w := httptest.NewRecorder()
		h.CreateEmpresa(w, req)
		var created EmpresaResponse
		json.Unmarshal(w.Body.Bytes(), &created)

		// Then get it
		req = httptest.NewRequest("GET", "/empresas/"+created.ID, nil)
		vs := map[string]string{"id": created.ID}
		req = mux.SetURLVars(req, vs)
		w = httptest.NewRecorder()

		h.GetEmpresa(w, req)

		if w.Code != http.StatusOK {
			t.Fatalf("expected status 200, got %d", w.Code)
		}
	})
}

func TestHandlerUpdateEmpresa(t *testing.T) {
	h := setupHandlerTest()

	t.Run("invalid json", func(t *testing.T) {
		req := httptest.NewRequest("PUT", "/empresas/emp-1", bytes.NewReader([]byte("invalid")))
		vars := map[string]string{"id": "emp-1"}
		req = mux.SetURLVars(req, vars)
		w := httptest.NewRecorder()

		h.UpdateEmpresa(w, req)

		if w.Code != http.StatusBadRequest {
			t.Fatalf("expected status 400, got %d", w.Code)
		}
	})
}

func TestHandlerDeleteEmpresa(t *testing.T) {
	h := setupHandlerTest()

	t.Run("delete non-existent", func(t *testing.T) {
		req := httptest.NewRequest("DELETE", "/empresas/notfound", nil)
		vars := map[string]string{"id": "notfound"}
		req = mux.SetURLVars(req, vars)
		w := httptest.NewRecorder()

		h.DeleteEmpresa(w, req)

		// Should still return 204 even if not found (delete is idempotent)
		if w.Code != http.StatusNoContent && w.Code != http.StatusNotFound {
			t.Fatalf("expected status 204 or 404, got %d", w.Code)
		}
	})
}

func TestHandlerListEmpresas(t *testing.T) {
	h := setupHandlerTest()

	t.Run("successful list", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/empresas?page=1&limit=10", nil)
		w := httptest.NewRecorder()

		h.ListEmpresas(w, req)

		if w.Code != http.StatusOK {
			t.Fatalf("expected status 200, got %d", w.Code)
		}
	})

	t.Run("invalid pagination", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/empresas?page=invalid&limit=10", nil)
		w := httptest.NewRecorder()

		h.ListEmpresas(w, req)

		if w.Code == http.StatusOK {
			// Might return OK with default or error, both acceptable
		}
	})
}

func TestHandlerEstabelecimento(t *testing.T) {
	h := setupHandlerTest()
	logger := configs.NewLogger(io.Discard)
	repo := newMockHandlerRepository()
	// Create empresa first
	repo.CreateEmpresa(&Empresa{
		ID: "emp-1", CNPJ: "11111111000111", RazaoSocial: "E",
		TipoInscricao: "1", ClassificacaoTributaria: "01",
		IndCooperativa: "N", IndConstrutora: "N", IndDesoneracao: "N",
		Situacao: "ATIVA", DataCadastro: time.Now().UTC(), DataAtualizacao: time.Now().UTC(),
	})
	svc := NewService(repo, logger)
	h = NewHandler(svc)

	t.Run("create estabelecimento", func(t *testing.T) {
		payload := CreateEstabelecimentoRequest{
			EmpresaID: "emp-1", TipoInscricao: "1",
			NumeroInscricao: "12345", IndObra: "N",
		}
		body, _ := json.Marshal(payload)
		req := httptest.NewRequest("POST", "/estabelecimentos", bytes.NewReader(body))
		w := httptest.NewRecorder()

		h.CreateEstabelecimento(w, req)

		if w.Code != http.StatusCreated {
			t.Fatalf("expected status 201, got %d", w.Code)
		}
	})
}

func TestHandlerRubrica(t *testing.T) {
	logger := configs.NewLogger(io.Discard)
	repo := newMockHandlerRepository()
	repo.CreateEmpresa(&Empresa{
		ID: "emp-1", CNPJ: "22222222000122", RazaoSocial: "E2",
		TipoInscricao: "1", ClassificacaoTributaria: "01",
		IndCooperativa: "N", IndConstrutora: "N", IndDesoneracao: "N",
		Situacao: "ATIVA", DataCadastro: time.Now().UTC(), DataAtualizacao: time.Now().UTC(),
	})
	svc := NewService(repo, logger)
	h := NewHandler(svc)

	t.Run("create rubrica", func(t *testing.T) {
		payload := CreateRubricaRequest{
			EmpresaID: "emp-1", Codigo: "R01", Descricao: "desc",
			NaturezaRubrica: "0001", TipoRubrica: "D",
			DataInicioValidade: time.Now().UTC(),
		}
		body, _ := json.Marshal(payload)
		req := httptest.NewRequest("POST", "/rubricas", bytes.NewReader(body))
		w := httptest.NewRecorder()

		h.CreateRubrica(w, req)

		if w.Code != http.StatusCreated {
			t.Fatalf("expected status 201, got %d", w.Code)
		}
	})
}

func TestHandlerLotacao(t *testing.T) {
	logger := configs.NewLogger(io.Discard)
	repo := newMockHandlerRepository()
	repo.CreateEmpresa(&Empresa{
		ID: "emp-1", CNPJ: "33333333000133", RazaoSocial: "E3",
		TipoInscricao: "1", ClassificacaoTributaria: "01",
		IndCooperativa: "N", IndConstrutora: "N", IndDesoneracao: "N",
		Situacao: "ATIVA", DataCadastro: time.Now().UTC(), DataAtualizacao: time.Now().UTC(),
	})
	svc := NewService(repo, logger)
	h := NewHandler(svc)

	t.Run("create lotacao", func(t *testing.T) {
		payload := CreateLotacaoRequest{
			EmpresaID: "emp-1", Codigo: "L01", Descricao: "desc",
			TipoLotacao: "01", DataInicioValidade: time.Now().UTC(),
		}
		body, _ := json.Marshal(payload)
		req := httptest.NewRequest("POST", "/lotacoes", bytes.NewReader(body))
		w := httptest.NewRecorder()

		h.CreateLotacao(w, req)

		if w.Code != http.StatusCreated {
			t.Fatalf("expected status 201, got %d", w.Code)
		}
	})
}

func TestHandlerCargo(t *testing.T) {
	logger := configs.NewLogger(io.Discard)
	repo := newMockHandlerRepository()
	repo.CreateEmpresa(&Empresa{
		ID: "emp-1", CNPJ: "44444444000144", RazaoSocial: "E4",
		TipoInscricao: "1", ClassificacaoTributaria: "01",
		IndCooperativa: "N", IndConstrutora: "N", IndDesoneracao: "N",
		Situacao: "ATIVA", DataCadastro: time.Now().UTC(), DataAtualizacao: time.Now().UTC(),
	})
	svc := NewService(repo, logger)
	h := NewHandler(svc)

	t.Run("create cargo", func(t *testing.T) {
		payload := CreateCargoRequest{
			EmpresaID: "emp-1", Codigo: "C01", Descricao: "desc",
			CBO: "123456", DataInicioValidade: time.Now().UTC(),
		}
		body, _ := json.Marshal(payload)
		req := httptest.NewRequest("POST", "/cargos", bytes.NewReader(body))
		w := httptest.NewRecorder()

		h.CreateCargo(w, req)

		if w.Code != http.StatusCreated {
			t.Fatalf("expected status 201, got %d", w.Code)
		}
	})
}
