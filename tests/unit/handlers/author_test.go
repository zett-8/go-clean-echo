package handlers

import (
	"database/sql"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/zett-8/go-clean-echo/handlers"
	"github.com/zett-8/go-clean-echo/logger"
	"github.com/zett-8/go-clean-echo/models"
	"github.com/zett-8/go-clean-echo/services"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockAuthorService struct {
	services.AuthorService
	MockGetAuthors       func() ([]models.Author, error)
	MockDeleteAuthorById func(id int) error
}

func (m *MockAuthorService) GetAuthors() ([]models.Author, error) {
	return m.MockGetAuthors()
}

func (m *MockAuthorService) DeleteAuthor(id int) error {
	return m.MockDeleteAuthorById(id)
}

func setUp_author_test() func() {
	logger.New()

	return func() {
		logger.Sync()
		logger.Delete()
	}
}

func TestGetAuthorsSuccessCase(t *testing.T) {
	defer setUp_author_test()()

	s := &MockAuthorService{
		MockGetAuthors: func() ([]models.Author, error) {
			var r []models.Author
			return r, nil
		},
	}

	mockService := &services.Services{Author: s}

	e := handlers.Echo()
	h := handlers.New(mockService)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/author", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	assert.NoError(t, h.AuthorHandler.GetAuthors(c))
	assert.Equal(t, rec.Code, http.StatusOK)
}

func TestGetAuthors500Case(t *testing.T) {
	defer setUp_author_test()()

	s := &MockAuthorService{
		MockGetAuthors: func() ([]models.Author, error) {
			var r []models.Author
			return r, errors.New("fake error")
		},
	}

	mockService := &services.Services{Author: s}

	e := handlers.Echo()
	h := handlers.New(mockService)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/author", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	assert.NoError(t, h.AuthorHandler.GetAuthors(c))
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
}

func TestDeleteAuthorSuccessCase(t *testing.T) {
	defer setUp_author_test()()

	s := &MockAuthorService{
		MockDeleteAuthorById: func(id int) error {
			return nil
		},
	}

	mockService := &services.Services{Author: s}

	e := handlers.Echo()
	h := handlers.New(mockService)

	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("api/v1/author/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	assert.NoError(t, h.AuthorHandler.DeleteAuthorById(c))
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestDeleteAuthor400Case(t *testing.T) {
	defer setUp_author_test()()

	s := &MockAuthorService{
		MockDeleteAuthorById: func(id int) error {
			return nil
		},
	}

	mockService := &services.Services{Author: s}

	e := handlers.Echo()
	h := handlers.New(mockService)

	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("api/v1/author/:id")

	assert.NoError(t, h.AuthorHandler.DeleteAuthorById(c))
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestDeleteAuthor404Case(t *testing.T) {
	defer setUp_author_test()()

	s := &MockAuthorService{
		MockDeleteAuthorById: func(id int) error {
			return sql.ErrNoRows
		},
	}

	mockService := &services.Services{Author: s}

	e := handlers.Echo()
	h := handlers.New(mockService)

	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("api/v1/author/:id")
	c.SetParamNames("id")
	c.SetParamValues("4242872")

	assert.NoError(t, h.AuthorHandler.DeleteAuthorById(c))
	assert.Equal(t, http.StatusNotFound, rec.Code)
}

func TestDeleteAuthor500Case(t *testing.T) {
	defer setUp_author_test()()

	s := &MockAuthorService{
		MockDeleteAuthorById: func(id int) error {
			return errors.New("fake error")
		},
	}

	mockService := &services.Services{Author: s}

	e := handlers.Echo()
	h := handlers.New(mockService)

	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("api/v1/author/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	assert.NoError(t, h.AuthorHandler.DeleteAuthorById(c))
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
}
