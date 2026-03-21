package trabalhador

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Handler defines the HTTP handlers for trabalhador operations
type Handler struct {
	service Service
}

// NewHandler creates a new trabalhador handler
func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

// CreateTrabalhador handles POST /trabalhadores
func (h *Handler) CreateTrabalhador(w http.ResponseWriter, r *http.Request) {
	var req CreateTrabalhadorRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	response, err := h.service.Create(r.Context(), &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

// GetTrabalhador handles GET /trabalhadores/{id}
func (h *Handler) GetTrabalhador(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	response, err := h.service.GetByID(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// GetTrabalhadorByCPF handles GET /trabalhadores/cpf/{cpf}
func (h *Handler) GetTrabalhadorByCPF(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cpf := vars["cpf"]

	response, err := h.service.GetByCPF(r.Context(), cpf)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// UpdateTrabalhador handles PUT /trabalhadores/{id}
func (h *Handler) UpdateTrabalhador(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var req UpdateTrabalhadorRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	response, err := h.service.Update(r.Context(), id, &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// DeleteTrabalhador handles DELETE /trabalhadores/{id}
func (h *Handler) DeleteTrabalhador(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := h.service.Delete(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// ListTrabalhadores handles GET /trabalhadores
func (h *Handler) ListTrabalhadores(w http.ResponseWriter, r *http.Request) {
	// Parse query parameters
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

	if situacao := r.URL.Query().Get("situacao"); situacao != "" {
		req.Situacao = &situacao
	}

	if sexo := r.URL.Query().Get("sexo"); sexo != "" {
		req.Sexo = &sexo
	}

	if possuiDeficienciaStr := r.URL.Query().Get("possui_deficiencia"); possuiDeficienciaStr != "" {
		if possuiDeficiencia, err := strconv.ParseBool(possuiDeficienciaStr); err == nil {
			req.PossuiDeficiencia = &possuiDeficiencia
		}
	}

	response, err := h.service.List(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// RegisterRoutes registers the trabalhador routes
func (h *Handler) RegisterRoutes(router *mux.Router) {
	trabalhadorRouter := router.PathPrefix("/trabalhadores").Subrouter()

	trabalhadorRouter.HandleFunc("", h.CreateTrabalhador).Methods("POST")
	trabalhadorRouter.HandleFunc("", h.ListTrabalhadores).Methods("GET")
	trabalhadorRouter.HandleFunc("/{id}", h.GetTrabalhador).Methods("GET")
	trabalhadorRouter.HandleFunc("/{id}", h.UpdateTrabalhador).Methods("PUT")
	trabalhadorRouter.HandleFunc("/{id}", h.DeleteTrabalhador).Methods("DELETE")
	trabalhadorRouter.HandleFunc("/cpf/{cpf}", h.GetTrabalhadorByCPF).Methods("GET")
}
