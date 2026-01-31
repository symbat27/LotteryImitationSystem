
package models

import "time"

type Ticket struct {
  ID        string    json:"id"
  UserID    string    json:"user_id"
  Numbers   []int     json:"numbers"
  CreatedAt time.Time json:"created_at"
}
