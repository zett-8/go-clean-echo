package db

import (
	"database/sql"
	"github.com/pressly/goose/v3"
	"os"
)

var GO_ENV = os.Getenv("GO_ENV")

func Load() (*sql.DB, error) {
	var uri string
	if GO_ENV == "development" {
		uri = "postgres://postgres:postgres@echo-db:5432/postgres?sslmode=disable"
	} else {
		uri = "postgres://postgres:postgres@echo-db:5432/postgres?sslmode=disable"
	}

	db, err := sql.Open("pgx", uri)
	if err != nil {
		return nil, err
	}

	if err := goose.SetDialect("postgres"); err != nil {
		db.Close()
		return nil, err
	}

	if err := goose.Up(db, "db/migrations"); err != nil {
		db.Close()
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
