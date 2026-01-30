package models

type PrizeType string

const (
	Money  PrizeType = "money"
	Travel PrizeType = "travel"
	Gift   PrizeType = "gift"
)

type Prize struct {
	TicketID string    `json:"ticket_id"`
	Type     PrizeType `json:"type"`
	Name     string    `json:"name"`
	Value    int       `json:"value"`
}
