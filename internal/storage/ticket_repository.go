package storage

import (
  "LotteryImitationSystem/internal/models"
  "sync"
)

type TicketRepository struct {
  mu sync.RWMutex
  db []models.Ticket
}

func NewTicketRepository() *TicketRepository {
  return &TicketRepository{}
}

func (r *TicketRepository) Save(t models.Ticket) {
  r.mu.Lock()
  defer r.mu.Unlock()
  r.db = append(r.db, t)
}
