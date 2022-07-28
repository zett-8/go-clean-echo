package store

import (
	"github.com/jmoiron/sqlx"
	"github.com/zett-8/go-clean-echo/models"
	"log"
)

type BookStore struct {
	db *sqlx.DB
}

func NewBooksStore(db *sqlx.DB) *BookStore {
	return &BookStore{
		db: db,
	}
}

func (s *BookStore) Get() ([]*models.Book, error) {
	books := make([]*models.Book, 0)

	query := "SELECT id, name, author_id from books;"

	err := s.db.Select(&books, query)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return books, nil
}
