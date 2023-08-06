package main

import (
	"html/template"
	"io"
	"net/http"
	"github.com/labstack/echo/v4"
)

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("./views/*.html")),
	}
	e.Renderer = renderer

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", map[string]interface{}{
			"appName": "naagaad",
		})
	})

	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "I am live :)")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
