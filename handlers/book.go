package handlers

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"github.com/zett-8/go-clean-echo/services"
	"log"
	"net/http"
	"strconv"
)

type BookHandler struct {
	*services.BookService
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
func (h *BookHandler) GetBooks(c echo.Context) error {
	r, err := h.BookService.GetBooks()

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
// @Failure 400 {string} string "ID is invalid"
// @Failure 404 {string} string "Not found"
// @Failure 500 {object} utils.Error
// @Router /api/v1/book/{id} [delete]
func (h *BookHandler) DeleteBook(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, "ID is Invalid")
	}

	err = h.BookService.DeleteBookById(id)

	if err == sql.ErrNoRows {
		return c.JSON(http.StatusNotFound, "not found")
	} else if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]error{"message": err})
	}

	return c.JSON(http.StatusOK, "OK")
}
