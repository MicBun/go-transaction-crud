package main

import (
	"github.com/MicBun/go-transaction-crud/config"
	"github.com/MicBun/go-transaction-crud/docs"
	"github.com/MicBun/go-transaction-crud/route"
	"github.com/joho/godotenv"
	"log"
)

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @termsOfService http://swagger.io/terms/

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, using default env")
	}

	description := "This is a sample server Activity Tracking server.\n\n" +
		"For this sample, you can use username `user1` and password `password1` after reset to test the authorization filters.\n\n" +
		"To reset the database, use the `reset` query parameter with any value.\n\n" +
		"Checkout my Github: https://github.com/MicBun/go-activity-tracking-api\n\n" +
		"Checkout my Linkedin: https://www.linkedin.com/in/MicBun\n\n"

	docs.SwaggerInfo.Title = "Activity Tracking API"
	docs.SwaggerInfo.Description = description

	database := config.ConnectDataBase()
	sqlDB, _ := database.DB()
	defer sqlDB.Close()
	r := route.SetupRouter(database)
	r.Run()
}
