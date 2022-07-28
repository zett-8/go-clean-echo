package models

type Book struct {
	ID       int    `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
	AuthorID int    `json:"author_id" db:"author_id"`
}
