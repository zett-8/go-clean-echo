package handlers

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"github.com/zett-8/go-clean-echo/services"
	"github.com/zett-8/go-clean-echo/utils"
	"log"
	"net/http"
	"strconv"
)

type BookHandler struct {
	services.BookService
}

// GetBooks
// @Summary Fetch a list of all books.
// @Description Fetch a list of all books.
// @Tags Book
// @Accept */*
// @Security Bearer Authentication
// @Param Authorization header string true "'Bearer _YOUR_TOKEN_'"
// @Produce json
// @Success 200 {object} []models.Book
// @Failure 500 {object} utils.Error
// @Router /api/v1/book [get]
func (h *BookHandler) GetBooks(c echo.Context) error {
	r, err := h.BookService.GetBooks()

	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, utils.Error{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, r)
}

// DeleteBook
// @Summary Delete a book by ID.
// @Description Delete a book by ID.
// @Tags Book
// @Accept */*
// @Security Bearer Authentication
// @Produce json
// @Param id path int true "Book id"
// @Param Authorization header string true "'Bearer _YOUR_TOKEN_'"
// @Success 200 {integer} int "Deleted Book ID"
// @Failure 400 {object} utils.Error
// @Failure 404 {object} utils.Error
// @Failure 500 {object} utils.Error
// @Router /api/v1/book/{id} [delete]
func (h *BookHandler) DeleteBook(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, utils.Error{Message: "ID is Invalid"})
	}

	err = h.BookService.DeleteBookById(id)

	if err == sql.ErrNoRows {
		return c.JSON(http.StatusNotFound, utils.Error{Message: "not found"})
	} else if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, utils.Error{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, "OK")
}
