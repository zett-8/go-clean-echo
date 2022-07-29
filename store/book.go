package store

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/zett-8/go-clean-echo/models"
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
		return nil, err
	}

	return books, nil
}

func (s *BookStore) DeleteById(id int) error {
	query := `
		DELETE FROM books
		WHERE books.id = $1
		RETURNING books.id;
`

	row, err := s.db.Exec(query, id)
	if err != nil {
		return err
	}

	if r, err := row.RowsAffected(); err != nil {
		return err
	} else if r == 0 {
		return sql.ErrNoRows
	}

	return nil
}
