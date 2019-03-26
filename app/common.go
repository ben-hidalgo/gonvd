package app

import (
	"net/http"
	"encoding/json"
)

func JsonResponse(w http.ResponseWriter, body interface{}, status int) {
  
  j, _ := json.Marshal(body)

  w.Header().Set("Content-Type", "application/json")

  w.WriteHeader(status)

  w.Write(j)
}

func JsonSuccess(w http.ResponseWriter, body interface{}) {
	JsonResponse(w, body, http.StatusOK)  
}
