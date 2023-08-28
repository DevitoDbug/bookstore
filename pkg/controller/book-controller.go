package controller

import (
	"encoding/json"
	"github.com/DevitoDbug/bookstore/pkg/models"
	"log"
	"net/http"
)

var NewBook models.Book

func GetBookById(w http.ResponseWriter, r *http.Request) {
	//get the id
	//using params
	//get the book form the database
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "pkg/json")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write(res)
	if err != nil {
		log.Printf("%v\n", err)
		return
	}
}

func CreateBook(w http.ResponseWriter, r *http.Request) {

}

func UpdateBook(w http.ResponseWriter, r *http.Request) {

}

func DeleteBook(w http.ResponseWriter, r *http.Request) {

}
