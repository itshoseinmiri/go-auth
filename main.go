package main

import (
	"github.com/gin-gonic/gin"
	"auth-api/routes"
	"auth-api/config"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.ConnectDB()

	r := gin.Default()
	routes.SetupRoutes(r)

	port := os.Getenv("PORT")
	r.Run(":" + port)
}
