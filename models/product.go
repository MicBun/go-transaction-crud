package models

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

// Product
type Product struct {
	ID                uint      `json:"id" gorm:"primary_key"`
	Name              string    `json:"name" gorm:"type:varchar(255)"`
	Description       string    `json:"description" gorm:"type:varchar(255)"`
	ImgUrl            string    `json:"img_url" gorm:"type:varchar(255)"`
	Stock             int       `json:"stock" gorm:"type:int"`
	ProductCategoryID uint      `json:"product_category_id" gorm:"type:int"`
	CreatedAt         time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt         time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (p *Product) GetProduct(ctx *gin.Context) (Product, error) {
	db := ctx.MustGet("db").(*gorm.DB)
	var product Product
	err := db.Where("id = ?", p.ID).First(&product).Error
	if err != nil {
		return Product{}, err
	}
	return product, nil
}

func GetProducts(ctx *gin.Context) ([]Product, error) {
	db := ctx.MustGet("db").(*gorm.DB)
	var products []Product
	err := db.Find(&products).Error
	if err != nil {
		return []Product{}, err
	}
	return products, nil
}

func (p *Product) CreateProduct(ctx *gin.Context) (Product, error) {
	db := ctx.MustGet("db").(*gorm.DB)
	var product Product
	err := db.Where("name = ?", p.Name).First(&product).Error
	if err != nil {
		err := db.Create(&p).Error
		if err != nil {
			return Product{}, err
		}
		return *p, nil
	}
	return Product{}, errors.New("product already exists")
}

func (p *Product) UpdateProduct(ctx *gin.Context) (Product, error) {
	db := ctx.MustGet("db").(*gorm.DB)
	err := db.Model(&p).Updates(p).Error
	if err != nil {
		return Product{}, err
	}
	return *p, nil
}

func (p *Product) DeleteProduct(ctx *gin.Context) error {
	db := ctx.MustGet("db").(*gorm.DB)
	err := db.Delete(&p).Error
	if err != nil {
		return err
	}
	return nil
}
