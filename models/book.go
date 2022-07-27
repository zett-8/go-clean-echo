package models

type Book struct {
	id       int    `json:"id"`
	name     string `json:"name"`
	authorID int    `json:"author_id"`
}
