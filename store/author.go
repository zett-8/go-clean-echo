package store

import (
	"github.com/jmoiron/sqlx"
	"github.com/zett-8/go-clean-echo/models"
)

type AuthorStore struct {
	db *sqlx.DB
}

func NewAuthorStore(db *sqlx.DB) *AuthorStore {
	return &AuthorStore{
		db: db,
	}
}

func (s *AuthorStore) Get() ([]*models.Author, error) {
	authors := make([]*models.Author, 0)

	err := s.db.Select(&authors, "SELECT id, name, country from authors")

	if err != nil {
		return nil, err
	}

	return authors, nil
}
