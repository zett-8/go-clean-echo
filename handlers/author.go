package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/zett-8/go-clean-echo/services"
)

//type AuthorHandler struct {
//	*services.AuthorService
//}

//func (h *AuthorHandler) GetAuthors(c echo.Context) error {
//
//	return c.JSON(http.StatusOK, "r")
//}

func NewAuthorHandler(e *echo.Group, s *services.AuthorService) {
	e.GET("/author", s.GetAuthors)
	e.DELETE("/author/:id", s.DeleteAuthor)
}
