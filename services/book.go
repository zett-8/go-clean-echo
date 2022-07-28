package services

import (
	"github.com/labstack/echo/v4"
	"github.com/zett-8/go-echo-without-orm/store"
	"net/http"
)

type BookService struct {
	store *store.BookStore
}

func NewBookService(s *store.BookStore) *BookService {
	return &BookService{
		store: s,
	}
}

func (s *BookService) GetBooks(c echo.Context) error {
	r, err := s.store.Get()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]error{"message": err})
	}

	return c.JSON(http.StatusOK, r)
}
