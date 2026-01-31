

package handlers

import (
  "LotteryImitationSystem/internal/services"
  "encoding/json"
  "net/http"
)

type TicketHandler struct {
  s *services.LotteryService
}

func NewTicketHandler(s *services.LotteryService) *TicketHandler {
  return &TicketHandler{s}
}

func (h *TicketHandler) Register(mux *http.ServeMux) {
  mux.HandleFunc("/tickets", h.create)
}

func (h *TicketHandler) create(w http.ResponseWriter, r *http.Request) {
  var body struct{ UserID string }
  json.NewDecoder(r.Body).Decode(&body)
  json.NewEncoder(w).Encode(h.s.CreateTicket(body.UserID))
}
