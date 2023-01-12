package models

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

const (
	STOCK_OUT = 0
	STOCK_IN  = 1
)

type TransactionDetail struct {
	ID            uint      `json:"id" gorm:"primary_key"`
	TransactionID uint      `json:"transaction_id" gorm:"type:int"`
	ProductID     uint      `json:"product_id" gorm:"type:int"`
	Quantity      int       `json:"quantity" gorm:"type:int"`
	CreatedAt     time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt     time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (t *TransactionDetail) GetTransactionDetails(ctx *gin.Context) ([]TransactionDetail, error) {
	db := ctx.MustGet("db").(*gorm.DB)
	var transactionDetails []TransactionDetail
	err := db.Find(&transactionDetails).Error
	if err != nil {
		return []TransactionDetail{}, err
	}
	return transactionDetails, nil
}

func (t *TransactionDetail) GetTransactionDetail(ctx *gin.Context) (TransactionDetail, error) {
	db := ctx.MustGet("db").(*gorm.DB)
	var transactionDetail TransactionDetail
	err := db.Where("id = ?", t.ID).First(&transactionDetail).Error
	if err != nil {
		return TransactionDetail{}, err
	}
	return transactionDetail, nil
}

func (t *TransactionDetail) CreateTransactionDetail(ctx *gin.Context) (TransactionDetail, error) {
	db := ctx.MustGet("db").(*gorm.DB)
	err := db.Create(&t).Error
	if err != nil {
		return TransactionDetail{}, err
	}
	return *t, nil
}

func (t *TransactionDetail) UpdateTransactionDetail(ctx *gin.Context) (TransactionDetail, error) {
	db := ctx.MustGet("db").(*gorm.DB)
	err := db.Model(&t).Updates(t).Error
	if err != nil {
		return TransactionDetail{}, err
	}
	return *t, nil
}

func (t *TransactionDetail) DeleteTransactionDetail(ctx *gin.Context) error {
	db := ctx.MustGet("db").(*gorm.DB)
	err := db.Delete(&t).Error
	if err != nil {
		return err
	}
	return nil
}
