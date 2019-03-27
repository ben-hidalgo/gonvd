package app

import (
	"encoding/json"
	"github.com/ben-hidalgo/gonvd/restful"
	"log"
	"net/http"
)

func JsonResponse(w http.ResponseWriter, body interface{}, status int) {

	j, err := json.Marshal(body)

	if err != nil {

		log.Printf("JsonResponse() err=%s", err)

		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	w.Header().Set(restful.HEADER_ContentType, restful.MIME_JSON)

	w.WriteHeader(status)

	w.Write(j)
}

func JsonSuccess(w http.ResponseWriter, body interface{}) {
	JsonResponse(w, body, http.StatusOK)
}
