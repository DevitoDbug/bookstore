package controller

import (
	"encoding/json"
	"fmt"
	"github.com/DevitoDbug/bookstore/pkg/models"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

var NewBook models.Book

func GetBookById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
		return
	}

	book, dbOperation := models.GetBookById(id)
	if dbOperation.Error != nil {
		log.Printf("DB: %v\n", dbOperation.Error)
		return
	}
	res, _ := json.Marshal(book)

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(res)
	if err != nil {
		log.Printf("%v\n", err)
		return
	}

}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "pkglication/json")
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
