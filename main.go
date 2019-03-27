package main

import (
	// "github.com/gorilla/handlers"
	"github.com/ben-hidalgo/gonvd/app"
	"log"
	"net/http"
)

// https://medium.com/statuscode/how-i-write-go-http-services-after-seven-years-37c208122831

func main() {

	config := &app.Config{}

	err := config.Init()
	if err != nil {
		log.Fatal("main() failed config.Init()", err)
	}

	server := &app.Server{}

	cveStore, err := config.InitCveStore()
	if err != nil {
		log.Fatal("main() failed config.InitCveStore()", err)
	}

	server.CVEStore = cveStore

	server.InitRoutes()

	log.Printf("main() listening on %s", config.MuxAddr)

	log.Fatal(http.ListenAndServe(config.MuxAddr, server.Router))
}
