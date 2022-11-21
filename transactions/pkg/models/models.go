package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	UserID int `json:"user_id"`
	BookID int `json:"book_id"`
}
