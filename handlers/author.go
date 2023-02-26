package handlers

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"github.com/zett-8/go-clean-echo/logger"
	"github.com/zett-8/go-clean-echo/models"
	"github.com/zett-8/go-clean-echo/services"
	"github.com/zett-8/go-clean-echo/utils"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

type (
	AuthorHandler interface {
		GetAuthors(c echo.Context) error
		CreateAuthor(c echo.Context) error
		UpdateAuthorById(c echo.Context) error
		DeleteAuthorById(c echo.Context) error
	}

	authorHandler struct {
		services.AuthorService
	}
)

// GetAuthors
// @Summary Fetch a list of all authors.
// @Description Fetch a list of all authors.
// @Tags Author
// @Accept */*
// @Param Authorization header string true "'Bearer _YOUR_TOKEN_'"
// @Security Bearer Authentication
// @Produce json
// @Success 200 {object} []models.Author
// @Failure 500 {object} utils.Error
// @Router /api/v1/author [get]
func (h *authorHandler) GetAuthors(c echo.Context) error {
	r, err := h.AuthorService.GetAuthors()

	if err != nil {
		logger.Error("failed to get author", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, utils.Error{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, r)
}

// CreateAuthor
// @Summary Create an author.
// @Description Create an author.
// @Tags Author
// @Accept */*
// @Param Authorization header string true "'Bearer _YOUR_TOKEN_'"
// @Security Bearer Authentication
// @Produce json
// @Success 200 {integer} integer "Created ID"
// @Failure 400 {object} utils.Error
// @Failure 500 {object} utils.Error
// @Router /api/v1/author [post]
func (h *authorHandler) CreateAuthor(c echo.Context) error {
	var a *models.Author

	if err := c.Bind(&a); err != nil {
		logger.Error("failed to bind req body", zap.Error(err))
		return c.JSON(http.StatusBadRequest, utils.Error{Message: err.Error()})
	}

	r, err := h.AuthorService.CreateAuthor(a)
	if err != nil {
		logger.Error("failed to create author", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, utils.Error{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, r)
}

// UpdateAuthorById
// @Summary Update an author.
// @Description Update an author.
// @Tags Author
// @Accept */*
// @Param Authorization header string true "'Bearer _YOUR_TOKEN_'"
// @Security Bearer Authentication
// @Produce json
// @Success 200 {integer} integer "Updated ID"
// @Failure 400 {object} utils.Error
// @Failure 404 {object} utils.Error
// @Failure 500 {object} utils.Error
// @Router /api/v1/author [put]
func (h *authorHandler) UpdateAuthorById(c echo.Context) error {
	var a *models.Author

	if err := c.Bind(&a); err != nil {
		logger.Error("failed to bind req body", zap.Error(err))
		return c.JSON(http.StatusBadRequest, utils.Error{Message: "args is invalid"})
	}

	r, err := h.AuthorService.UpdateAuthorById(a)
	if err == sql.ErrNoRows {
		return c.JSON(http.StatusNotFound, utils.Error{Message: "not found"})
	} else if err != nil {
		logger.Error("failed to update author", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, utils.Error{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, r)
}

// DeleteAuthorById
// @Summary Delete an author by ID.
// @Description Delete an author by ID.
// @Tags Author
// @Accept */*
// @Security Bearer Authentication
// @Produce json
// @Param id path int true "Author id"
// @Param Authorization header string true "'Bearer _YOUR_TOKEN_'"
// @Success 200 {string} string "OK"
// @Failure 400 {object} utils.Error
// @Failure 404 {object} utils.Error
// @Failure 500 {object} utils.Error
// @Router /api/v1/author/{id} [delete]
func (h *authorHandler) DeleteAuthorById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error("failed to parse id", zap.Error(err))
		return c.JSON(http.StatusBadRequest, utils.Error{Message: "ID is invalid"})
	}

	err = h.AuthorService.DeleteAuthor(id)

	if err == sql.ErrNoRows {
		return c.JSON(http.StatusNotFound, "not found")
	} else if err != nil {
		logger.Error("failed to delete author", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, utils.Error{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, "OK")
}
