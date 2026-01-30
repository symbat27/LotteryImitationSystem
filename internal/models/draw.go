package models

type Draw struct {
	ID             string `json:"id"`
	WinningNumbers []int  `json:"winning_numbers"`
	Status         string `json:"status"`
}
