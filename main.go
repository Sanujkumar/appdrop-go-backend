package main


import (
	"log"
	"os"

	"appdrop/config"
	"appdrop/models"
	"appdrop/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"appdrop/middleware"

)

func main() {
	godotenv.Load()

	config.ConnectDB()

	config.DB.AutoMigrate(&models.Brand{},&models.Page{}, &models.Widget{})

	r := gin.Default()
	r.Use(middleware.RequestLogger())
	
	routes.SetupRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server running on port", port)
	r.Run(":" + port)
}
