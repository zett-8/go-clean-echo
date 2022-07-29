package stores

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/zett-8/go-clean-echo/models"
)

type AuthorStore struct {
	*sqlx.DB
}

func (s *AuthorStore) Get() ([]*models.Author, error) {
	authors := make([]*models.Author, 0)

	err := s.Select(&authors, "SELECT id, name, country from authors")

	if err != nil {
		return nil, err
	}

	return authors, nil
}

func (s *AuthorStore) DeleteById(id int) error {
	query := `
		DELETE FROM authors
		WHERE authors.id = $1
		RETURNING authors.id;
`

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
