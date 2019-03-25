package main

import (
	// "github.com/gorilla/handlers"
	"github.com/ben-hidalgo/gonvd/common"
	"log"
)

// https://medium.com/statuscode/how-i-write-go-http-services-after-seven-years-37c208122831

func main() {

	config := &common.Config{}

	config.Init()

	log.Printf("main() listening on %s", config.MuxAddr)

	//log.Fatal(http.ListenAndServe(muxAddr, handlers.CORS()(server.Router)))
}
