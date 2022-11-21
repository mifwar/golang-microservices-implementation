package main

import (
	"log"
	"net/http"
	repo "users/pkg/repository"
	"users/pkg/routes"

	"github.com/gorilla/mux"
)

func main() {
	repo.InitDB()

	r := mux.NewRouter()
	routes.SetupHandle(r)

	port := "localhost:8080"
	log.Fatal(http.ListenAndServe(port, r))
}
