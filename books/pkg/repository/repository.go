package repository

import (
	"gorm.io/gorm"

	d "books/pkg/database"
	m "books/pkg/models"
)

var db *gorm.DB

func InitDB() {
	db = d.GetDB()
	db.AutoMigrate(&m.Book{})
}

func CreateBook(b *m.Book) {
	db.Create(b)
}

func GetAllBooks() []m.Book {
	var books []m.Book
	db.Find(&books)
	return books
}

func GetBookById(id int) m.Book {
	var book m.Book
	db.Where("ID = ?", id).Find(&book)
	return book
}

func DeleteBook(id int) []m.Book {
	var book m.Book
	var remainedBooks []m.Book
	db.Where("ID = ?", id).Delete(&book)

	db.Find(&remainedBooks)
	return remainedBooks
}

func UpdateBook(id int, updatedBook m.Book) bool {
	var book m.Book
	db.Where("ID = ?", id).Find(&book)

	if book.ID == 0 {
		return false
	}

	book.Author = updatedBook.Author
	book.Title = updatedBook.Title
	book.Author = updatedBook.Author
	book.Genre = updatedBook.Genre
	book.Publisher = updatedBook.Publisher

	db.Save(&book)

	return true
}
