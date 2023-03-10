package stores

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/zett-8/go-clean-echo/db"
	"github.com/zett-8/go-clean-echo/models"
	"github.com/zett-8/go-clean-echo/stores"
	"testing"
)

func TestBookStore_GetSuccessCase(t *testing.T) {
	mockDB, mock := db.Mock()
	defer mockDB.Close()

	books := []models.Book{
		{ID: 1, Name: "test1", AuthorID: 1},
		{ID: 2, Name: "test2", AuthorID: 1},
	}

	rows := mock.NewRows([]string{"id", "name", "author_id"})
	for _, b := range books {
		rows.AddRow(b.ID, b.Name, b.AuthorID)
	}

	mock.
		ExpectQuery("SELECT id, name, author_id from books").
		WillReturnRows(rows)

	s := stores.New(mockDB)

	r, err := s.Book.Get(nil)

	assert.NoError(t, err)
	assert.Equal(t, books, r)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestBookStore_DeleteByIdSuccessCase(t *testing.T) {
	mockDB, mock := db.Mock()
	defer mockDB.Close()

	books := []models.Book{
		{ID: 1, Name: "test1", AuthorID: 1},
		{ID: 2, Name: "test2", AuthorID: 1},
	}

	deletingID := books[0].ID

	rows := mock.NewRows([]string{"id", "name", "author_id"})
	for _, b := range books {
		rows.AddRow(b.ID, b.Name, b.AuthorID)
	}
	mock.
		ExpectExec("DELETE FROM books WHERE books.id = $1 RETURNING books.id").
		WithArgs(deletingID).
		WillReturnResult(sqlmock.NewResult(int64(deletingID), 1))
	mock.
		ExpectExec("DELETE FROM books WHERE books.id = $1 RETURNING books.id").
		WithArgs(deletingID).
		WillReturnResult(sqlmock.NewResult(int64(deletingID), 0))

	s := stores.New(mockDB)

	assert.NoError(t, s.Book.DeleteById(nil, deletingID))
	assert.Equal(t, s.Book.DeleteById(nil, deletingID), sql.ErrNoRows)
	assert.NoError(t, mock.ExpectationsWereMet())
}
