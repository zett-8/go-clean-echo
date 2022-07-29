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

func TestGetAuthorsSuccessCase(t *testing.T) {
	mockDB, sqlmock := db.Mock()
	defer mockDB.Close()

	authors := []models.Author{
		{ID: 1, Name: "test1", Country: "US"},
		{ID: 2, Name: "test2", Country: "UK"},
	}

	rows := sqlmock.NewRows([]string{"id", "name", "country"})
	for _, a := range authors {
		rows.AddRow(a.ID, a.Name, a.Country)
	}
	sqlmock.MatchExpectationsInOrder(false)
	sqlmock.ExpectBegin()
	sqlmock.
		ExpectQuery("SELECT id, name, country from authors").
		WillReturnRows(rows)
	sqlmock.ExpectCommit()

	e := Echo()
	s := stores.New(mockDB)
	ss := services.New(s)
	h := New(ss)

	req := httptest.NewRequest(http.MethodGet, "/api/v1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	assert.NoError(t, h.AuthorHandler.GetAuthors(c))
	assert.Equal(t, rec.Code, http.StatusOK)

	_expected, _ := json.Marshal(authors)
	expected := string(_expected)
	got := strings.TrimSuffix(rec.Body.String(), "\n")

	assert.Equal(t, expected, got)
}
