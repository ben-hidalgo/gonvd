package app

import (
	"net/http"
	// "github.com/gorilla/handlers"
	// "log"
)


func (s *Server) handleHealth(hideBody bool) http.HandlerFunc {
    
    return func(w http.ResponseWriter, r *http.Request) {
        
        /*
		  js, err := json.Marshal(profile)
		  if err != nil {
		    http.Error(w, err.Error(), http.StatusInternalServerError)
		    return
		  }

		  w.Header().Set("Content-Type", "application/json")
		  w.Write(js)        
        */
    }
}
