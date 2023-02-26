package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"github.com/zett-8/go-clean-echo/configs"
	"github.com/zett-8/go-clean-echo/services"
	"github.com/zett-8/go-clean-echo/utils"
	"net/http"
	"strings"
)

type Handlers struct {
	AuthorHandler
	BookHandler
}

func New(s *services.Services) *Handlers {
	return &Handlers{
		AuthorHandler: &authorHandler{s.Author},
		BookHandler:   &bookHandler{s.Book},
	}
}

func SetDefault(e *echo.Echo) {
	utils.SetHTMLTemplateRenderer(e)

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "data", configs.Auth0Config)
	})
	e.GET("/healthcheck", HealthCheckHandler)
	e.GET("/swagger/*", echoSwagger.WrapHandler)
}

func SetApi(e *echo.Echo, h *Handlers, m echo.MiddlewareFunc) {
	g := e.Group("/api/v1")
	g.Use(m)

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
