package stores

import "github.com/jmoiron/sqlx"

type Stores struct {
	AuthorStore
	BookStore
}

func New(db *sqlx.DB) *Stores {
	return &Stores{
		AuthorStore: AuthorStore{db},
		BookStore:   BookStore{db},
	}
}
