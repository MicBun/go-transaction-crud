package controllers

import (
	"github.com/MicBun/go-transaction-crud/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

type CreateProductInput struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	ImgUrl      string `json:"img_url" binding:"required"`
}

type UpdateProductInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	ImgUrl      string `json:"img_url"`
}

// CreateProduct godoc
// @Summary Create a product.
// @Description Create a product.
// @Tags Product
// @Accept  json
// @Produce  json
// @Param product body CreateProductInput true "Product"
// @Success 200 {object} models.Product
// @Router /products [post]
func CreateProduct(c *gin.Context) {
	// Get model if exist
	var input CreateProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	product := models.Product{Name: input.Name, Description: input.Description, ImgUrl: input.ImgUrl}
	productCreated, err := product.CreateProduct(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"product": productCreated})
}

// GetProducts godoc
// @Summary Get all products.
// @Description Get all products.
// @Tags Product
// @Accept  json
// @Produce  json
// @Success 200 {object} []models.Product
// @Router /products [get]
func GetProducts(c *gin.Context) {
	products, err := models.GetProducts(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"products": products})
}

// GetProduct godoc
// @Summary Get a product.
// @Description Get a product.
// @Tags Product
// @Accept  json
// @Produce  json
// @Param id path int true "Product ID"
// @Success 200 {object} models.Product
// @Router /products/{id} [get]
func GetProduct(c *gin.Context) {
	var product models.Product
	id, _ := strconv.Atoi(c.Param("id"))
	product = models.Product{ID: uint(id)}
	product, err := product.GetProduct(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"product": product})
}

// UpdateProduct godoc
// @Summary Update a product.
// @Description Update a product.
// @Tags Product
// @Accept  json
// @Produce  json
// @Param id path int true "Product ID"
// @Param product body UpdateProductInput true "Product"
// @Success 200 {object} models.Product
// @Router /products/{id} [put]
func UpdateProduct(c *gin.Context) {
	var input UpdateProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	var product models.Product
	id, _ := strconv.Atoi(c.Param("id"))
	product = models.Product{ID: uint(id), Name: input.Name, Description: input.Description, ImgUrl: input.ImgUrl}
	product, err := product.UpdateProduct(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"product": product})
}

// DeleteProduct godoc
// @Summary Delete a product.
// @Description Delete a product.
// @Tags Product
// @Accept  json
// @Produce  json
// @Param id path int true "Product ID"
// @Success 200 {object} string
// @Router /products/{id} [delete]
func DeleteProduct(c *gin.Context) {
	var product models.Product
	id, _ := strconv.Atoi(c.Param("id"))
	product = models.Product{ID: uint(id)}
	err := product.DeleteProduct(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"product": "deleted"})
}
