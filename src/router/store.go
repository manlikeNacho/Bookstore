package router

import (
	"github.com/gorilla/mux"
	"week-1/src/controller"
)

var Register = func(router *mux.Router) {
	router.HandleFunc("/createBooks", controller.CreateBook).Methods("POST")
	router.HandleFunc("/getAllBooks", controller.GetAllBooks).Methods("GET")
	router.HandleFunc("/book/{Id}", controller.GetBookById).Methods("GET")
}
