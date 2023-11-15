package staticfileserver

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (a *App) Routes(router chi.Router) {
	router.Get("/dist", http.RedirectHandler("/dist/", http.StatusMovedPermanently).ServeHTTP)
	router.Get("/dist/*", http.FileServer(http.FS(a.dist)).ServeHTTP)
}
