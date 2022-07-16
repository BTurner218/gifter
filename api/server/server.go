package server

import (
	"net/http"
	"time"

	"github.com/BTurner218/gifter/api/gifter"
	"github.com/BTurner218/gifter/api/postgres"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type Server struct {
	server      *http.Server
	router      *chi.Mux
	userService gifter.UserService
}

type ErrResponse struct {
	Err            error `json:"-"`
	HTTPStatusCode int   `json:"-"`

	StatusText string `json:"status"`
	AppCode    int64  `json:"code,omitempty"`
	ErrorText  string `json:"error,omitempty"`
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func ErrInvalidRequest(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 400,
		StatusText:     "Invalid request.",
		ErrorText:      err.Error(),
	}
}

func NewServer(db *postgres.DB) *Server {
	s := Server{
		server: &http.Server{
			Addr:         ":8080",
			WriteTimeout: 10 * time.Second,
			ReadTimeout:  10 * time.Second,
		},
		router: chi.NewRouter(),
	}

	s.routes()
	s.userService = postgres.NewUserService(db)
	s.server.Handler = s.router

	return &s
}

func (s *Server) Run() error {
	return s.server.ListenAndServe()
}
