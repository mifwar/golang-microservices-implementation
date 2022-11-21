package usecase

import (
	m "books/pkg/models"
	repo "books/pkg/repository"
	u "books/pkg/utils"
	"fmt"
	"net/http"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book m.Book
	u.ParseBody(r, &book)

	repo.CreateBook(&book)
	u.SendData(w, book)
}

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	books := repo.GetAllBooks()
	u.SendData(w, books)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	id := u.GetParamID(r, "id")
	book := repo.GetBookById(id)

	if book.ID == 0 {
		res := map[string]string{"result": fmt.Sprintf("can't get book with ID =  %d", id)}
		u.SendData(w, res)
	} else {
		u.SendData(w, book)
	}
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	id := u.GetParamID(r, "id")
	book := repo.GetBookById(id)

	if book.ID == 0 {
		res := map[string]string{"result": fmt.Sprintf("can't get book with ID = %d", id)}
		u.SendData(w, res)
	} else {
		books := repo.DeleteBook(id)
		u.SendData(w, books)
	}
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	id := u.GetParamID(r, "id")
	updated := m.Book{}
	u.ParseBody(r, &updated)

	isSuccess := repo.UpdateBook(id, updated)
	if !isSuccess {
		res := map[string]string{"result": fmt.Sprintf("can't get book with ID = %d", id)}
		u.SendData(w, res)
	} else {
		books := repo.GetAllBooks()
		u.SendData(w, books)
	}
}
