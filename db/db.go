package db

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose/v3"
	"io/ioutil"
	"log"
)

func New(development bool) (*sqlx.DB, error) {
	var uri string
	if development {
		uri = "postgres://postgres:postgres@echo-db:5432/postgres?sslmode=disable"
	} else {
		uri = "postgres://postgres:postgres@echo-db:5432/postgres?sslmode=disable"
	}

	db, err := sqlx.Open("pgx", uri)
	if err != nil {
		return nil, err
	}

	if err := goose.SetDialect("postgres"); err != nil {
		db.Close()
		return nil, err
	}

	if err := goose.Up(db.DB, "db/migrations"); err != nil {
		db.Close()
		return nil, err
	}

	// Run all seed sql file if on dev
	if development {
		seedFiles, err := ioutil.ReadDir("db/seed/")
		if err != nil {
			fmt.Println(err)
		}

		for _, f := range seedFiles {
			c, err := ioutil.ReadFile("db/seed/" + f.Name())
			if err != nil {
				fmt.Println(err)
			}

			sql := string(c)

			_, err = db.Exec(sql)
			if err != nil {
				fmt.Println(err)
			}
		}

		fmt.Println("seeded database")
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func Mock() (*sqlx.DB, sqlmock.Sqlmock) {
	_db, mock, err := sqlmock.New()

	if err != nil {
		log.Fatal(err)
	}

	db := sqlx.NewDb(_db, "pgx")

	_ = goose.SetDialect("postgres")
	_ = goose.Up(db.DB, "db/migrations")

	return db, mock
}
