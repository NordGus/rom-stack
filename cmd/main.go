package main

import (
	"github.com/NordGus/rom-stack/html"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
)

func main() {
	htmlapp, err := html.New()
	if err != nil {
		log.Fatalln(err)
	}

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(devCORSMiddleware)

	htmlapp.Routes(router)

	err = http.ListenAndServe(":4269", router)
	if err != nil {
		log.Fatalf("something went wrong initalizing http server: %v\n", err)
	}
}

func devCORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")

		next.ServeHTTP(writer, request)
	})
}
