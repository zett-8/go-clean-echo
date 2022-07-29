package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/zett-8/go-clean-echo/services"
)

func NewAuthorHandler(e *echo.Group, s *services.AuthorService) {
	e.GET("/author", s.GetAuthors)
	e.DELETE("/author/:id", s.DeleteAuthor)
}
