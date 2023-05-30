package main

import (
	"log"
	"net/http"

	"github.com/SemaMolchanov/go-bookstore/packages/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterURLs(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
