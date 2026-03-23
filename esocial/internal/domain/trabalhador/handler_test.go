package trabalhador

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gorilla/mux"
)

type mockServiceForHandler struct {
	createFn   func(ctx context.Context, req *CreateTrabalhadorRequest) (*TrabalhadorResponse, error)
	getByIDFn  func(ctx context.Context, id string) (*TrabalhadorResponse, error)
	getByCPFFn func(ctx context.Context, cpf string) (*TrabalhadorResponse, error)
	updateFn   func(ctx context.Context, id string, req *UpdateTrabalhadorRequest) (*TrabalhadorResponse, error)
	deleteFn   func(ctx context.Context, id string) error
	listFn     func(ctx context.Context, req *ListRequest) (*ListResponse[TrabalhadorResponse], error)
}

func (m *mockServiceForHandler) Create(ctx context.Context, req *CreateTrabalhadorRequest) (*TrabalhadorResponse, error) {
	if m.createFn != nil {
		return m.createFn(ctx, req)
	}
	return nil, nil
}

func (m *mockServiceForHandler) GetByID(ctx context.Context, id string) (*TrabalhadorResponse, error) {
	if m.getByIDFn != nil {
		return m.getByIDFn(ctx, id)
	}
	return nil, nil
}

func (m *mockServiceForHandler) GetByCPF(ctx context.Context, cpf string) (*TrabalhadorResponse, error) {
	if m.getByCPFFn != nil {
		return m.getByCPFFn(ctx, cpf)
	}
	return nil, nil
}

func (m *mockServiceForHandler) Update(ctx context.Context, id string, req *UpdateTrabalhadorRequest) (*TrabalhadorResponse, error) {
	if m.updateFn != nil {
		return m.updateFn(ctx, id, req)
	}
	return nil, nil
}

func (m *mockServiceForHandler) Delete(ctx context.Context, id string) error {
	if m.deleteFn != nil {
		return m.deleteFn(ctx, id)
	}
	return nil
}

func (m *mockServiceForHandler) List(ctx context.Context, req *ListRequest) (*ListResponse[TrabalhadorResponse], error) {
	if m.listFn != nil {
		return m.listFn(ctx, req)
	}
	return nil, nil
}

// TestCreateTrabalhadorHandler tests the create endpoint
func TestCreateTrabalhadorHandler(t *testing.T) {
	t.Run("successful create", func(t *testing.T) {
		mockSvc := &mockServiceForHandler{
			createFn: func(ctx context.Context, req *CreateTrabalhadorRequest) (*TrabalhadorResponse, error) {
				return &TrabalhadorResponse{
					ID:        "t1",
					EmpresaID: "e1",
					CPF:       req.CPF,
					Nome:      req.Nome,
				}, nil
			},
		}
		handler := NewHandler(mockSvc)

		body := CreateTrabalhadorRequest{
			EmpresaID:         "e1",
			CPF:               "11144477735",
			Nome:              "João Silva",
			DataNascimento:    time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
			Sexo:              "M",
			PaisNascimento:    "BRA",
			PaisNacionalidade: "BRA",
		}
		bodyBytes, _ := json.Marshal(body)

		req := httptest.NewRequest("POST", "/trabalhadores", bytes.NewReader(bodyBytes))
		w := httptest.NewRecorder()

		handler.CreateTrabalhador(w, req)

		if w.Code != http.StatusCreated {
			t.Errorf("expected status %d, got %d", http.StatusCreated, w.Code)
		}

		var resp TrabalhadorResponse
		json.NewDecoder(w.Body).Decode(&resp)
		if resp.ID != "t1" {
			t.Errorf("expected ID t1, got %s", resp.ID)
		}
	})

	t.Run("invalid json", func(t *testing.T) {
		mockSvc := &mockServiceForHandler{}
		handler := NewHandler(mockSvc)

		req := httptest.NewRequest("POST", "/trabalhadores", bytes.NewReader([]byte("invalid")))
		w := httptest.NewRecorder()

		handler.CreateTrabalhador(w, req)

		if w.Code != http.StatusBadRequest {
			t.Errorf("expected status %d, got %d", http.StatusBadRequest, w.Code)
		}
	})

	t.Run("service error", func(t *testing.T) {
		mockSvc := &mockServiceForHandler{
			createFn: func(ctx context.Context, req *CreateTrabalhadorRequest) (*TrabalhadorResponse, error) {
				return nil, errors.New("database error")
			},
		}
		handler := NewHandler(mockSvc)

		body := CreateTrabalhadorRequest{
			EmpresaID: "e1",
			CPF:       "11144477735",
			Nome:      "João Silva",
		}
		bodyBytes, _ := json.Marshal(body)

		req := httptest.NewRequest("POST", "/trabalhadores", bytes.NewReader(bodyBytes))
		w := httptest.NewRecorder()

		handler.CreateTrabalhador(w, req)

		if w.Code != http.StatusBadRequest {
			t.Errorf("expected status %d, got %d", http.StatusBadRequest, w.Code)
		}
	})
}

