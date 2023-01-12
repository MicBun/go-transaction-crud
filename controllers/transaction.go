package controllers

import (
	"github.com/MicBun/go-transaction-crud/models"
	"github.com/gin-gonic/gin"
)

const (
	STOCK_OUT = false
	STOCK_IN  = true
)

type TransactionDetailInput struct {
	ProductID uint `json:"product_id" binding:"required"`
	Quantity  int  `json:"quantity" binding:"required"`
}

type TransactionInput struct {
	TransactionType    bool                     `json:"transaction_type" binding:"required"`
	AdminID            uint                     `json:"admin_id" binding:"required"`
	TransactionDetails []TransactionDetailInput `json:"transaction_details" binding:"required"`
}

type TransactionDetailOutput struct {
	ProductID     uint `json:"product_id"`
	Quantity      int  `json:"quantity"`
	TransactionID uint `json:"transaction_id"`
}

type TransactionOutput struct {
	TransactionType    bool                      `json:"transaction_type"`
	AdminID            uint                      `json:"admin_id"`
	TransactionDetails []TransactionDetailOutput `json:"transaction_details"`
}

// CreateTransaction godoc
// @Summary Create a transaction.
// @Description Create a transaction.
// @Tags Transaction
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Param transaction body TransactionInput true "Transaction"
// @Success 200 {object} TransactionOutput
// @Failure 400 {object} string
// @Router /transactions [post]
func CreateTransaction(c *gin.Context) {
	var input TransactionInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	transaction := models.Transaction{
		TransactionType: input.TransactionType,
		AdminID:         input.AdminID,
	}

	transactionDetails := []models.TransactionDetail{}
	for _, transactionDetail := range input.TransactionDetails {
		if transactionDetail.Quantity <= 0 {
			c.JSON(400, gin.H{"error": "Quantity must be greater than 0"})
			return
		}
		product := models.Product{ID: transactionDetail.ProductID}
		product, err := product.GetProduct(c)
		if err != nil {
			c.JSON(400, gin.H{"error": "Product not found"})
			return
		}
		if input.TransactionType == STOCK_OUT && product.Stock < transactionDetail.Quantity {
			c.JSON(400, gin.H{"error": "Stock not enough"})
			return
		}
		transactionDetails = append(transactionDetails, models.TransactionDetail{
			ProductID: transactionDetail.ProductID,
			Quantity:  transactionDetail.Quantity,
		})
	}

	transaction, err := transaction.CreateTransaction(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": transaction})
}

// GetTransactions godoc
// @Summary Get all transactions.
// @Description Get all transactions.
// @Tags Transaction
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} []models.Transaction
// @Failure 400 {object} string
// @Router /transactions [get]
func GetTransactions(c *gin.Context) {
	transactions, err := models.GetTransactions(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": transactions})
}

func GetTransaction(c *gin.Context) {
	var transaction models.Transaction
	if err := c.ShouldBindUri(&transaction); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	transaction, err := transaction.GetTransaction(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": transaction})
}
