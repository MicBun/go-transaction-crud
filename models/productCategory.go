package models

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

// ProductCategory
type ProductCategory struct {
	ID          uint      `json:"id" gorm:"primary_key"`
	Name        string    `json:"name" gorm:"type:varchar(255)"`
	Description string    `json:"description" gorm:"type:varchar(255)"`
	Product     []Product `json:"product" gorm:"foreignKey:ProductCategoryID"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (p *ProductCategory) GetProductCategory(ctx *gin.Context) (ProductCategory, error) {
	db := ctx.MustGet("db").(*gorm.DB)
	var productCategory ProductCategory
	err := db.Where("id = ?", p.ID).First(&productCategory).Error
	if err != nil {
		return ProductCategory{}, err
	}
	return productCategory, nil
}

func (p *ProductCategory) GetProductCategories(ctx *gin.Context) ([]ProductCategory, error) {
	db := ctx.MustGet("db").(*gorm.DB)
	var productCategories []ProductCategory
	err := db.Find(&productCategories).Error
	if err != nil {
		return []ProductCategory{}, err
	}
	return productCategories, nil
}

func (p *ProductCategory) CreateProductCategory(ctx *gin.Context) (ProductCategory, error) {
	db := ctx.MustGet("db").(*gorm.DB)
	var productCategory ProductCategory
	err := db.Where("name = ?", p.Name).First(&productCategory).Error
	if err != nil {
		err := db.Create(&p).Error
		if err != nil {
			return ProductCategory{}, err
		}
		return *p, nil
	}
	return ProductCategory{}, errors.New("product category already exists")
}

func (p *ProductCategory) UpdateProductCategory(ctx *gin.Context) (ProductCategory, error) {
	db := ctx.MustGet("db").(*gorm.DB)
	err := db.Model(&p).Updates(p).Error
	if err != nil {
		return ProductCategory{}, err
	}
	return *p, nil
}

func (p *ProductCategory) DeleteProductCategory(ctx *gin.Context) error {
	db := ctx.MustGet("db").(*gorm.DB)
	err := db.Delete(&p).Error
	if err != nil {
		return err
	}
	return nil
}
