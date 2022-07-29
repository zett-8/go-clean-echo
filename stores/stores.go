package stores

import (
	"database/sql"
)

type Stores struct {
	AuthorStore
	BookStore
}

func New(db *sql.DB) *Stores {
	return &Stores{
		AuthorStore: AuthorStore{db},
		BookStore:   BookStore{db},
	}
}