// TestGetTrabalhadorHandler tests the get endpoint
func TestGetTrabalhadorHandler(t *testing.T) {
	t.Run("successful get", func(t *testing.T) {
		mockSvc := &mockServiceForHandler{
			getByIDFn: func(ctx context.Context, id string) (*TrabalhadorResponse, error) {
				return &TrabalhadorResponse{
					ID:        id,
					EmpresaID: "e1",
					CPF:       "11144477735",
					Nome:      "João Silva",
				}, nil
			},
		}
		handler := NewHandler(mockSvc)

		req := httptest.NewRequest("GET", "/trabalhadores/t1", nil)
		req = mux.SetURLVars(req, map[string]string{"id": "t1"})
		w := httptest.NewRecorder()

		handler.GetTrabalhador(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("expected status %d, got %d", http.StatusOK, w.Code)
		}

		var resp TrabalhadorResponse
		json.NewDecoder(w.Body).Decode(&resp)
		if resp.ID != "t1" {
			t.Errorf("expected ID t1, got %s", resp.ID)
		}
	})

	t.Run("not found", func(t *testing.T) {
		mockSvc := &mockServiceForHandler{
			getByIDFn: func(ctx context.Context, id string) (*TrabalhadorResponse, error) {
				return nil, errors.New("not found")
			},
		}
		handler := NewHandler(mockSvc)

		req := httptest.NewRequest("GET", "/trabalhadores/invalid", nil)
		req = mux.SetURLVars(req, map[string]string{"id": "invalid"})
		w := httptest.NewRecorder()

		handler.GetTrabalhador(w, req)

		if w.Code != http.StatusNotFound {
			t.Errorf("expected status %d, got %d", http.StatusNotFound, w.Code)
		}
	})
}

// TestGetTrabalhadorByCPFHandler tests the get by CPF endpoint
func TestGetTrabalhadorByCPFHandler(t *testing.T) {
	t.Run("successful get by cpf", func(t *testing.T) {
		mockSvc := &mockServiceForHandler{
			getByCPFFn: func(ctx context.Context, cpf string) (*TrabalhadorResponse, error) {
				if cpf == "11144477735" {
					return &TrabalhadorResponse{
						ID:  "t1",
						CPF: cpf,
					}, nil
				}
				return nil, errors.New("not found")
			},
		}
		handler := NewHandler(mockSvc)

		req := httptest.NewRequest("GET", "/trabalhadores/cpf/11144477735", nil)
		req = mux.SetURLVars(req, map[string]string{"cpf": "11144477735"})
		w := httptest.NewRecorder()

		handler.GetTrabalhadorByCPF(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("expected status %d, got %d", http.StatusOK, w.Code)
		}
	})

	t.Run("cpf not found", func(t *testing.T) {
		mockSvc := &mockServiceForHandler{
			getByCPFFn: func(ctx context.Context, cpf string) (*TrabalhadorResponse, error) {
				return nil, errors.New("not found")
			},
		}
		handler := NewHandler(mockSvc)

		req := httptest.NewRequest("GET", "/trabalhadores/cpf/00000000000", nil)
		req = mux.SetURLVars(req, map[string]string{"cpf": "00000000000"})
		w := httptest.NewRecorder()

		handler.GetTrabalhadorByCPF(w, req)

		if w.Code != http.StatusNotFound {
			t.Errorf("expected status %d, got %d", http.StatusNotFound, w.Code)
		}
	})
}

