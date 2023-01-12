package models

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

// Admin
type Admin struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Email     string    `json:"email" gorm:"type:varchar(255)"`
	Password  string    `json:"password" gorm:"type:varchar(255)"`
	FirstName string    `json:"first_name" gorm:"type:varchar(255)"`
	LastName  string    `json:"last_name" gorm:"type:varchar(255)"`
	BirthDate time.Time `json:"birth_date" gorm:"type:varchar(255)"`
	Sex       bool      `json:"sex" gorm:"type:varchar(255)"`
	LastLogin time.Time `json:"last_login" gorm:"type:varchar(255)"`
	LoggedIn  bool      `json:"logged_in" gorm:"type:bool"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (u *Admin) LoginUser(ctx *gin.Context) (Admin, error) {
	db := ctx.MustGet("db").(*gorm.DB)
	var user Admin
	err := db.Where("email = ? AND password = ?", u.Email, u.Password).First(&user).Error
	if err != nil {
		return Admin{}, err
	}
	user.LastLogin = time.Now()
	user.LoggedIn = true
	err = db.Model(&user).Updates(user).Error
	if err != nil {
		return Admin{}, err
	}
	return user, nil
}

func (u *Admin) LogoutUser(ctx *gin.Context) (Admin, error) {
	db := ctx.MustGet("db").(*gorm.DB)
	user, err := u.GetUser(ctx)
	if err != nil {
		return Admin{}, err
	}
	user.LoggedIn = false
	err = db.Model(&user).Updates(user).Error
	if err != nil {
		return Admin{}, err
	}
	return user, nil
}

func (u *Admin) RegisterUser(ctx *gin.Context) (Admin, error) {
	db := ctx.MustGet("db").(*gorm.DB)
	err := db.Create(&u).Error
	if err != nil {
		return Admin{}, err
	}
	return *u, nil
}

func (u *Admin) GetUser(ctx *gin.Context) (Admin, error) {
	db := ctx.MustGet("db").(*gorm.DB)
	var user Admin
	err := db.Where("id = ?", u.ID).First(&user).Error
	if err != nil {
		return Admin{}, err
	}
	return user, nil
}

func (u *Admin) UpdateUser(ctx *gin.Context) (Admin, error) {
	db := ctx.MustGet("db").(*gorm.DB)
	err := db.Model(&u).Updates(u).Error
	if err != nil {
		return Admin{}, err
	}
	return *u, nil
}

func (u *Admin) DeleteUser(ctx *gin.Context) error {
	db := ctx.MustGet("db").(*gorm.DB)
	err := db.Delete(&u).Error
	if err != nil {
		return err
	}
	return nil
}

func GetUsers(ctx *gin.Context) ([]Admin, error) {
	db := ctx.MustGet("db").(*gorm.DB)
	var users []Admin
	err := db.Find(&users).Error
	if err != nil {
		return []Admin{}, err
	}
	return users, nil
}
