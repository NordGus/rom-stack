package html

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (a *App) Routes(router chi.Router) {
	router.Get("/", func(writer http.ResponseWriter, _ *http.Request) {
		err := a.templates.ExecuteTemplate(writer, "layout.gohtml", nil)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}
	})

	router.Get("/app", func(writer http.ResponseWriter, _ *http.Request) {
		err := a.templates.ExecuteTemplate(writer, "app.gohtml", nil)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}
	})
}
