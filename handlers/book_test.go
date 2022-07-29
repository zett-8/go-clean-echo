package handlers

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/zett-8/go-clean-echo/db"
	"github.com/zett-8/go-clean-echo/models"
	"github.com/zett-8/go-clean-echo/services"
	"github.com/zett-8/go-clean-echo/stores"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetBooksSuccessCase(t *testing.T) {
	mockDB, sqlmock := db.Mock()
	defer mockDB.Close()

	books := []models.Book{
		{ID: 1, Name: "test1", AuthorID: 1},
		{ID: 2, Name: "test2", AuthorID: 1},
	}

	rows := sqlmock.NewRows([]string{"id", "name", "author_id"})
	for _, b := range books {
		rows.AddRow(b.ID, b.Name, b.AuthorID)
	}
	sqlmock.MatchExpectationsInOrder(false)
	sqlmock.ExpectBegin()
	sqlmock.
		ExpectQuery("SELECT id, name, author_id from books;").
		WillReturnRows(rows)
	sqlmock.ExpectCommit()

	e := Echo()
	s := stores.New(mockDB)
	ss := services.New(s)
	h := New(ss)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/book", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	assert.NoError(t, h.BookHandler.GetBooks(c))
	assert.Equal(t, rec.Code, http.StatusOK)

	_expected, _ := json.Marshal(books)
	expected := string(_expected)
	got := strings.TrimSuffix(rec.Body.String(), "\n")

	assert.Equal(t, expected, got)
}
