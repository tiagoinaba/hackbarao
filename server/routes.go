package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (s *Server) RegisterRoutes() {
	mux := chi.NewRouter()
	mux.Use(middleware.Logger)
}
