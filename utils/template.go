package utils

import (
	"github.com/labstack/echo/v4"
	"github.com/zett-8/go-clean-echo/public"
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
		templates: template.Must(template.ParseFS(public.Files, "*.html")),
	}

	e.Renderer = t
}
