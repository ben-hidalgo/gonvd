package main

import (
	// "github.com/gorilla/handlers"
	"github.com/ben-hidalgo/gonvd/app"
	"net/http"
	"log"
)

// https://medium.com/statuscode/how-i-write-go-http-services-after-seven-years-37c208122831

func main() {

	config := &app.Config{}

	config.Init()

	log.Printf("main() listening on %s", config.MuxAddr)

	server := &app.Server{}

	server.InitRoutes()

	log.Fatal(http.ListenAndServe(config.MuxAddr, server.Router))
}
