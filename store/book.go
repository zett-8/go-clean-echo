package store

import (
	"database/sql"
	"github.com/zett-8/go-echo-without-orm/models"
	"log"
)

type BookStore struct {
	db *sql.DB
}

func NewBooksStore(db *sql.DB) *BookStore {
	return &BookStore{
		db: db,
	}
}

func (s *BookStore) Get() ([]*models.Book, error) {
	rows, err := s.db.Query("SELECT id, name, author_id from books")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer rows.Close()

	var books []*models.Book
	for rows.Next() {
		book := &models.Book{}

		err := rows.Scan(&book.ID, &book.Name, &book.AuthorID)
		if err != nil {
			log.Fatal(err)
		}

		books = append(books, book)
	}

	return books, nil
}
