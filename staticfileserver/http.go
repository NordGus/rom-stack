package staticfileserver

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (s *Service) Routes(router chi.Router) {
	router.Get("/dist", http.RedirectHandler("/dist/", http.StatusMovedPermanently).ServeHTTP)
	router.Get("/dist/*", http.FileServer(http.FS(s.dist)).ServeHTTP)
}
