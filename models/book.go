package models

type Book struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	AuthorID int    `json:"author_id"`
}
