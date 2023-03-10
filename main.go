package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"week-1/src/router"
	"week-1/src/setup"
)

func main() {
	setup.Connection()
	r := mux.NewRouter()
	router.Register(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", r))
}
