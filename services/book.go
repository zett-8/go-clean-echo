package services

import (
	"github.com/zett-8/go-clean-echo/models"
	"github.com/zett-8/go-clean-echo/stores"
)

type BookService interface {
	GetBooks() ([]models.Book, error)
	DeleteBookById(id int) error
}

type BookServiceContext struct {
	stores *stores.Stores
}

func (s *BookServiceContext) GetBooks() ([]models.Book, error) {
	r, err := s.stores.BookStore.Get(nil)
	return r, err
}

func (s *BookServiceContext) DeleteBookById(id int) error {
	err := s.stores.BookStore.DeleteById(nil, id)
	return err
}
