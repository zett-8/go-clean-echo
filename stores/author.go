package stores

import (
	"database/sql"
	"github.com/zett-8/go-clean-echo/models"
	"log"
)

type AuthorStore interface {
	Get() ([]models.Author, error)
	Create(author *models.Author) (int64, error)
	UpdateById(author *models.Author) (int64, error)
	DeleteById(id int) error
}

type AuthorStoreContext struct {
	*sql.DB
}

func (s *AuthorStoreContext) Get() ([]models.Author, error) {
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

func (s *AuthorStoreContext) Create(author *models.Author) (int64, error) {
	query, err := s.Prepare("INSERT INTO authors (name, country) VALUES ($1, $2) RETURNING id")
	if err != nil {
		return 0, err
	}

	var id int64

	if err = query.QueryRow(author.Name, author.Country).Scan(&id); err != nil {
		log.Println(err)
		return 0, err
	}

	return id, nil
}

func (s *AuthorStoreContext) UpdateById(author *models.Author) (int64, error) {
	query, err := s.Prepare("UPDATE authors SET name = $1, country = $2 WHERE authors.id = $3 RETURNING id")
	if err != nil {
		return 0, err
	}

	var id int64

	err = query.QueryRow(author.Name, author.Country, author.ID).Scan(&id)
	if id == 0 {
		return 0, sql.ErrNoRows
	} else if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *AuthorStoreContext) DeleteById(id int) error {
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
