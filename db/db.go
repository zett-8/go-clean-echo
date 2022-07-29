package db

import (
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/pressly/goose/v3"
	"io/ioutil"
	"log"
)

func New(development bool) (*sql.DB, error) {
	var uri string
	if development {
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

func Mock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	//db := sql.Open("sqlmock", sqlmock.dsn)
	if err != nil {
		log.Fatal(err)
	}

	_ = goose.SetDialect("sqlmock")
	_ = goose.Up(db, "db/migrations")

	return db, mock
}
