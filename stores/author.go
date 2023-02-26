package stores

import (
	"database/sql"
	"github.com/zett-8/go-clean-echo/logger"
	"github.com/zett-8/go-clean-echo/models"
	"go.uber.org/zap"
)

type (
	AuthorStore interface {
		Get(tx *sql.Tx) ([]models.Author, error)
		Create(tx *sql.Tx, author *models.Author) (int64, error)
		UpdateById(tx *sql.Tx, author *models.Author) (int64, error)
		DeleteById(tx *sql.Tx, id int) error
	}

	authorStore struct {
		*sql.DB
	}
)

func (s *authorStore) Get(tx *sql.Tx) ([]models.Author, error) {
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

func (s *authorStore) Create(tx *sql.Tx, author *models.Author) (int64, error) {
	var err error

	query := "INSERT INTO authors (name, country) VALUES ($1, $2) RETURNING id"
	if err != nil {
		return 0, err
	}

	var id int64

	if tx != nil {
		err = tx.QueryRow(query, author.Name, author.Country).Scan(&id)
	} else {
		err = s.QueryRow(query, author.Name, author.Country).Scan(&id)
	}

	if err != nil {
		logger.Error("failed to create author", zap.Error(err))
		return 0, err
	}

	return id, nil
}

func (s *authorStore) UpdateById(tx *sql.Tx, author *models.Author) (int64, error) {
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

func (s *authorStore) DeleteById(tx *sql.Tx, id int) error {
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
