
package storage

import (
  "LotteryImitationSystem/internal/models"
  "sync"
)

type PrizeRepository struct {
  mu sync.RWMutex
  db []models.Prize
}

func NewPrizeRepository() *PrizeRepository {
  return &PrizeRepository{}
}

func (r *PrizeRepository) Save(p models.Prize) {
  r.mu.Lock()
  defer r.mu.Unlock()
  r.db = append(r.db, p)
}

func (r *PrizeRepository) List() []models.Prize {
  r.mu.RLock()
  defer r.mu.RUnlock()
  return r.db
}
