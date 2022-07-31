package utils

import (
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func SetHTMLTemplateRenderer(e *echo.Echo) {
	t := &Template{
		templates: template.Must(template.ParseGlob("public/index.html")),
	}

	e.Renderer = t
}
