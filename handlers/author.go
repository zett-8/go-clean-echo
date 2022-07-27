package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/zett-8/go-echo-without-orm/services"
)

func NewAuthorHandler(e *echo.Echo, s *services.AuthorService) {
	e.GET("/author", s.GetAuthors)
}
