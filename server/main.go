package main

import (
	"log"
	"net/http"

	"google.golang.org/appengine"

	"github.com/abhijit-kar/unite-society/apis"
)

func main() {
	log.Printf("Server started")

	router := apis.NewRouter()

	http.Handle("/", router)

	appengine.Main()

	// log.Fatal(http.ListenAndServe(":8080", router))
}
