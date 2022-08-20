package db

import (
	"database/sql"
	"embed"
	_ "embed"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"io/ioutil"
	"log"
)

//go:embed migrations/*.sql
var migrations embed.FS

func New(development bool) (*sql.DB, error) {
	var uri string
	if development {
		uri = "postgres://postgres:postgres@echo-db:5432/postgres?sslmode=disable"
	} else {
		uri = "postgres://postgres:postgres@echo-db:5432/postgres?sslmode=disable"
	}

	db, err := sql.Open("postgres", uri)
	if err != nil {
		return nil, err
	}

	if err := goose.SetDialect("postgres"); err != nil {
		db.Close()
		return nil, err
	}
	goose.SetBaseFS(migrations)

	if err := goose.Up(db, "migrations"); err != nil {
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

			sqlCode := string(c)

			_, err = db.Exec(sqlCode)
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
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	if err != nil {
		log.Fatal(err)
	}

	return db, mock
}
