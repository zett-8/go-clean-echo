package services

import (
	"github.com/zett-8/go-clean-echo/models"
	"github.com/zett-8/go-clean-echo/stores"
)

type AuthorService interface {
	GetAuthors() ([]models.Author, error)
	DeleteAuthor(id int) error
}

type AuthorServiceContext struct {
	store stores.AuthorStore
}

func (s *AuthorServiceContext) GetAuthors() ([]models.Author, error) {
	r, err := s.store.Get()
	return r, err
}

func (s *AuthorServiceContext) DeleteAuthor(id int) error {
	err := s.store.DeleteById(id)
	return err
}
