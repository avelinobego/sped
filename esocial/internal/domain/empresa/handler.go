package empresa

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Handler defines the HTTP handlers for empresa operations
type Handler struct {
	service *Service
}

// NewHandler creates a new empresa handler
func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

// Empresa Handlers

// CreateEmpresa handles POST /empresas
func (h *Handler) CreateEmpresa(w http.ResponseWriter, r *http.Request) {
	var req CreateEmpresaRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	response, err := h.service.CreateEmpresa(&req)
	if err != nil {
		if err == ErrInvalidCNPJ || err == ErrEmpresaAlreadyExists {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

// GetEmpresa handles GET /empresas/{id}
func (h *Handler) GetEmpresa(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	response, err := h.service.GetEmpresaByID(id)
	if err != nil {
		if err == ErrEmpresaNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// UpdateEmpresa handles PUT /empresas/{id}
func (h *Handler) UpdateEmpresa(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var req UpdateEmpresaRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	response, err := h.service.UpdateEmpresa(id, &req)
	if err != nil {
		if err == ErrEmpresaNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// DeleteEmpresa handles DELETE /empresas/{id}
func (h *Handler) DeleteEmpresa(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := h.service.DeleteEmpresa(id)
	if err != nil {
		if err == ErrEmpresaNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// ListEmpresas handles GET /empresas
func (h *Handler) ListEmpresas(w http.ResponseWriter, r *http.Request) {
	req := &ListRequest{}

	if pageStr := r.URL.Query().Get("page"); pageStr != "" {
		if page, err := strconv.Atoi(pageStr); err == nil {
			req.Page = page
		}
	}

	if limitStr := r.URL.Query().Get("limit"); limitStr != "" {
		if limit, err := strconv.Atoi(limitStr); err == nil {
			req.Limit = limit
		}
	}

	if empresaID := r.URL.Query().Get("empresa_id"); empresaID != "" {
		req.EmpresaID = &empresaID
	}

	if search := r.URL.Query().Get("search"); search != "" {
		req.Search = &search
	}

	if ativoStr := r.URL.Query().Get("ativo"); ativoStr != "" {
		if ativo, err := strconv.ParseBool(ativoStr); err == nil {
			req.Ativo = &ativo
		}
	}

	response, err := h.service.ListEmpresas(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Estabelecimento Handlers

// CreateEstabelecimento handles POST /estabelecimentos
func (h *Handler) CreateEstabelecimento(w http.ResponseWriter, r *http.Request) {
	var req CreateEstabelecimentoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	response, err := h.service.CreateEstabelecimento(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

// GetEstabelecimento handles GET /estabelecimentos/{id}
func (h *Handler) GetEstabelecimento(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	response, err := h.service.GetEstabelecimentoByID(id)
	if err != nil {
		if err == ErrEstabelecimentoNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// UpdateEstabelecimento handles PUT /estabelecimentos/{id}
func (h *Handler) UpdateEstabelecimento(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var req UpdateEstabelecimentoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	response, err := h.service.UpdateEstabelecimento(id, &req)
	if err != nil {
		if err == ErrEstabelecimentoNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// DeleteEstabelecimento handles DELETE /estabelecimentos/{id}
func (h *Handler) DeleteEstabelecimento(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := h.service.DeleteEstabelecimento(id)
	if err != nil {
		if err == ErrEstabelecimentoNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// ListEstabelecimentos handles GET /estabelecimentos
func (h *Handler) ListEstabelecimentos(w http.ResponseWriter, r *http.Request) {
	req := &ListRequest{}

	if pageStr := r.URL.Query().Get("page"); pageStr != "" {
		if page, err := strconv.Atoi(pageStr); err == nil {
			req.Page = page
		}
	}

	if limitStr := r.URL.Query().Get("limit"); limitStr != "" {
		if limit, err := strconv.Atoi(limitStr); err == nil {
			req.Limit = limit
		}
	}

	if empresaID := r.URL.Query().Get("empresa_id"); empresaID != "" {
		req.EmpresaID = &empresaID
	}

	if search := r.URL.Query().Get("search"); search != "" {
		req.Search = &search
	}

	if ativoStr := r.URL.Query().Get("ativo"); ativoStr != "" {
		if ativo, err := strconv.ParseBool(ativoStr); err == nil {
			req.Ativo = &ativo
		}
	}

	response, err := h.service.ListEstabelecimentos(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Rubrica Handlers

// CreateRubrica handles POST /rubricas
func (h *Handler) CreateRubrica(w http.ResponseWriter, r *http.Request) {
	var req CreateRubricaRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	response, err := h.service.CreateRubrica(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

// GetRubrica handles GET /rubricas/{id}
func (h *Handler) GetRubrica(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	response, err := h.service.GetRubricaByID(id)
	if err != nil {
		if err == ErrRubricaNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// UpdateRubrica handles PUT /rubricas/{id}
func (h *Handler) UpdateRubrica(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var req UpdateRubricaRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	response, err := h.service.UpdateRubrica(id, &req)
	if err != nil {
		if err == ErrRubricaNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// DeleteRubrica handles DELETE /rubricas/{id}
func (h *Handler) DeleteRubrica(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := h.service.DeleteRubrica(id)
	if err != nil {
		if err == ErrRubricaNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// ListRubricas handles GET /rubricas
func (h *Handler) ListRubricas(w http.ResponseWriter, r *http.Request) {
	req := &ListRequest{}

	if pageStr := r.URL.Query().Get("page"); pageStr != "" {
		if page, err := strconv.Atoi(pageStr); err == nil {
			req.Page = page
		}
	}

	if limitStr := r.URL.Query().Get("limit"); limitStr != "" {
		if limit, err := strconv.Atoi(limitStr); err == nil {
			req.Limit = limit
		}
	}

	if empresaID := r.URL.Query().Get("empresa_id"); empresaID != "" {
		req.EmpresaID = &empresaID
	}

	if search := r.URL.Query().Get("search"); search != "" {
		req.Search = &search
	}

	if ativoStr := r.URL.Query().Get("ativo"); ativoStr != "" {
		if ativo, err := strconv.ParseBool(ativoStr); err == nil {
			req.Ativo = &ativo
		}
	}

	response, err := h.service.ListRubricas(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Lotacao Handlers

// CreateLotacao handles POST /lotacoes
func (h *Handler) CreateLotacao(w http.ResponseWriter, r *http.Request) {
	var req CreateLotacaoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	response, err := h.service.CreateLotacao(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

// GetLotacao handles GET /lotacoes/{id}
func (h *Handler) GetLotacao(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	response, err := h.service.GetLotacaoByID(id)
	if err != nil {
		if err == ErrLotacaoNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// UpdateLotacao handles PUT /lotacoes/{id}
func (h *Handler) UpdateLotacao(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var req UpdateLotacaoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	response, err := h.service.UpdateLotacao(id, &req)
	if err != nil {
		if err == ErrLotacaoNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// DeleteLotacao handles DELETE /lotacoes/{id}
func (h *Handler) DeleteLotacao(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := h.service.DeleteLotacao(id)
	if err != nil {
		if err == ErrLotacaoNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// ListLotacoes handles GET /lotacoes
func (h *Handler) ListLotacoes(w http.ResponseWriter, r *http.Request) {
	req := &ListRequest{}

	if pageStr := r.URL.Query().Get("page"); pageStr != "" {
		if page, err := strconv.Atoi(pageStr); err == nil {
			req.Page = page
		}
	}

	if limitStr := r.URL.Query().Get("limit"); limitStr != "" {
		if limit, err := strconv.Atoi(limitStr); err == nil {
			req.Limit = limit
		}
	}

	if empresaID := r.URL.Query().Get("empresa_id"); empresaID != "" {
		req.EmpresaID = &empresaID
	}

	if search := r.URL.Query().Get("search"); search != "" {
		req.Search = &search
	}

	if ativoStr := r.URL.Query().Get("ativo"); ativoStr != "" {
		if ativo, err := strconv.ParseBool(ativoStr); err == nil {
			req.Ativo = &ativo
		}
	}

	response, err := h.service.ListLotacoes(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Cargo Handlers

// CreateCargo handles POST /cargos
func (h *Handler) CreateCargo(w http.ResponseWriter, r *http.Request) {
	var req CreateCargoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	response, err := h.service.CreateCargo(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

// GetCargo handles GET /cargos/{id}
func (h *Handler) GetCargo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	response, err := h.service.GetCargoByID(id)
	if err != nil {
		if err == ErrCargoNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// UpdateCargo handles PUT /cargos/{id}
func (h *Handler) UpdateCargo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var req UpdateCargoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	response, err := h.service.UpdateCargo(id, &req)
	if err != nil {
		if err == ErrCargoNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// DeleteCargo handles DELETE /cargos/{id}
func (h *Handler) DeleteCargo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := h.service.DeleteCargo(id)
	if err != nil {
		if err == ErrCargoNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// ListCargos handles GET /cargos
func (h *Handler) ListCargos(w http.ResponseWriter, r *http.Request) {
	req := &ListRequest{}

	if pageStr := r.URL.Query().Get("page"); pageStr != "" {
		if page, err := strconv.Atoi(pageStr); err == nil {
			req.Page = page
		}
	}

	if limitStr := r.URL.Query().Get("limit"); limitStr != "" {
		if limit, err := strconv.Atoi(limitStr); err == nil {
			req.Limit = limit
		}
	}

	if empresaID := r.URL.Query().Get("empresa_id"); empresaID != "" {
		req.EmpresaID = &empresaID
	}

	if search := r.URL.Query().Get("search"); search != "" {
		req.Search = &search
	}

	if ativoStr := r.URL.Query().Get("ativo"); ativoStr != "" {
		if ativo, err := strconv.ParseBool(ativoStr); err == nil {
			req.Ativo = &ativo
		}
	}

	response, err := h.service.ListCargos(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// RegisterRoutes registers all empresa routes
func (h *Handler) RegisterRoutes(router *mux.Router) {
	// Empresa routes
	empresaRouter := router.PathPrefix("/empresas").Subrouter()
	empresaRouter.HandleFunc("", h.CreateEmpresa).Methods("POST")
	empresaRouter.HandleFunc("", h.ListEmpresas).Methods("GET")
	empresaRouter.HandleFunc("/{id}", h.GetEmpresa).Methods("GET")
	empresaRouter.HandleFunc("/{id}", h.UpdateEmpresa).Methods("PUT")
	empresaRouter.HandleFunc("/{id}", h.DeleteEmpresa).Methods("DELETE")

	// Estabelecimento routes
	estabelecimentoRouter := router.PathPrefix("/estabelecimentos").Subrouter()
	estabelecimentoRouter.HandleFunc("", h.CreateEstabelecimento).Methods("POST")
	estabelecimentoRouter.HandleFunc("", h.ListEstabelecimentos).Methods("GET")
	estabelecimentoRouter.HandleFunc("/{id}", h.GetEstabelecimento).Methods("GET")
	estabelecimentoRouter.HandleFunc("/{id}", h.UpdateEstabelecimento).Methods("PUT")
	estabelecimentoRouter.HandleFunc("/{id}", h.DeleteEstabelecimento).Methods("DELETE")

	// Rubrica routes
	rubricaRouter := router.PathPrefix("/rubricas").Subrouter()
	rubricaRouter.HandleFunc("", h.CreateRubrica).Methods("POST")
	rubricaRouter.HandleFunc("", h.ListRubricas).Methods("GET")
	rubricaRouter.HandleFunc("/{id}", h.GetRubrica).Methods("GET")
	rubricaRouter.HandleFunc("/{id}", h.UpdateRubrica).Methods("PUT")
	rubricaRouter.HandleFunc("/{id}", h.DeleteRubrica).Methods("DELETE")

	// Lotacao routes
	lotacaoRouter := router.PathPrefix("/lotacoes").Subrouter()
	lotacaoRouter.HandleFunc("", h.CreateLotacao).Methods("POST")
	lotacaoRouter.HandleFunc("", h.ListLotacoes).Methods("GET")
	lotacaoRouter.HandleFunc("/{id}", h.GetLotacao).Methods("GET")
	lotacaoRouter.HandleFunc("/{id}", h.UpdateLotacao).Methods("PUT")
	lotacaoRouter.HandleFunc("/{id}", h.DeleteLotacao).Methods("DELETE")

	// Cargo routes
	cargoRouter := router.PathPrefix("/cargos").Subrouter()
	cargoRouter.HandleFunc("", h.CreateCargo).Methods("POST")
	cargoRouter.HandleFunc("", h.ListCargos).Methods("GET")
	cargoRouter.HandleFunc("/{id}", h.GetCargo).Methods("GET")
	cargoRouter.HandleFunc("/{id}", h.UpdateCargo).Methods("PUT")
	cargoRouter.HandleFunc("/{id}", h.DeleteCargo).Methods("DELETE")
}
