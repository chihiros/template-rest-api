package infra

import (
	"tamaribacms/ent"
	"tamaribacms/interfaces/controller"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func NewRouter(conn *ent.Client) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	controller := controller.NewController(conn)
	r.Route("/api", func(r chi.Router) {
		r.Route("/hello", func(r chi.Router) {
			r.Get("/", func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("Hello, World!"))
			})
		})
	})

	return r
}
