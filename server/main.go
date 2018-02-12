package main

import (
	"log"
	"net/http"

	"google.golang.org/appengine"

	apis "github.com/abhijit-kar/restful-go/apis"
)

func main() {
	log.Printf("Server started")

	router := apis.NewRouter()

	http.Handle("/", router)

	appengine.Main()
}
