package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/alecthomas/kingpin"

	"zavamed.com/parcel-company/data"
	"zavamed.com/parcel-company/handlers"
)

var (
	port = kingpin.Flag("port", "Server port").Default("8080").Int()
)

func main() {
	kingpin.Parse()

	address := fmt.Sprintf(":%d", *port)

	myApp := data.NewApp()

	router := handlers.CreateRouter(myApp)

	log.Fatal(http.ListenAndServe(address, router))
}
