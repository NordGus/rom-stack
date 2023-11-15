package html

import (
	"embed"
	"errors"
	"html/template"
)

var (
	//go:embed templates
	templates embed.FS
)

type App struct {
	templates *template.Template
}

func New() (*App, error) {
	tmpls, err := template.ParseFS(templates, "templates/*.gohtml")
	if err != nil {
		return nil, errors.Join(errors.New("html: something went wrong parsing templates"), err)
	}

	return &App{
		templates: tmpls,
	}, nil
}
