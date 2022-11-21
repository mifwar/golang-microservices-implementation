package usecase

import (
	"fmt"
	"net/http"
	d "transactions/pkg/domain"
	m "transactions/pkg/models"
	repo "transactions/pkg/repository"
	u "transactions/pkg/utils"
)

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var transaction m.Transaction
	u.ParseBody(r, &transaction)

	user_id := transaction.UserID
	book_id := transaction.BookID

	dataUser, err := d.FetchUserID(user_id)
	dataBook, err := d.FetchBookID(book_id)

	u.LogError(err, "fetch user id")

	userID := dataUser["ID"]
	bookID := dataBook["ID"]

	if userID == 0 || userID == nil {
		res := map[string]string{"result": fmt.Sprintf("can't find user with ID = %d", user_id)}
		u.SendData(w, res)
	} else if bookID == 0 || bookID == nil {
		res := map[string]string{"result": fmt.Sprintf("can't find book with ID = %d", book_id)}
		u.SendData(w, res)
	} else {
		repo.CreateTransaction(&transaction)
		u.SendData(w, transaction)
	}
}

func GetAllTransactions(w http.ResponseWriter, r *http.Request) {
	transactions := repo.GetAllTransactions()
	u.SendData(w, transactions)
}

func GetTransactionById(w http.ResponseWriter, r *http.Request) {
	id := u.GetParamID(r, "id")
	transaction := repo.GetTransactionById(id)

	if transaction.ID == 0 {
		res := map[string]string{"result": fmt.Sprintf("can't find transaction with ID = %d", id)}
		u.SendData(w, res)
	} else {
		u.SendData(w, transaction)
	}
}

func DeleteTransaction(w http.ResponseWriter, r *http.Request) {
	id := u.GetParamID(r, "id")
	transaction := repo.GetTransactionById(id)

	if transaction.ID == 0 {
		res := map[string]string{"result": fmt.Sprintf("can't find transaction with ID = %d", id)}
		u.SendData(w, res)
	} else {
		transactions := repo.DeleteTransaction(id)
		u.SendData(w, transactions)
	}
}

func UpdateTransaction(w http.ResponseWriter, r *http.Request) {
	id := u.GetParamID(r, "id")
	updated := m.Transaction{}
	u.ParseBody(r, &updated)

	isSuccess := repo.UpdateTransaction(id, updated)
	if !isSuccess {
		res := map[string]string{"result": fmt.Sprintf("can't find transaction with ID = %d", id)}
		u.SendData(w, res)
	} else {
		transactions := repo.GetAllTransactions()
		u.SendData(w, transactions)
	}
}
