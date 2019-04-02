package app

import (
	"github.com/gorilla/mux"
)

func (s *Server) InitRoutes() {

	s.Router = mux.NewRouter()

	s.Router.HandleFunc("/", s.handleHealth(false))
	s.Router.HandleFunc("/health", s.handleHealth(true))

	s.Router.HandleFunc("/cveitems", s.GetCveItems())
	s.Router.HandleFunc("/cveitems/{id}", s.GetCveItem())
}
