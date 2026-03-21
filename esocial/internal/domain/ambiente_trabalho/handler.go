package ambiente_trabalho

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/avelinobego/esocial/configs"
	"github.com/gorilla/mux"
)

type ServiceInterface interface {
	// Horario methods
	CreateHorario(*CreateHorarioRequest) (*HorarioResponse, error)
	GetHorarioByID(string) (*HorarioResponse, error)
	UpdateHorario(string, *UpdateHorarioRequest) (*HorarioResponse, error)
	DeleteHorario(string) error
	ListHorarios(*ListRequest) (*ListResponse[HorarioResponse], error)

	// HorarioDetalhe methods
	CreateHorarioDetalhe(*CreateHorarioDetalheRequest) (*HorarioDetalheResponse, error)
	UpdateHorarioDetalhe(string, *UpdateHorarioDetalheRequest) (*HorarioDetalheResponse, error)
	DeleteHorarioDetalhe(string) error

	// AmbienteTrabalho methods
	CreateAmbienteTrabalho(*CreateAmbienteTrabalhoRequest) (*AmbienteTrabalhoResponse, error)
	GetAmbienteTrabalhoByID(string) (*AmbienteTrabalhoResponse, error)
	UpdateAmbienteTrabalho(string, *UpdateAmbienteTrabalhoRequest) (*AmbienteTrabalhoResponse, error)
	DeleteAmbienteTrabalho(string) error
	ListAmbientesTrabalho(*ListRequest) (*ListResponse[AmbienteTrabalhoResponse], error)
}

type Handler struct {
	service ServiceInterface
	logger  *configs.Logger
}

func NewHandler(service ServiceInterface, logger *configs.Logger) *Handler {
	return &Handler{
		service: service,
		logger:  logger,
	}
}

// Horario Handlers
func (h *Handler) CreateHorario(w http.ResponseWriter, r *http.Request) {
	var req CreateHorarioRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.respondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	response, err := h.service.CreateHorario(&req)
	if err != nil {
		h.handleServiceError(w, err)
		return
	}

	h.respondWithJSON(w, http.StatusCreated, response)
}

func (h *Handler) GetHorario(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	response, err := h.service.GetHorarioByID(id)
	if err != nil {
		h.handleServiceError(w, err)
		return
	}

	h.respondWithJSON(w, http.StatusOK, response)
}

func (h *Handler) UpdateHorario(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var req UpdateHorarioRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.respondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	response, err := h.service.UpdateHorario(id, &req)
	if err != nil {
		h.handleServiceError(w, err)
		return
	}

	h.respondWithJSON(w, http.StatusOK, response)
}

func (h *Handler) DeleteHorario(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := h.service.DeleteHorario(id)
	if err != nil {
		h.handleServiceError(w, err)
		return
	}

	h.respondWithJSON(w, http.StatusNoContent, nil)
}

func (h *Handler) ListHorarios(w http.ResponseWriter, r *http.Request) {
	req := h.parseListRequest(r)

	response, err := h.service.ListHorarios(req)
	if err != nil {
		h.handleServiceError(w, err)
		return
	}

	h.respondWithJSON(w, http.StatusOK, response)
}

// HorarioDetalhe Handlers
func (h *Handler) CreateHorarioDetalhe(w http.ResponseWriter, r *http.Request) {
	var req CreateHorarioDetalheRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.respondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	response, err := h.service.CreateHorarioDetalhe(&req)
	if err != nil {
		h.handleServiceError(w, err)
		return
	}

	h.respondWithJSON(w, http.StatusCreated, response)
}

func (h *Handler) UpdateHorarioDetalhe(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var req UpdateHorarioDetalheRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.respondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	response, err := h.service.UpdateHorarioDetalhe(id, &req)
	if err != nil {
		h.handleServiceError(w, err)
		return
	}

	h.respondWithJSON(w, http.StatusOK, response)
}

func (h *Handler) DeleteHorarioDetalhe(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := h.service.DeleteHorarioDetalhe(id)
	if err != nil {
		h.handleServiceError(w, err)
		return
	}

	h.respondWithJSON(w, http.StatusNoContent, nil)
}

// AmbienteTrabalho Handlers
func (h *Handler) CreateAmbienteTrabalho(w http.ResponseWriter, r *http.Request) {
	var req CreateAmbienteTrabalhoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.respondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	response, err := h.service.CreateAmbienteTrabalho(&req)
	if err != nil {
		h.handleServiceError(w, err)
		return
	}

	h.respondWithJSON(w, http.StatusCreated, response)
}

func (h *Handler) GetAmbienteTrabalho(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	response, err := h.service.GetAmbienteTrabalhoByID(id)
	if err != nil {
		h.handleServiceError(w, err)
		return
	}

	h.respondWithJSON(w, http.StatusOK, response)
}

func (h *Handler) UpdateAmbienteTrabalho(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var req UpdateAmbienteTrabalhoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.respondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	response, err := h.service.UpdateAmbienteTrabalho(id, &req)
	if err != nil {
		h.handleServiceError(w, err)
		return
	}

	h.respondWithJSON(w, http.StatusOK, response)
}

func (h *Handler) DeleteAmbienteTrabalho(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := h.service.DeleteAmbienteTrabalho(id)
	if err != nil {
		h.handleServiceError(w, err)
		return
	}

	h.respondWithJSON(w, http.StatusNoContent, nil)
}

func (h *Handler) ListAmbientesTrabalho(w http.ResponseWriter, r *http.Request) {
	req := h.parseListRequest(r)

	response, err := h.service.ListAmbientesTrabalho(req)
	if err != nil {
		h.handleServiceError(w, err)
		return
	}

	h.respondWithJSON(w, http.StatusOK, response)
}

// Helper methods
func (h *Handler) parseListRequest(r *http.Request) *ListRequest {
	req := &ListRequest{}

	if page := r.URL.Query().Get("page"); page != "" {
		if p, err := strconv.Atoi(page); err == nil && p > 0 {
			req.Page = p
		}
	}

	if limit := r.URL.Query().Get("limit"); limit != "" {
		if l, err := strconv.Atoi(limit); err == nil && l > 0 && l <= 100 {
			req.Limit = l
		}
	}

	if empresaID := r.URL.Query().Get("empresa_id"); empresaID != "" {
		req.EmpresaID = &empresaID
	}

	if horarioID := r.URL.Query().Get("horario_id"); horarioID != "" {
		req.HorarioID = &horarioID
	}

	if search := r.URL.Query().Get("search"); search != "" {
		req.Search = &search
	}

	if ativo := r.URL.Query().Get("ativo"); ativo != "" {
		if a, err := strconv.ParseBool(ativo); err == nil {
			req.Ativo = &a
		}
	}

	return req
}

func (h *Handler) handleServiceError(w http.ResponseWriter, err error) {
	switch err {
	case ErrHorarioNotFound, ErrHorarioDetalheNotFound, ErrAmbienteTrabalhoNotFound:
		h.respondWithError(w, http.StatusNotFound, err.Error())
	case ErrHorarioDetalheAlreadyExists:
		h.respondWithError(w, http.StatusConflict, err.Error())
	default:
		h.logger.Error("Internal server error", "error", err)
		h.respondWithError(w, http.StatusInternalServerError, "Internal server error")
	}
}

func (h *Handler) respondWithJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if payload != nil {
		json.NewEncoder(w).Encode(payload)
	}
}

func (h *Handler) respondWithError(w http.ResponseWriter, status int, message string) {
	h.respondWithJSON(w, status, map[string]string{"error": message})
}
