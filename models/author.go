package models

type Author struct {
	ID      int    `json:"id" db:"id"`
	Name    string `json:"name" db:"name"`
	Country string `json:"country" db:"country"`
}
