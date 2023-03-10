package db

import (
	"database/sql"
	"embed"
	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"
	"github.com/zett-8/go-clean-echo/logger"
	"go.uber.org/zap"
	"io/ioutil"
)

//go:embed migrations/*.sql
var migrationsFS embed.FS

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

	migrations := &migrate.EmbedFileSystemMigrationSource{
		FileSystem: migrationsFS,
		Root:       "migrations",
	}

	if _, err := migrate.Exec(db, "postgres", migrations, migrate.Up); err != nil {
		return nil, err
	}

	// Run all seed sql file if on dev
	if development {
		seedFiles, err := ioutil.ReadDir("db/seed/")
		if err != nil {
			logger.Error("failed to read seed files", zap.Error(err))
		}

		for _, f := range seedFiles {
			c, err := ioutil.ReadFile("db/seed/" + f.Name())
			if err != nil {
				logger.Error("failed to read seed file", zap.Error(err))
			}

			sqlCode := string(c)

			_, err = db.Exec(sqlCode)
			if err != nil {
				logger.Error("failed to seed database", zap.Error(err))
			}
		}

		logger.Info("seeded database")
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
		logger.Fatal("failed to create mock db", zap.Error(err))
	}

	return db, mock
}
