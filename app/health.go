package app

import (
	"log"
	"net/http"
	// "github.com/ben-hidalgo/gonvd/restful"
)

type health struct {
	Status string `json:"status,omitempty"`
}

// Passing in parameters is one reason why the HandlerFunc wrapper is useful.

func (s *Server) handleHealth(showBody bool) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		body := &health{}

		if showBody {
			log.Println("showBody true")
			body.Status = "UP"
		}

		JsonSuccess(w, body)
	}
}
