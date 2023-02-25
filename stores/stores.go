package stores

import (
	"database/sql"
)

type Stores struct {
	DB     *sql.DB
	Author AuthorStore
	Book   BookStore
}

func New(db *sql.DB) *Stores {
	return &Stores{
		DB:     db,
		Author: &authorStore{db},
		Book:   &bookStore{db},
	}
}

func (s *Stores) Begin() (*sql.Tx, error) {
	return s.DB.Begin()
}

func (s *Stores) Commit(tx *sql.Tx) error {
	return tx.Commit()
}

func (s *Stores) RollBack(tx *sql.Tx) error {
	return tx.Rollback()
}
