package routes

import (
	uc "users/pkg/usecase"

	"github.com/gorilla/mux"
)

func SetupHandle(r *mux.Router) {
	r.HandleFunc("/user/", uc.CreateUser).Methods("POST")
	r.HandleFunc("/user/", uc.GetAllUsers).Methods("GET")
	r.HandleFunc("/user/{id}/", uc.GetUserById).Methods("GET")
	r.HandleFunc("/user/{id}/", uc.DeleteUser).Methods("DELETE")
	r.HandleFunc("/user/{id}/", uc.UpdateUser).Methods("PUT")
}
