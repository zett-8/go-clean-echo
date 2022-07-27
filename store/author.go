package store

import (
	"database/sql"
	"fmt"
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

func (s *AuthorStore) Get() (string, error) {
	var author *models.Author

	rows, err := s.db.Query("SELECT id, name, country from authors")
	if err != nil {
		return "", err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&author)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(author)
	}

	return "everything's done", nil
}
