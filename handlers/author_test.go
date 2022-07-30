package handlers

import (
	"github.com/zett-8/go-clean-echo/models"
	"github.com/zett-8/go-clean-echo/services"
	"testing"
)

type MockAuthorService struct {
	services.AuthorService
	MockGetAuthors   func() ([]models.Author, error)
	MockDeleteAuthor func() error
}

func (m *MockAuthorService) GetAuthors() ([]models.Author, error) {
	return m.MockGetAuthors()
}

func (m *MockAuthorService) DeleteAuthor() error {
	return m.MockDeleteAuthor()
}

func TestGetAuthorsSuccessCase(t *testing.T) {

	//mockService := &services.Services{AuthorService: MockAuthorService}
	//e := Echo()
	//h := New(mockService)
	//
	//req := httptest.NewRequest(http.MethodGet, "/api/v1/author", nil)
	//rec := httptest.NewRecorder()
	//c := e.NewContext(req, rec)
	//
	//assert.NoError(t, h.AuthorHandler.GetAuthors(c))
	//assert.Equal(t, rec.Code, http.StatusOK)
}

func TestDeleteAuthorSuccessCase(t *testing.T) {
	//mockDB, mock := db.Mock()
	//defer mockDB.Close()
	//
	//authors := []models.Author{
	//	{ID: 1, Name: "test1", Country: "US"},
	//	{ID: 2, Name: "test2", Country: "UK"},
	//}
	//
	//deletingID := authors[0].ID
	//
	//rows := mock.NewRows([]string{"id", "name", "country"})
	//for _, a := range authors {
	//	rows.AddRow(a.ID, a.Name, a.Country)
	//}
	//mock.MatchExpectationsInOrder(false)
	//mock.ExpectBegin()
	//mock.ExpectExec(`
	//	DELETE FROM authors
	//	WHERE authors.id = $1
	//	RETURNING authors.id;
	//`).
	//	WithArgs(deletingID).
	//	WillReturnResult(sqlmock.NewResult(1, 1))
	//mock.ExpectCommit()
	//
	//e := Echo()
	//s := stores.New(mockDB)
	//ss := services.New(s)
	//h := New(ss)
	//
	//req := httptest.NewRequest(http.MethodDelete, "/", nil)
	//rec := httptest.NewRecorder()
	//c := e.NewContext(req, rec)
	//c.SetPath("/api/v1/author/:id")
	//c.SetParamNames("id")
	//c.SetParamValues(fmt.Sprint(deletingID))
	//
	//assert.NoError(t, h.AuthorHandler.DeleteAuthor(c))
	//assert.Equal(t, rec.Code, http.StatusOK)
	//
	//_expected, _ := json.Marshal(authors)
	//expected := string(_expected)
	//got := strings.TrimSuffix(rec.Body.String(), "\n")
	//
	//assert.Equal(t, expected, got)
}
