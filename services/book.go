package services

import (
	"github.com/zett-8/go-clean-echo/models"
	"github.com/zett-8/go-clean-echo/stores"
)

type BookService struct {
	store *stores.BookStore
}

func (s *BookService) GetBooks() ([]*models.Book, error) {
	r, err := s.store.Get()
	return r, err
}

func (s *BookService) DeleteBookById(id int) error {
	err := s.store.DeleteById(id)
	return err
}