// TestUpdateTrabalhadorHandler tests the update endpoint
func TestUpdateTrabalhadorHandler(t *testing.T) {
	t.Run("successful update", func(t *testing.T) {
		mockSvc := &mockServiceForHandler{
			updateFn: func(ctx context.Context, id string, req *UpdateTrabalhadorRequest) (*TrabalhadorResponse, error) {
				return &TrabalhadorResponse{
					ID:              id,
					EmpresaID:       "e1",
					CPF:             "11144477735",
					DataAtualizacao: time.Now(),
				}, nil
			},
		}
		handler := NewHandler(mockSvc)

		body := UpdateTrabalhadorRequest{}
		bodyBytes, _ := json.Marshal(body)

		req := httptest.NewRequest("PUT", "/trabalhadores/t1", bytes.NewReader(bodyBytes))
		req = mux.SetURLVars(req, map[string]string{"id": "t1"})
		w := httptest.NewRecorder()

		handler.UpdateTrabalhador(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("expected status %d, got %d", http.StatusOK, w.Code)
		}

		var resp TrabalhadorResponse
		json.NewDecoder(w.Body).Decode(&resp)
		if resp.ID != "t1" {
			t.Errorf("expected ID t1, got %s", resp.ID)
		}
	})

	t.Run("invalid json", func(t *testing.T) {
		mockSvc := &mockServiceForHandler{}
		handler := NewHandler(mockSvc)

		req := httptest.NewRequest("PUT", "/trabalhadores/t1", bytes.NewReader([]byte("invalid")))
		req = mux.SetURLVars(req, map[string]string{"id": "t1"})
		w := httptest.NewRecorder()

		handler.UpdateTrabalhador(w, req)

		if w.Code != http.StatusBadRequest {
			t.Errorf("expected status %d, got %d", http.StatusBadRequest, w.Code)
		}
	})

	t.Run("service error", func(t *testing.T) {
		mockSvc := &mockServiceForHandler{
			updateFn: func(ctx context.Context, id string, req *UpdateTrabalhadorRequest) (*TrabalhadorResponse, error) {
				return nil, errors.New("validation error")
			},
		}
		handler := NewHandler(mockSvc)

		body := UpdateTrabalhadorRequest{}
		bodyBytes, _ := json.Marshal(body)

		req := httptest.NewRequest("PUT", "/trabalhadores/t1", bytes.NewReader(bodyBytes))
		req = mux.SetURLVars(req, map[string]string{"id": "t1"})
		w := httptest.NewRecorder()

		handler.UpdateTrabalhador(w, req)

		if w.Code != http.StatusBadRequest {
			t.Errorf("expected status %d, got %d", http.StatusBadRequest, w.Code)
		}
	})
}

// TestDeleteTrabalhadorHandler tests the delete endpoint
func TestDeleteTrabalhadorHandler(t *testing.T) {
	t.Run("successful delete", func(t *testing.T) {
		mockSvc := &mockServiceForHandler{
			deleteFn: func(ctx context.Context, id string) error {
				return nil
			},
		}
		handler := NewHandler(mockSvc)

		req := httptest.NewRequest("DELETE", "/trabalhadores/t1", nil)
		req = mux.SetURLVars(req, map[string]string{"id": "t1"})
		w := httptest.NewRecorder()

		handler.DeleteTrabalhador(w, req)

		if w.Code != http.StatusNoContent {
			t.Errorf("expected status %d, got %d", http.StatusNoContent, w.Code)
		}
	})

	t.Run("not found", func(t *testing.T) {
		mockSvc := &mockServiceForHandler{
			deleteFn: func(ctx context.Context, id string) error {
				return errors.New("not found")
			},
		}
		handler := NewHandler(mockSvc)

		req := httptest.NewRequest("DELETE", "/trabalhadores/invalid", nil)
		req = mux.SetURLVars(req, map[string]string{"id": "invalid"})
		w := httptest.NewRecorder()

		handler.DeleteTrabalhador(w, req)

		if w.Code != http.StatusNotFound {
			t.Errorf("expected status %d, got %d", http.StatusNotFound, w.Code)
		}
	})
}

