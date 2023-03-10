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

type MockBookService struct {
	services.BookService
	MockGetBooks       func() ([]models.Book, error)
	MockDeleteBookById func(id int) error
}

func (m *MockBookService) GetBooks() ([]models.Book, error) {
	return m.MockGetBooks()
}

func (m *MockBookService) DeleteBookById(id int) error {
	return m.MockDeleteBookById(id)
}

func setUp_book_test() func() {
	logger.New()

	return func() {
		logger.Sync()
		logger.Delete()
	}
}

func TestGetBooksSuccessCase(t *testing.T) {
	defer setUp_book_test()()

	s := &MockBookService{
		MockGetBooks: func() ([]models.Book, error) {
			var r []models.Book
			return r, nil
		},
	}

	mockService := &services.Services{Book: s}

	e := handlers.Echo()
	h := handlers.New(mockService)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/book", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	assert.NoError(t, h.BookHandler.GetBooks(c))
	assert.Equal(t, rec.Code, http.StatusOK)
}

func TestGetBooks500Case(t *testing.T) {
	defer setUp_book_test()()

	s := &MockBookService{
		MockGetBooks: func() ([]models.Book, error) {
			var r []models.Book
			return r, errors.New("fake error")
		},
	}

	mockService := &services.Services{Book: s}

	e := handlers.Echo()
	h := handlers.New(mockService)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/book", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	assert.NoError(t, h.BookHandler.GetBooks(c))
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
}

func TestDeleteBookSuccessCase(t *testing.T) {
	defer setUp_book_test()()

	s := &MockBookService{
		MockDeleteBookById: func(id int) error {
			return nil
		},
	}

	mockService := &services.Services{Book: s}

	e := handlers.Echo()
	h := handlers.New(mockService)

	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("api/v1/book/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	assert.NoError(t, h.BookHandler.DeleteBook(c))
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestDeleteBook400Case(t *testing.T) {
	defer setUp_book_test()()

	s := &MockBookService{
		MockDeleteBookById: func(id int) error {
			return nil
		},
	}

	mockService := &services.Services{Book: s}

	e := handlers.Echo()
	h := handlers.New(mockService)

	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("api/v1/Book/:id")

	assert.NoError(t, h.BookHandler.DeleteBook(c))
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestDeleteBook404Case(t *testing.T) {
	defer setUp_book_test()()

	s := &MockBookService{
		MockDeleteBookById: func(id int) error {
			return sql.ErrNoRows
		},
	}

	mockService := &services.Services{Book: s}

	e := handlers.Echo()
	h := handlers.New(mockService)

	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("api/v1/Book/:id")
	c.SetParamNames("id")
	c.SetParamValues("4242872")

	assert.NoError(t, h.BookHandler.DeleteBook(c))
	assert.Equal(t, http.StatusNotFound, rec.Code)
}

func TestDeleteBook500Case(t *testing.T) {
	defer setUp_book_test()()

	s := &MockBookService{
		MockDeleteBookById: func(id int) error {
			return errors.New("fake error")
		},
	}

	mockService := &services.Services{Book: s}

	e := handlers.Echo()
	h := handlers.New(mockService)

	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("api/v1/book/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	assert.NoError(t, h.BookHandler.DeleteBook(c))
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
}
