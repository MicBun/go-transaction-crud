package controllers

import (
	"github.com/MicBun/go-transaction-crud/models"
	"github.com/MicBun/go-transaction-crud/utils/jwtAuth"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type LoginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Login godoc
// @Summary Login with credential.
// @Description Logging in to get jwt token to access admin or user api by roles.
// @Tags Auth
// @Param Body body LoginInput true "the body to login a user"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /login [post]
func Login(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	userPassword := models.Admin{Email: input.Email, Password: input.Password}

	if userPassword.Email == "admin" && userPassword.Password == "admin" {
		token, err := jwtAuth.GenerateToken(0)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"token": token})
		return
	}

	user, err := userPassword.LoginUser(c)
	if err != nil {
		c.JSON(401, gin.H{"error": "username or password is incorrect."})
		return
	}
	jwtToken, err := jwtAuth.GenerateToken(user.ID)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"token": jwtToken})
}

// Logout godoc
// @Summary Logout.
// @Description Logout.
// @Tags Auth
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /logout [post]
func Logout(c *gin.Context) {
	decoded, err := jwtAuth.ExtractTokenID(c)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}
	user := models.Admin{ID: decoded}
	user.LogoutUser(c)
	c.JSON(200, gin.H{"message": "user logged out"})
}

type RegisterInput struct {
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	BirthDate string `json:"birth_date" binding:"required"`
	Sex       bool   `json:"sex" binding:"required"`
}

// Register godoc
// @Summary Register a new user.
// @Description Register a new user.
// @Param Body body RegisterInput true "the body to register a new user"
// @Tags Admin
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /register [post]
func Register(c *gin.Context) {
	var input RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	decoded, err := jwtAuth.ExtractTokenID(c)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	if decoded != 0 {
		c.JSON(401, gin.H{"error": "Only super admin can create new admin."})
		return
	}

	birthDate, err := time.Parse("2006-01-02", input.BirthDate)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	user := models.Admin{
		Email:     input.Email,
		Password:  input.Password,
		FirstName: input.FirstName,
		LastName:  input.LastName,
		BirthDate: birthDate,
		Sex:       input.Sex,
	}
	userCreated, err := user.RegisterUser(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": userCreated})
}

// GetProfile godoc
// @Summary Get user profile.
// @Description Get user profile.
// @Tags Admin
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /user/me [get]
func GetProfile(c *gin.Context) {
	decoded, err := jwtAuth.ExtractTokenID(c)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}
	user := models.Admin{ID: decoded}
	userProfile, err := user.GetUser(c)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": userProfile})
}

type UpdateProfileInput struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	BirthDate time.Time `json:"birth_date"`
	Password  string    `json:"password"`
}

// UpdateProfile godoc
// @Summary Update user profile.
// @Description Update user profile.
// @Tags Admin
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Param Body body UpdateProfileInput true "the body to update user profile"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /user/me [put]
func UpdateProfile(c *gin.Context) {
	var input UpdateProfileInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	decoded, err := jwtAuth.ExtractTokenID(c)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}
	user := models.Admin{
		ID:        decoded,
		FirstName: input.FirstName,
		LastName:  input.LastName,
		BirthDate: input.BirthDate,
		Password:  input.Password,
	}
	userUpdated, err := user.UpdateUser(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": userUpdated})
}

// DeleteProfile godoc
// @Summary Delete user profile.
// @Description Delete user profile.
// @Tags Admin
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /user/:id [delete]
func DeleteProfile(c *gin.Context) {
	decoded, err := jwtAuth.ExtractTokenID(c)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}
	if decoded != 0 {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}
	id, _ := strconv.Atoi(c.Param("id"))
	user := models.Admin{ID: uint(id)}
	err = user.DeleteUser(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "user deleted"})
}

// GetUsers godoc
// @Summary Get all users.
// @Description Get all users.
// @Tags Admin
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /user/all [get]
func GetUsers(c *gin.Context) {
	decoded, err := jwtAuth.ExtractTokenID(c)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}
	if decoded != 0 {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}
	var users []models.Admin
	users, _ = models.GetUsers(c)
	c.JSON(200, gin.H{"data": users})
}
