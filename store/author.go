package store

import (
	"database/sql"
	"github.com/zett-8/go-echo-without-orm/models"
	"log"
)

type AuthorStore struct {
	db *sql.DB
}

func NewAuthorStore(db *sql.DB) *AuthorStore {
	return &AuthorStore{
		db: db,
	}
}

func (s *AuthorStore) Get() ([]*models.Author, error) {
	var authors []*models.Author

	rows, err := s.db.Query("SELECT id, name, country from authors")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		author := &models.Author{}

		err := rows.Scan(&author.ID, &author.Name, &author.Country)
		if err != nil {
			log.Fatal(err)
		}

		authors = append(authors, author)
	}

	return authors, nil
}
