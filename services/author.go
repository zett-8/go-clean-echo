package services

import (
	"github.com/labstack/echo/v4"
	"github.com/zett-8/go-echo-without-orm/store"
	"net/http"
)

type AuthorService struct {
	*store.AuthorStore
}

func NewAuthorService(s *store.AuthorStore) *AuthorService {
	return &AuthorService{s}
}

func (s *AuthorService) GetAuthors(c echo.Context) error {
	r, err := s.Get()

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]error{"message": err})
	}

	return c.String(http.StatusOK, r)
}
