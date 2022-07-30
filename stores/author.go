package stores

import (
	"database/sql"
	"github.com/zett-8/go-clean-echo/models"
)

type AuthorStore struct {
	*sql.DB
}

func (s *AuthorStore) Get() ([]models.Author, error) {
	authors := make([]models.Author, 0)

	rows, err := s.Query("SELECT id, name, country from authors")

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var a models.Author
		err = rows.Scan(&a.ID, &a.Name, &a.Country)
		authors = append(authors, a)
	}

	return authors, nil
}

func (s *AuthorStore) DeleteById(id int) error {
	row, err := s.Exec("DELETE FROM authors WHERE authors.id = $1 RETURNING authors.id", id)
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
