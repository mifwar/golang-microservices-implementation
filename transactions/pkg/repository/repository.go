package repository

import (
	d "transactions/pkg/database"
	m "transactions/pkg/models"

	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() {
	db = d.GetDB()
	db.AutoMigrate(&m.Transaction{})
}

func CreateTransaction(t *m.Transaction) {
	db.Create(t)
}

func GetAllTransactions() []m.Transaction {
	var transactions []m.Transaction
	db.Find(&transactions)
	return transactions
}

func GetTransactionById(id int) m.Transaction {
	var transaction m.Transaction
	db.Where("ID = ?", id).Find(&transaction)
	return transaction
}

func DeleteTransaction(id int) []m.Transaction {
	var transaction m.Transaction
	var remainedTransactions []m.Transaction
	db.Where("ID = ?", id).Delete(&transaction)

	db.Find(&remainedTransactions)
	return remainedTransactions
}

func UpdateTransaction(id int, updatedTransaction m.Transaction) bool {
	var transaction m.Transaction
	db.Where("ID = ?", id).Find(&transaction)

	if transaction.ID == 0 {
		return false
	}

	transaction.UserID = updatedTransaction.UserID
	transaction.BookID = updatedTransaction.BookID

	db.Save(&transaction)

	return true
}
