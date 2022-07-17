package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

func (s *Server) routes() {
	s.router.Use(middleware.Logger)
	s.router.Use(render.SetContentType(render.ContentTypeJSON))

	s.router.Get("/", hello)
	s.router.Route("/user", func(r chi.Router) {
		r.Post("/", s.createUser)
	})
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}
