package handlers

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"github.com/zett-8/go-clean-echo/services"
	"log"
	"net/http"
	"strconv"
)

type AuthorHandler struct {
	*services.AuthorService
}

// GetAuthors
// @Summary Fetch a list of all authors.
// @Description Fetch a list of all authors.
// @Tags Author
// @Accept */*
// @Produce json
// @Success 200 {object} []models.Author
// @Failure 500 {object} utils.Error
// @Router /api/v1/author [get]
func (h *AuthorHandler) GetAuthors(c echo.Context) error {
	r, err := h.AuthorService.GetAuthors()

	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]error{"message": err})
	}

	return c.JSON(http.StatusOK, r)
}

// DeleteAuthor
// @Summary Delete an author by ID.
// @Description Delete an author by ID.
// @Tags Author
// @Accept */*
// @Produce json
// @Param id path int true "Author id"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "ID is invalid"
// @Failure 404 {string} string "Not found"
// @Failure 500 {object} utils.Error
// @Router /api/v1/author/{id} [delete]
func (h *AuthorHandler) DeleteAuthor(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "ID is invalid")
	}

	err = h.AuthorService.DeleteAuthor(id)

	if err == sql.ErrNoRows {
		return c.JSON(http.StatusNotFound, "not found")
	} else if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]error{"message": err})
	}

	return c.JSON(http.StatusOK, "OK")
}
