package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose/v3"
	"io/ioutil"
	"os"
)

var GO_ENV = os.Getenv("GO_ENV")

func Load() (*sqlx.DB, error) {
	var uri string
	if GO_ENV == "development" {
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

	if GO_ENV == "development" {
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
