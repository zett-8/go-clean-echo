package stores

import (
	"database/sql"
	"github.com/zett-8/go-clean-echo/models"
)

type (
	BookStore interface {
		Get(tx *sql.Tx) ([]models.Book, error)
		DeleteById(tx *sql.Tx, id int) error
	}

	bookStore struct {
		*sql.DB
	}
)

func (s *bookStore) Get(tx *sql.Tx) ([]models.Book, error) {
	books := make([]models.Book, 0)

	query := "SELECT id, name, author_id from books"

	rows, err := s.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var b models.Book
		err = rows.Scan(&b.ID, &b.Name, &b.AuthorID)
		books = append(books, b)
	}

	return books, nil
}

func (s *bookStore) DeleteById(tx *sql.Tx, id int) error {
	query := "DELETE FROM books WHERE books.id = $1 RETURNING books.id"

	row, err := s.Exec(query, id)
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
