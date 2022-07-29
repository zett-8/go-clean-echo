package services

import (
	"github.com/zett-8/go-clean-echo/models"
	"github.com/zett-8/go-clean-echo/stores"
)

type AuthorService struct {
	store *stores.AuthorStore
}

func (s *AuthorService) GetAuthors() ([]models.Author, error) {
	r, err := s.store.Get()
	return r, err
}

func (s *AuthorService) DeleteAuthor(id int) error {
	err := s.store.DeleteById(id)
	return err
}
