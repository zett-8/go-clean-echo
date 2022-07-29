package services

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"github.com/zett-8/go-clean-echo/stores"
	"log"
	"net/http"
	"strconv"
)

type BookService struct {
	store *stores.BookStore
}

// GetBooks
// @Summary Fetch a list of all books.
// @Description Fetch a list of all books.
// @Tags Book
// @Accept */*
// @Produce json
// @Success 200 {object} []models.Book
// @Failure 500 {object} utils.Error
// @Router /api/v1/book [get]
func (s *BookService) GetBooks(c echo.Context) error {
	r, err := s.store.Get()

	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]error{"message": err})
	}

	return c.JSON(http.StatusOK, r)
}

// DeleteBook
// @Summary Delete a book by ID.
// @Description Delete a book by ID.
// @Tags Book
// @Accept */*
// @Produce json
// @Param id path int true "Book id"
// @Success 200 {integer} int "Deleted Book ID"
// @Failure 500 {object} utils.Error
// @Router /api/v1/book/{id} [delete]
func (s *BookService) DeleteBook(c echo.Context) error {
	_id := c.Param("id")
	id, _ := strconv.Atoi(_id)

	err := s.store.DeleteById(id)
	if err == sql.ErrNoRows {
		return c.JSON(http.StatusNotFound, "not found")
	} else if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]error{"message": err})
	}

	return c.JSON(http.StatusOK, "OK")
}
