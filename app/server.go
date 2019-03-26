package app

import (
	"github.com/gorilla/mux"
)

type Server struct {
	Config Config
	Router *mux.Router
}
