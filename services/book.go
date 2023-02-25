package services

import (
	"github.com/zett-8/go-clean-echo/models"
	"github.com/zett-8/go-clean-echo/stores"
)

type (
	BookService interface {
		GetBooks() ([]models.Book, error)
		DeleteBookById(id int) error
	}

	bookService struct {
		stores *stores.Stores
	}
)

func (s *bookService) GetBooks() ([]models.Book, error) {
	r, err := s.stores.Book.Get(nil)
	return r, err
}

func (s *bookService) DeleteBookById(id int) error {
	err := s.stores.Book.DeleteById(nil, id)
	return err
}
