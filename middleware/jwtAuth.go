package middleware

import (
	"github.com/MicBun/go-transaction-crud/models"
	"github.com/MicBun/go-transaction-crud/utils/jwtAuth"
	"github.com/gin-gonic/gin"
	"time"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := jwtAuth.TokenValid(c)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": err.Error()})
			return
		}
		c.Next()
	}
}

func JwtAndLoggedInMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := jwtAuth.TokenValid(c)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": err.Error()})
			return
		}

		adminID, err := jwtAuth.ExtractTokenID(c)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": err.Error()})
			return
		}

		user := models.Admin{ID: adminID}
		userResult, err := user.GetUser(c)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": err.Error()})
			return
		}

		if userResult.LoggedIn == true && time.Now().Sub(userResult.LastLogin).Hours() < 24 {
			c.Next()
		}

		c.AbortWithStatusJSON(400, gin.H{"error": "user have not logged in yet"})
	}
}

//func JwtAndClockInMiddleware() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		err := jwtAuth.TokenValid(c)
//		if err != nil {
//			c.AbortWithStatusJSON(401, gin.H{"error": err.Error()})
//			return
//		}
//
//		userID, err := jwtAuth.ExtractTokenID(c)
//		if err != nil {
//			c.AbortWithStatusJSON(401, gin.H{"error": err.Error()})
//			return
//		}
//
//		attendance := models.Attendance{UserID: userID}
//		attendanceResult, err := attendance.GetAttendanceByDate(c, time.Now().Format("2006-01-02"))
//		if err == nil {
//			if attendanceResult.UpdatedAt == attendanceResult.CreatedAt {
//				c.Next()
//			}
//		}
//
//		c.AbortWithStatusJSON(400, gin.H{"error": "user have not clocked in yet or already clocked out"})
//	}
//}
