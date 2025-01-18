package models

type Customer struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	PointBalance int    `json:"point_balance"`
}
