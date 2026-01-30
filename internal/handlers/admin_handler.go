package handlers

import (
	"LotteryImitationSystem/internal/services"
	"encoding/json"
	"net/http"
)

type AdminHandler struct {
	s *services.LotteryService
}

func NewAdminHandler(s *services.LotteryService) *AdminHandler {
	return &AdminHandler{s}
}

func (h *AdminHandler) Register(mux *http.ServeMux) {
	mux.HandleFunc("/admin/stats", h.stats)
}

func (h *AdminHandler) stats(w http.ResponseWriter, _ *http.Request) {
	if err := json.NewEncoder(w).Encode(h.s.Stats()); err != nil {
		http.Error(w, "failed to encode stats", http.StatusInternalServerError)
	}
}
