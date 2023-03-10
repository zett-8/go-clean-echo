package services

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/zett-8/go-clean-echo/db"
	"github.com/zett-8/go-clean-echo/models"
	services2 "github.com/zett-8/go-clean-echo/services"
	"github.com/zett-8/go-clean-echo/stores"
	"testing"
)

func TestAuthorServiceContext_CreateAuthorWithBooks_Success(t *testing.T) {
	mockDB, mock := db.Mock()
	defer mockDB.Close()

	s := stores.New(mockDB)
	services := services2.New(s)

	a := &models.Author{
		Name:    "test",
		Country: "US",
	}

	mock.ExpectBegin()
	mock.
		ExpectQuery("INSERT INTO authors (name, country) VALUES ($1, $2) RETURNING id").
		WithArgs(a.Name, a.Country).
		WillReturnRows(
			sqlmock.NewRows([]string{"id"}).AddRow(1),
		)
	mock.ExpectCommit()

	r, err := services.Author.CreateAuthorWithBooks(a, nil)

	assert.NoError(t, err)
	assert.Equal(t, int64(1), r)
}
