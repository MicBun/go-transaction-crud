package models

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

type Transaction struct {
	ID                 uint                `json:"id" gorm:"primary_key"`
	TransactionType    bool                `json:"transaction_type" gorm:"type:bool"`
	AdminID            uint                `json:"admin_id" gorm:"type:int"`
	TransactionDetails []TransactionDetail `json:"transaction_details" gorm:"foreignKey:TransactionID"`
	CreatedAt          time.Time           `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt          time.Time           `json:"updated_at" gorm:"autoUpdateTime"`
}

func GetTransactions(ctx *gin.Context) ([]Transaction, error) {
	db := ctx.MustGet("db").(*gorm.DB)
	var transactions []Transaction
	err := db.Find(&transactions).Error
	if err != nil {
		return []Transaction{}, err
	}
	return transactions, nil
}

func (t *Transaction) GetTransaction(ctx *gin.Context) (Transaction, error) {
	db := ctx.MustGet("db").(*gorm.DB)
	var transaction Transaction
	err := db.Where("id = ?", t.ID).First(&transaction).Error
	if err != nil {
		return Transaction{}, err
	}
	return transaction, nil
}

func (t *Transaction) CreateTransaction(ctx *gin.Context) (Transaction, error) {
	db := ctx.MustGet("db").(*gorm.DB)
	err := db.Create(&t).Error
	if err != nil {
		return Transaction{}, err
	}
	for _, transactionDetail := range t.TransactionDetails {
		transactionDetail.TransactionID = t.ID
		transactionDetail.CreateTransactionDetail(ctx)
	}
	return *t, nil
}

func (t *Transaction) UpdateTransaction(ctx *gin.Context) (Transaction, error) {
	db := ctx.MustGet("db").(*gorm.DB)
	err := db.Model(&t).Updates(t).Error
	if err != nil {
		return Transaction{}, err
	}
	return *t, nil
}

func (t *Transaction) DeleteTransaction(ctx *gin.Context) error {
	db := ctx.MustGet("db").(*gorm.DB)
	err := db.Delete(&t).Error
	if err != nil {
		return err
	}
	return nil
}
