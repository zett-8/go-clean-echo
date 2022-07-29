package services

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"github.com/zett-8/go-clean-echo/stores"
	"log"
	"net/http"
	"strconv"
)

type AuthorService struct {
	store *stores.AuthorStore
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
func (s *AuthorService) GetAuthors(c echo.Context) error {
	r, err := s.store.Get()

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
// @Failure 500 {object} utils.Error
// @Router /api/v1/author/{id} [delete]
func (s *AuthorService) DeleteAuthor(c echo.Context) error {
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
