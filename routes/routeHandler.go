package routes

import (
	"github.com/go-chi/chi"
	"net/http"
)

func GetRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Group(func(router chi.Router) {
		router.Get("/", Hello)
	})

	return router
}

func Hello(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Hello World!!!"))
}
