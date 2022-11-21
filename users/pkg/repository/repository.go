package repository

import (
	d "users/pkg/database"
	m "users/pkg/models"

	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() {
	db = d.GetDB()
	db.AutoMigrate(&m.User{})
}

func CreateUser(u *m.User) {
	db.Create(u)
}

func GetAllUsers() []m.User {
	var users []m.User
	db.Find(&users)
	return users
}

func GetUserById(id int) *m.User {
	var user *m.User
	db.Where("ID = ?", id).Find(&user)
	return user
}

func DeleteUser(id int) []m.User {
	var user *m.User
	var remainedUsers []m.User
	db.Where("ID = ?", id).Delete(&user)
	db.Find(&remainedUsers)
	return remainedUsers
}

func UpdateUser(id int, updated m.User) bool {
	var user *m.User

	db.Where("ID = ?", id).Find(&user)

	if user.ID == 0 {
		return false
	}

	user.Name = updated.Name
	user.IsBorrowing = updated.IsBorrowing
	db.Save(user)
	return true
}
