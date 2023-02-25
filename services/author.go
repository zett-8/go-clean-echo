package services

import (
	"github.com/zett-8/go-clean-echo/models"
	"github.com/zett-8/go-clean-echo/stores"
)

type (
	AuthorService interface {
		GetAuthors() ([]models.Author, error)
		CreateAuthor(a *models.Author) (int64, error)
		CreateAuthorWithBooks(a *models.Author, bs *[]models.Book) (int64, error)
		UpdateAuthorById(a *models.Author) (int64, error)
		DeleteAuthor(id int) error
	}

	authorService struct {
		stores *stores.Stores
	}
)

func (s *authorService) GetAuthors() ([]models.Author, error) {
	r, err := s.stores.AuthorStore.Get(nil)
	return r, err
}

func (s *authorService) CreateAuthor(a *models.Author) (int64, error) {
	r, err := s.stores.AuthorStore.Create(nil, a)
	return r, err
}

func (s *authorService) CreateAuthorWithBooks(a *models.Author, bs *[]models.Book) (int64, error) {
	tx, err := s.stores.DB.Begin()
	if err != nil {
		return 0, err
	}

	id, err := s.stores.AuthorStore.Create(tx, a)
	if err != nil {
		s.stores.RollBack(tx)
		return 0, err
	}

	// run other SQL

	err = s.stores.Commit(tx)
	if err != nil {
		s.stores.RollBack(tx)
		return 0, err
	}

	return id, nil
}

func (s *authorService) UpdateAuthorById(a *models.Author) (int64, error) {
	r, err := s.stores.AuthorStore.UpdateById(nil, a)
	return r, err
}

func (s *authorService) DeleteAuthor(id int) error {
	err := s.stores.AuthorStore.DeleteById(nil, id)
	return err
}
