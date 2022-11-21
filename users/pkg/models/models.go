package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name        string `json:"name"`
	IsBorrowing bool   `json:"isBorrowing"`
}

//user       : id, name, is_borrowing
//transaction: user_id, book_id, date
//catalogue  : book_id, title, published, genre
