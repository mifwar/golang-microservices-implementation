package main

import (
	repo "books/pkg/repository"
	"books/pkg/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	repo.InitDB()

	r := mux.NewRouter()
	routes.SetupHandle(r)

	port := "localhost:8090"
	log.Fatal(http.ListenAndServe(port, r))

}
