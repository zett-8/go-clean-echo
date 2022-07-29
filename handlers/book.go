package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/zett-8/go-clean-echo/services"
)

func NewBookHandler(e *echo.Group, s *services.BookService) {
	e.GET("/book", s.GetBooks)
	e.DELETE("/book/:id", s.DeleteBook)
}
