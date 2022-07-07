package server

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	server *http.Server
	router *chi.Mux
}

func NewServer() *Server {
	s := Server{
		server: &http.Server{
			Addr:         ":8080",
			WriteTimeout: 10 * time.Second,
			ReadTimeout:  10 * time.Second,
		},
		router: chi.NewRouter(),
	}

	s.routes()
	s.server.Handler = s.router

	return &s
}

func (s *Server) Run() error {
	return s.server.ListenAndServe()
}
