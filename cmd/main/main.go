package main

import (
	"github.com/DevitoDbug/bookstore/pkg/routes"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	port := ":8080"
	r := mux.NewRouter()

	routes.RegiserBookStore(r)
	http.Handle("/", r)

	err := http.ListenAndServe(port, r)
	if err != nil {
		log.Fatalf("%v\n", err)
		return
	}
}
