package main

import (
	"fmt"
	"go-jwt/database"
	"go-jwt/router"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// load env variables
	loadEnv()

	// connect to database
	database.Connect()

	app := gin.Default()

	// set up routes
	router.Route(app)
	port := os.Getenv("SERVER_PORT")

	// run server on port 8080
	app.Run(port)
	fmt.Println("Server running on " + port + " port")
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
