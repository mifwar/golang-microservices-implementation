package usecase

import (
	"fmt"
	"net/http"
	m "users/pkg/models"
	repo "users/pkg/repository"
	u "users/pkg/utils"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	user := m.User{}
	u.ParseBody(r, &user)

	repo.CreateUser(&user)
	u.SendData(w, user)
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users := repo.GetAllUsers()
	u.SendData(w, users)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	id := u.GetParamID(r, "id")
	user := repo.GetUserById(id)

	if user.ID == 0 {
		res := map[string]string{"result": fmt.Sprintf("can't find user with ID = %d", id)}
		u.SendData(w, res)
	} else {
		u.SendData(w, user)
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := u.GetParamID(r, "id")
	user := repo.GetUserById(id)

	if user.ID == 0 {
		res := map[string]string{"result": fmt.Sprintf("can't find user with ID = %d", id)}
		u.SendData(w, res)
	} else {
		users := repo.DeleteUser(id)
		u.SendData(w, users)
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := u.GetParamID(r, "id")
	updated := m.User{}
	u.ParseBody(r, &updated)

	isSuccess := repo.UpdateUser(id, updated)
	if !isSuccess {
		res := map[string]string{"result": fmt.Sprintf("can't find user with ID = %d", id)}
		u.SendData(w, res)
	} else {
		users := repo.GetAllUsers()
		u.SendData(w, users)
	}
}