// TestListTrabalhadorHandler tests the list endpoint
func TestListTrabalhadorHandler(t *testing.T) {
	t.Run("successful list", func(t *testing.T) {
		mockSvc := &mockServiceForHandler{
			listFn: func(ctx context.Context, req *ListRequest) (*ListResponse[TrabalhadorResponse], error) {
				return &ListResponse[TrabalhadorResponse]{
					Total:      2,
					Page:       1,
					Limit:      10,
					TotalPages: 1,
					Data: []TrabalhadorResponse{
						{ID: "t1", Nome: "João Silva"},
						{ID: "t2", Nome: "Maria Santos"},
					},
				}, nil
			},
		}
		handler := NewHandler(mockSvc)

		req := httptest.NewRequest("GET", "/trabalhadores?page=1&limit=10", nil)
		w := httptest.NewRecorder()

		handler.ListTrabalhadores(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("expected status %d, got %d", http.StatusOK, w.Code)
		}

		var resp ListResponse[TrabalhadorResponse]
		json.NewDecoder(w.Body).Decode(&resp)
		if resp.Total != 2 {
			t.Errorf("expected total 2, got %d", resp.Total)
		}
		if len(resp.Data) != 2 {
			t.Errorf("expected 2 items, got %d", len(resp.Data))
		}
	})

	t.Run("list with filters", func(t *testing.T) {
		mockSvc := &mockServiceForHandler{
			listFn: func(ctx context.Context, req *ListRequest) (*ListResponse[TrabalhadorResponse], error) {
				return &ListResponse[TrabalhadorResponse]{
					Total:      1,
					Page:       1,
					Limit:      10,
					TotalPages: 1,
					Data: []TrabalhadorResponse{
						{ID: "t1", Nome: "João Silva", EmpresaID: "e1"},
					},
				}, nil
			},
		}
		handler := NewHandler(mockSvc)

		req := httptest.NewRequest("GET", "/trabalhadores?page=1&limit=10&empresa_id=e1&sexo=M", nil)
		w := httptest.NewRecorder()

		handler.ListTrabalhadores(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("expected status %d, got %d", http.StatusOK, w.Code)
		}
	})

	t.Run("list with invalid pagination", func(t *testing.T) {
		mockSvc := &mockServiceForHandler{
			listFn: func(ctx context.Context, req *ListRequest) (*ListResponse[TrabalhadorResponse], error) {
				return &ListResponse[TrabalhadorResponse]{
					Total:      0,
					Page:       1,
					Limit:      10,
					TotalPages: 1,
					Data:       []TrabalhadorResponse{},
				}, nil
			},
		}
		handler := NewHandler(mockSvc)

		req := httptest.NewRequest("GET", "/trabalhadores?page=invalid&limit=invalid", nil)
		w := httptest.NewRecorder()

		handler.ListTrabalhadores(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("expected status %d, got %d", http.StatusOK, w.Code)
		}
	})

	t.Run("service error", func(t *testing.T) {
		mockSvc := &mockServiceForHandler{
			listFn: func(ctx context.Context, req *ListRequest) (*ListResponse[TrabalhadorResponse], error) {
				return nil, errors.New("database error")
			},
		}
		handler := NewHandler(mockSvc)

		req := httptest.NewRequest("GET", "/trabalhadores", nil)
		w := httptest.NewRecorder()

		handler.ListTrabalhadores(w, req)

		if w.Code != http.StatusInternalServerError {
			t.Errorf("expected status %d, got %d", http.StatusInternalServerError, w.Code)
		}
	})
}

// Helper to inject mux.Vars
func setURLVars(r *http.Request, vars map[string]string) *http.Request {
	return mux.SetURLVars(r, vars)
}
