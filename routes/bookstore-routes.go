package routes

import (
	"github.com/DevitoDbug/bookstore/pkg/controller"
	"github.com/gorilla/mux"
)

var RegiserBookStore = func(r *mux.Router) {
	r.HandleFunc("/book/", controller.CreateBook).Methods("POST")
	r.HandleFunc("/book", controller.GetBook).Methods("GET")
	r.HandleFunc("/book", controller.UpdateBook).Methods("PUT")
	r.HandleFunc("/book", controller.DeleteBook).Methods("DELETE")
	r.HandleFunc("/book", controller.GetBooks).Methods("GET")

}
