package app

import (
	"net/http"
	// "github.com/ben-hidalgo/gonvd/restful"
)

type health struct {
	Status string `json:"status"`
}

func (s *Server) handleHealth(hideBody bool) http.HandlerFunc {
    
    return func(w http.ResponseWriter, r *http.Request) {
        
        JsonSuccess(w, health{Status: "UP"})
    }
}
