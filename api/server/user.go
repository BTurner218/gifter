package server

import (
	"errors"
	"net/http"

	"github.com/BTurner218/gifter/api/gifter"
	"github.com/go-chi/render"
)

type UserRequest struct {
	*gifter.User
}

func (u *UserRequest) Bind(r *http.Request) error {
	if u.User == nil {
		return errors.New("missing required User fields.")
	}

	return nil
}

func (u *UserRequest) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *Server) createUser(w http.ResponseWriter, r *http.Request) {
	data := &UserRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}
	user := data.User
	s.userService.CreateUser(user)

	render.Status(r, http.StatusCreated)
	render.Render(w, r, &UserRequest{})
}
