package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"week-1/src/router"
	"week-1/src/setup"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world from Deez suacse, %s", r.URL.Path[1:])
}

func handleSomn(r http.ResponseWriter, v *http.Request) {

}

func main() {
	setup.Connection()
	r := mux.NewRouter()
	router.Register(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", r))
}
