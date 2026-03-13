package auth

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {

	var req RegisterRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token, err := h.service.Register(
		r.Context(),
		req.Email,
		req.Password,
	)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(AuthResponse{
		Token: token,
	})
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {

	var req LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token, err := h.service.Login(
		r.Context(),
		req.Email,
		req.Password,
	)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(AuthResponse{
		Token: token,
	})
}
