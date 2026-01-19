package models

type Ticket struct {
	ID              int
	UserID          int
	DrawID          int
	SelectedNumbers []int
}
