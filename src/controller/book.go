package controller

import (
	"github.com/gorilla/mux"
	"net/http"
	"week-1/src/models"
)

var NewBook models.Book

func createBook(w http.ResponseWriter, r *http.Request) {
	_ = mux.Vars(r)
}
