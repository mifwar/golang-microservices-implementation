package main

import (
	"log"
	"net/http"
	repo "transactions/pkg/repository"
	"transactions/pkg/routes"

	"github.com/gorilla/mux"
)

func main() {
	repo.InitDB()

	r := mux.NewRouter()
	routes.SetupHandle(r)

	port := "localhost:8070"
	log.Fatal(http.ListenAndServe(port, r))

}
