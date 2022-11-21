package routes

import (
	uc "books/pkg/usecase"

	"github.com/gorilla/mux"
)

func SetupHandle(r *mux.Router) {
	r.HandleFunc("/book/", uc.CreateBook).Methods("POST")
	r.HandleFunc("/book/", uc.GetAllBooks).Methods("GET")
	r.HandleFunc("/book/{id}/", uc.GetBookById).Methods("GET")
	r.HandleFunc("/book/{id}/", uc.DeleteBook).Methods("DELETE")
	r.HandleFunc("/book/{id}/", uc.UpdateBook).Methods("PUT")
}
