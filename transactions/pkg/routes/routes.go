package routes

import (
	uc "transactions/pkg/usecase"

	"github.com/gorilla/mux"
)

func SetupHandle(r *mux.Router) {
	r.HandleFunc("/transaction/", uc.CreateTransaction).Methods("POST")
	r.HandleFunc("/transaction/", uc.GetAllTransactions).Methods("GET")
	r.HandleFunc("/transaction/{id}/", uc.GetTransactionById).Methods("GET")
	r.HandleFunc("/transaction/{id}/", uc.DeleteTransaction).Methods("DELETE")
	r.HandleFunc("/transaction/{id}/", uc.UpdateTransaction).Methods("PUT")
}
