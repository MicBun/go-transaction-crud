package route

import (
	"github.com/MicBun/go-transaction-crud/controllers"
	"github.com/MicBun/go-transaction-crud/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	r.POST("/login", controllers.Login)
	r.POST("/register", controllers.Register)

	userRoute := r.Group("/user")
	userRoute.Use(middleware.JwtAndLoggedInMiddleware())
	userRoute.GET("/me", controllers.GetProfile)
	userRoute.PUT("/me", controllers.UpdateProfile)
	userRoute.DELETE("/:id", controllers.DeleteProfile)
	userRoute.POST("/logout", controllers.Logout)
	userRoute.GET("/all", controllers.GetUsers)

	productCategoryRoute := r.Group("/product-category")
	productCategoryRoute.Use(middleware.JwtAndLoggedInMiddleware())
	productCategoryRoute.GET("/", controllers.GetProductCategories)
	productCategoryRoute.POST("/", controllers.CreateProductCategory)
	productCategoryRoute.GET("/:id", controllers.GetProductCategory)
	productCategoryRoute.PUT("/:id", controllers.UpdateProductCategory)
	productCategoryRoute.DELETE("/:id", controllers.DeleteProductCategory)

	productRoute := r.Group("/product")
	productRoute.Use(middleware.JwtAndLoggedInMiddleware())
	productRoute.GET("/", controllers.GetProducts)
	productRoute.POST("/", controllers.CreateProduct)
	productRoute.GET("/:id", controllers.GetProduct)
	productRoute.PUT("/:id", controllers.UpdateProduct)
	productRoute.DELETE("/:id", controllers.DeleteProduct)

	transactionRoute := r.Group("/transaction")
	transactionRoute.Use(middleware.JwtAndLoggedInMiddleware())
	transactionRoute.POST("/", controllers.CreateTransaction)
	transactionRoute.GET("/", controllers.GetTransactions)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
