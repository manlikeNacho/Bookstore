package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"week-1/src/models"
	"week-1/src/utils"
)

var NewBook models.Book

func CreateBook(w http.ResponseWriter, r *http.Request) {
	Book := &models.Book{}
	utils.ParseBody(r, Book)
	b := Book.CreateBook()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	books := models.GetAllBooks()
	res, _ := json.Marshal(books)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["Id"]
	Id, err := strconv.ParseInt(id, 0, 0)

	if err != nil {
		fmt.Println("error while parsing")
	}

	book := models.GetBooksById(int(Id))
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
