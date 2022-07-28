package services

import (
	"github.com/labstack/echo/v4"
	"github.com/zett-8/go-clean-echo/store"
	"log"
	"net/http"
)

type AuthorService struct {
	store *store.AuthorStore
}

func NewAuthorService(s *store.AuthorStore) *AuthorService {
	return &AuthorService{
		store: s,
	}
}

func (s *AuthorService) GetAuthors(c echo.Context) error {
	r, err := s.store.Get()

	if err != nil {
		log.Print(err)
		return c.JSON(http.StatusInternalServerError, map[string]error{"message": err})
	}

	return c.JSON(http.StatusOK, r)
}
