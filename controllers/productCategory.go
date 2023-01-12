package controllers

import (
	"github.com/MicBun/go-transaction-crud/models"
	"github.com/MicBun/go-transaction-crud/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CreateProductCategoryInput struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

// CreateProductCategory godoc
// @Summary Create a product category.
// @Description Create a product category.
// @Tags ProductCategory
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Param productCategory body CreateProductCategoryInput true "ProductCategory"
// @Success 200 {object} models.ProductCategory
// @Router /product-categories [post]
func CreateProductCategory(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var input CreateProductCategoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.HandleError(c, err)
		return
	}

	productCategory := models.ProductCategory{Name: input.Name, Description: input.Description}
	if err := db.Create(&productCategory).Error; err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(200, gin.H{"message": "success", "data": productCategory})
}

// GetProductCategories godoc
// @Summary Get all product categories.
// @Description Get all product categories.
// @Tags ProductCategory
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} []models.ProductCategory
// @Router /product-categories [get]
func GetProductCategories(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var productCategories []models.ProductCategory
	if err := db.Find(&productCategories).Error; err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(200, gin.H{"message": "success", "data": productCategories})
}

// GetProductCategory godoc
// @Summary Get a product category.
// @Description Get a product category.
// @Tags ProductCategory
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Param id path int true "ProductCategory ID"
// @Success 200 {object} models.ProductCategory
// @Router /product-categories/{id} [get]
func GetProductCategory(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var productCategory models.ProductCategory
	if err := db.Where("id = ?", c.Param("id")).First(&productCategory).Error; err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(200, gin.H{"message": "success", "data": productCategory})
}

type UpdateProductCategoryInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// UpdateProductCategory godoc
// @Summary Update a product category.
// @Description Update a product category.
// @Tags ProductCategory
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Param id path int true "ProductCategory ID"
// @Param productCategory body UpdateProductCategoryInput true "ProductCategory"
// @Success 200 {object} models.ProductCategory
// @Router /product-categories/{id} [put]
func UpdateProductCategory(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var productCategory models.ProductCategory
	if err := db.Where("id = ?", c.Param("id")).First(&productCategory).Error; err != nil {
		utils.HandleError(c, err)
		return
	}

	var input UpdateProductCategoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.HandleError(c, err)
		return
	}

	if err := db.Model(&productCategory).Updates(input).Error; err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(200, gin.H{"message": "success", "data": productCategory})
}

// DeleteProductCategory godoc
// @Summary Delete a product category.
// @Description Delete a product category.
// @Tags ProductCategory
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Param id path int true "ProductCategory ID"
// @Success 200 {object} string
// @Router /product-categories/{id} [delete]
func DeleteProductCategory(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var productCategory models.ProductCategory
	if err := db.Where("id = ?", c.Param("id")).First(&productCategory).Error; err != nil {
		utils.HandleError(c, err)
		return
	}

	if err := db.Delete(&productCategory).Error; err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(200, gin.H{"message": "success"})
}
