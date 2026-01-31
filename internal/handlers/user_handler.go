package handlers

import (
  "LotteryImitationSystem/internal/services"
  "encoding/json"
  "net/http"
)

type UserHandler struct {
  s *services.LotteryService
}

func NewUserHandler(s *services.LotteryService) *UserHandler {
  return &UserHandler{s}
}

func (h *UserHandler) Register(mux *http.ServeMux) {
  mux.HandleFunc("/users", h.create)
}

func (h *UserHandler) create(w http.ResponseWriter, r *http.Request) {
  var body struct {
    Username string json:"username"
  }

  if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
    http.Error(w, "invalid JSON", http.StatusBadRequest)
    return
  }

  user := h.s.CreateUser(body.Username)

  if err := json.NewEncoder(w).Encode(user); err != nil {
    http.Error(w, "failed to encode response", http.StatusInternalServerError)
  }
}
