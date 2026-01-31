package storage

import (
  "LotteryImitationSystem/internal/models"
  "sync"
)

type UserRepository struct {
  mu sync.RWMutex
  db map[string]models.User
}

func NewUserRepository() *UserRepository {
  return &UserRepository{db: map[string]models.User{}}
}

func (r *UserRepository) Save(u models.User) {
  r.mu.Lock()
  defer r.mu.Unlock()
  r.db[u.ID] = u
}
