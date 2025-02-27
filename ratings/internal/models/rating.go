package models

type Rating struct {
	ID       int64   `json:"rating_id" db:"rating_id"`
	FromUser string  `json:"from_user" db:"from_user"`
	ToUser   string  `json:"to_user" db:"to_user"`
	Rating   float64 `json:"rating" db:"rating"`
} // @name rating
