package main

import (
	"io"
	"text/template"

	"github.com/labstack/echo/v4"
)

type TemplateRegistry struct {
	templates map[string]*template.Template
}

func (t *TemplateRegistry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates[name].Execute(w, data)
}

func RegisterTemplates(htmlPages []HtmlPage) *TemplateRegistry {
	templates := make(map[string]*template.Template)
	for _, page := range htmlPages {
		templates[page.Name] = template.Must(
			template.ParseFiles(
				page.Path,
				"templates/header.html",
				"templates/footer.html"))
	}
	t := &TemplateRegistry{
		templates: templates,
	}
	return t
}
