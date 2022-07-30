package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	echoSwagger "github.com/swaggo/echo-swagger"
	"github.com/zett-8/go-clean-echo/services"
	"strings"
)

type Handlers struct {
	AuthorHandler
	BookHandler
}

func New(s *services.Services) *Handlers {
	return &Handlers{
		AuthorHandler: AuthorHandler{s.AuthorService},
		BookHandler:   BookHandler{s.BookService},
	}
}

func Set(e *echo.Echo, h *Handlers) {
	e.GET("/", IndexHandler)
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	g := e.Group("/api/v1")
	// Author
	g.GET("/author", h.AuthorHandler.GetAuthors)
	g.POST("/author", h.AuthorHandler.CreateAuthor)
	g.PUT("/author", h.AuthorHandler.UpdateAuthorById)
	g.DELETE("/author/:id", h.AuthorHandler.DeleteAuthorById)
	// Book
	g.GET("/book", h.BookHandler.GetBooks)
	g.DELETE("/book/:id", h.BookHandler.DeleteBook)
}

func Echo() *echo.Echo {
	e := echo.New()

	e.Logger.SetLevel(log.DEBUG)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Skipper: func(c echo.Context) bool {
			if strings.Contains(c.Request().URL.Path, "swagger") {
				return true
			}
			return false
		},
	}))

	return e
}
