package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/nunusavi/task-manager/internal/middleware"
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

	// Protected routes
	r.Group(func(r chi.Router) {
		r.Use(middleware.Verifier())
		r.Use(middleware.Authenticator())

		// Your new /tasks routes
		r.Route("/tasks", func(r chi.Router) {
			r.Get("/", ListTasksHandler)
			r.Post("/", CreateTaskHandler)
		})
	})
	return r
}
