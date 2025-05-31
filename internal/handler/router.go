package handler

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)


func NewRouter() http.Handler {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to Task Manager API!"))
	})

	// Auth routes
	r.Route("/auth", func(r chi.Router) {
		r.Post("/register", RegisterUserHandler)
		r.Post("/login", LoginUserHandler)
	})
	return r
}
